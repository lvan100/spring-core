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

type Resolver struct {
	cmd          conf.Properties
	env          conf.Properties
	profileFiles []conf.Properties
	appFiles     []conf.Properties
	prop         conf.Properties
}

func (r *Resolver) sources() []conf.Properties {
	var sources []conf.Properties
	sources = append(sources, r.cmd, r.env)
	sources = append(sources, r.profileFiles...)
	sources = append(sources, r.appFiles...)
	sources = append(sources, r.prop)
	return sources
}

// Exists checks whether a key exists.
func (r *Resolver) Exists(key string) bool {
	for _, s := range r.sources() {
		if s.Exists(key) {
			return true
		}
	}
	return false
}

// Lookup returns the value for a given key, and whether it exists.
func (r *Resolver) Lookup(key string) (string, bool) {
	for _, s := range r.sources() {
		if v, ok := s.Lookup(key); ok {
			return v, true
		}
	}
	return "", false
}

// Refresh refreshes the configuration by merging multiple sources.
func (c *AppConfig) Refresh() (conf.Properties, error) {
	// 1. -----
	env := conf.New()
	cmd := conf.New()

	if err := NewEnvironment().CopyTo(env); err != nil {
		return nil, err
	}
	if err := NewCommandArgs().CopyTo(cmd); err != nil {
		return nil, err
	}

	// 2. -----
	resolver := &Resolver{
		cmd:  cmd,
		env:  env,
		prop: c.Properties,
	}

	confDir, err := conf.ResolveString(resolver, "${spring.app.config.dir:=./conf}")
	if err != nil {
		return nil, err
	}

	var activeProfiles []string
	err = conf.Bind(resolver, &activeProfiles, "${spring.profiles.active:=}")
	if err != nil {
		return nil, err
	}

	// 3. -----
	appFiles, err := loadFiles(resolver, confDir, nil)
	if err != nil {
		return nil, errutil.Stack(err, "refresh error in source local")
	}

	// 4. -----
	profileFiles, err := loadFiles(resolver, confDir, activeProfiles)
	if err != nil {
		return nil, errutil.Stack(err, "refresh error in source local")
	}

	// 5. -----
	var sources []*NamedPropertyCopier
	sources = append(sources, NewNamedPropertyCopier("app", c.Properties))
	sources = append(sources, appFiles...)
	sources = append(sources, profileFiles...)
	sources = append(sources, NewNamedPropertyCopier("env", env))
	sources = append(sources, NewNamedPropertyCopier("cmd", cmd))
	return merge(sources...)
}

// loadFiles loads all candidate configuration files in order and returns
// them as NamedPropertyCopier instances. Non-existent files are skipped,
// while other loading errors abort the process.
func loadFiles(resolver *Resolver, dir string, activeProfiles []string) ([]*NamedPropertyCopier, error) {
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

	var ret []*NamedPropertyCopier
	for _, s := range files {
		filename, err := conf.ResolveString(resolver, s)
		if err != nil {
			return nil, err
		}
		// Load the file
		p, err := conf.Load(filename)
		if err != nil {
			// Don't use `os.IsNotExist`
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
		ret = append(ret, NewNamedPropertyCopier(filename, p))
		if activeProfiles == nil {
			resolver.appFiles = append(resolver.appFiles, p)
		} else {
			resolver.profileFiles = append(resolver.profileFiles, p)
		}

		// Load the file imports
		names, sources, err := loadFileImports(resolver)
		if err != nil {
			return nil, err
		}
		for i, source := range sources {
			ret = append(ret, NewNamedPropertyCopier(names[i], source))
			if activeProfiles == nil {
				resolver.appFiles = append(resolver.appFiles, source)
			} else {
				resolver.profileFiles = append(resolver.profileFiles, source)
			}
		}
	}
	return ret, nil
}

func loadFileImports(p *Resolver) ([]string, []conf.Properties, error) {

	var i struct {
		Imports []string `value:"${spring.app.imports:=}"`
	}
	if err := conf.Bind(p, &i); err != nil {
		return nil, nil, err
	}

	var names []string
	var sources []conf.Properties
	for _, source := range i.Imports {
		c, err := conf.Load(source)
		if err != nil {
			return nil, nil, err
		}
		names = append(names, source)
		sources = append(sources, c)
	}
	return names, sources, nil
}
