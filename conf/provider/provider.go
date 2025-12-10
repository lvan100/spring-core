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

package provider

import (
	"strings"

	"github.com/go-spring/spring-base/util"
	"github.com/go-spring/spring-core/conf/reader"
	"github.com/lvan100/golib/flatten"
)

var providers = map[string]Provider{}

func init() {
	RegisterProvider("file", LoadFile)
}

// Provider provides configuration data from a specific source.
type Provider func(source string) (*flatten.Storage, error)

// RegisterProvider registers a Provider for a specific configuration source.
func RegisterProvider(name string, p Provider) {
	providers[name] = p
}

// Load loads a configuration source and returns its content as a map.
func Load(source string) (*flatten.Storage, error) {
	name := "file"
	if i := strings.Index(source, ":"); i > 0 {
		name = source[:i]
		source = source[i+1:]
	}
	p, ok := providers[name]
	if !ok {
		err := util.FormatError(nil, "unsupported provider type %s", name)
		return nil, util.FormatError(err, "read config %s error", source)
	}
	return p(source)
}

// LoadFile loads a file and returns its content as a map.
func LoadFile(source string) (*flatten.Storage, error) {
	m, err := reader.ReadFile(source)
	if err != nil {
		return nil, err
	}
	s := flatten.NewStorage()
	fileID := s.AddFile(source)
	for k, v := range flatten.Flatten(m) {
		if err = s.Set(k, v, fileID); err != nil {
			return nil, err
		}
	}
	return s, nil
}
