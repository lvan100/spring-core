/*
 * Copyright 2024 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package gs_conf provides a layered configuration system for Go-Spring
// applications. It consolidates multiple configuration sources into a
// single immutable property set, supporting profile-specific files
// and optional import of additional configuration files.
//
// Supported configuration sources include:
//   - Built-in system defaults (SysConf)
//   - Local configuration files (e.g., ./conf/app.yaml)
//   - Remote configuration files (from config servers)
//   - Dynamically supplied remote properties
//   - Operating system environment variables
//   - Command-line arguments
//
// Sources are applied in a defined order; later sources override
// earlier ones when the same key is defined multiple times.
package gs_conf

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-spring/spring-core/conf"
	"github.com/go-spring/stdlib/errutil"
)

// PropertyCopier defines the interface for any configuration source
// that can copy its key-value pairs into a MutableProperties instance.
type PropertyCopier interface {
	CopyTo(out *conf.MutableProperties) error
}

// NamedPropertyCopier is a wrapper around PropertyCopier that carries
// a human-readable name. The Name field is used for logging, debugging,
// or error reporting when merging multiple configuration sources.
type NamedPropertyCopier struct {
	PropertyCopier
	Name string
}

// NewNamedPropertyCopier creates a new NamedPropertyCopier instance
// with the given name and underlying PropertyCopier.
func NewNamedPropertyCopier(name string, p PropertyCopier) *NamedPropertyCopier {
	return &NamedPropertyCopier{PropertyCopier: p, Name: name}
}

// CopyTo copies the properties from the underlying PropertyCopier into
// the given MutableProperties instance. If PropertyCopier is nil, this
// method does nothing and returns nil.
func (c *NamedPropertyCopier) CopyTo(out *conf.MutableProperties) error {
	if c.PropertyCopier != nil {
		return c.PropertyCopier.CopyTo(out)
	}
	return nil
}

// AppConfig represents the layered configuration of an application.
// The typical merge order is:
//  1. System defaults (SysConf)
//  2. Local configuration files
//  3. Remote configuration files
//  4. Dynamically supplied remote properties
//  5. Environment variables
//  6. Command-line arguments
//
// Later layers override earlier ones in case of key conflicts.
type AppConfig struct {
	Properties *conf.MutableProperties
}

// NewAppConfig creates a new AppConfig instance.
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Properties: conf.New(),
	}
}

// merge combines multiple NamedPropertyCopier instances into a single
// Properties instance. Sources are applied in order; later sources
// override earlier ones. If any source fails to copy, merge aborts
// and returns an error identifying the failing source.
func merge(sources ...*NamedPropertyCopier) (conf.Properties, error) {
	out := conf.New()
	for _, s := range sources {
		if s != nil {
			if err := s.CopyTo(out); err != nil {
				return nil, errutil.Stack(err, "merge error in source %s", s.Name)
			}
		}
	}
	return out, nil
}

// Refresh merges all configuration layers into a read-only Properties instance.
// If useImport is true, it additionally loads and merges imported configuration
// files defined via the "spring.app.imports" property.
func (c *AppConfig) Refresh(useImport bool) (conf.Properties, error) {
	env := NewEnvironment()
	cmd := NewCommandArgs()

	p, err := merge(
		NewNamedPropertyCopier("app", c.Properties),
		NewNamedPropertyCopier("env", env),
		NewNamedPropertyCopier("cmd", cmd),
	)
	if err != nil {
		return nil, err
	}

	// Load local configuration files
	localFiles, err := loadFiles(p)
	if err != nil {
		return nil, errutil.Stack(err, "refresh error in source local")
	}

	var sources []*NamedPropertyCopier
	sources = append(sources, NewNamedPropertyCopier("app", c.Properties))
	sources = append(sources, localFiles...)
	sources = append(sources, NewNamedPropertyCopier("env", env))
	sources = append(sources, NewNamedPropertyCopier("cmd", cmd))
	if p, err = merge(sources...); err != nil {
		return nil, err
	}

	// Skip imports if not enabled
	if !useImport {
		return p, nil
	}

	var i struct {
		Imports []string `value:"${spring.app.imports:=}"`
	}
	if err = p.Bind(&i); err != nil {
		return nil, err
	}

	sources = []*NamedPropertyCopier{}
	sources = append(sources, NewNamedPropertyCopier("app", c.Properties))
	sources = append(sources, localFiles...)
	for _, source := range i.Imports {
		if p, err = conf.Load(source); err != nil {
			return nil, err
		}
		if p != nil {
			sources = append(sources, NewNamedPropertyCopier(source, p))
		}
	}
	sources = append(sources, NewNamedPropertyCopier("env", env))
	sources = append(sources, NewNamedPropertyCopier("cmd", cmd))
	return merge(sources...)
}

// getAppFiles generates a list of candidate configuration file paths,
// including both base files (app.yaml, app.properties, etc.) and
// profile-specific variants (app-dev.yaml, app-prod.properties, etc.).
func getAppFiles(dir string, activeProfiles []string) ([]string, error) {
	extensions := []string{".properties", ".yaml", ".yml", ".toml", ".tml", ".json"}

	var files []string
	for _, ext := range extensions {
		files = append(files, filepath.Join(dir, "app"+ext))
	}

	for _, s := range activeProfiles {
		for _, ext := range extensions {
			files = append(files, filepath.Join(dir, "app-"+s+ext))
		}
	}
	return files, nil
}

// loadFiles loads all candidate configuration files in order and returns
// them as NamedPropertyCopier instances. Non-existent files are skipped,
// while other loading errors abort the process.
func loadFiles(resolver conf.Properties) ([]*NamedPropertyCopier, error) {
	dir, err := resolver.Resolve("${spring.app.config.dir:=./conf}")
	if err != nil {
		return nil, err
	}

	strActiveProfiles, err := resolver.Resolve("${spring.profiles.active:=}")
	if err != nil {
		return nil, err
	}

	var activeProfiles []string
	for s := range strings.SplitSeq(strActiveProfiles, ",") {
		if s = strings.TrimSpace(s); s != "" {
			activeProfiles = append(activeProfiles, s)
		}
	}

	files, err := getAppFiles(dir, activeProfiles)
	if err != nil {
		return nil, err
	}

	var ret []*NamedPropertyCopier
	for _, s := range files {
		filename, err := resolver.Resolve(s)
		if err != nil {
			return nil, err
		}
		c, err := conf.Load(filename)
		if err != nil {
			// Don't use `os.IsNotExist`
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
		ret = append(ret, NewNamedPropertyCopier(filename, c))
	}
	return ret, nil
}
