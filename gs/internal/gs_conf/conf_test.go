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

package gs_conf

import (
	"errors"
	"os"
	"testing"

	"github.com/go-spring/spring-base/testing/assert"
	"github.com/go-spring/spring-core/conf"
)

func clean() {
	os.Args = nil
	os.Clearenv()
}

func TestAppConfig(t *testing.T) {
	clean()

	t.Run("local dir resolve error", func(t *testing.T) {
		t.Cleanup(clean)
		_ = os.Setenv("GS_SPRING_APP_CONFIG_DIR", "${a}")
		_, err := NewAppConfig().Refresh()
		assert.Error(t, err).Matches(`resolve string "\${a}" error: property \"a\" not exist`)
	})

	t.Run("success", func(t *testing.T) {
		t.Cleanup(clean)
		_ = os.Setenv("GS_SPRING_APP_CONFIG_DIR", "./testdata/conf")
		p, err := NewAppConfig().Refresh()
		assert.That(t, err).Nil()
		assert.That(t, p.Data()).Equal(map[string]string{
			"spring.app.config.dir": "./testdata/conf",
			"spring.app.name":       "test",
			"http.server.addr":      "0.0.0.0:8080",
		})
	})

	t.Run("merge error - env", func(t *testing.T) {
		t.Cleanup(clean)
		_ = os.Setenv("GS_A", "a")
		_ = os.Setenv("GS_A_B", "a.b")
		_, err := NewAppConfig().Refresh()
		assert.Error(t, err).Matches("property conflict at path a.b")
	})

	t.Run("merge error - sys", func(t *testing.T) {
		t.Cleanup(clean)
		_ = os.Setenv("GS_SPRING_APP_CONFIG_DIR", "./testdata/conf")
		c := NewAppConfig()
		fileID := c.Properties.AddFile("conf_test.go")
		_ = c.Properties.Set("http.server[0].addr", "0.0.0.0:8080", fileID)
		_, err := c.Refresh()
		assert.Error(t, err).Matches("property conflict at path http.server.addr")
	})

	t.Run("load from sys conf", func(t *testing.T) {
		t.Cleanup(clean)
		c := NewAppConfig()
		fileID := c.Properties.AddFile("test")
		_ = c.Properties.Set("spring.app.name", "sysconf-test", fileID)
		p, err := c.Refresh()
		assert.That(t, err).Nil()
		assert.That(t, p.Get("spring.app.name")).Equal("sysconf-test")
	})
}

func TestPropertySources(t *testing.T) {
	clean()

	t.Run("non existent directory", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddDir("non_existent_dir")
		assert.That(t, 1).Equal(len(ps.extraDirs))
	})

	t.Run("dir is file", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		assert.Panic(t, func() {
			ps.AddDir("./testdata/conf/app.properties")
		}, "should be a directory")
	})

	t.Run("dir access denied", func(t *testing.T) {
		t.Cleanup(clean)
		defer func() { osStat = os.Stat }()
		osStat = func(name string) (os.FileInfo, error) {
			return nil, errors.New("permission denied")
		}
		ps := NewPropertySources()
		assert.Panic(t, func() {
			ps.AddDir("./testdata/conf/app.properties")
		}, "permission denied")
	})

	t.Run("non existent file", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddFile("non_existent_file")
		assert.That(t, 1).Equal(len(ps.extraFiles))
	})

	t.Run("file is directory", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		assert.Panic(t, func() {
			ps.AddFile("./testdata/conf")
		}, "should be a file")
	})

	t.Run("file access denied", func(t *testing.T) {
		t.Cleanup(clean)
		defer func() { osStat = os.Stat }()
		osStat = func(name string) (os.FileInfo, error) {
			return nil, errors.New("permission denied")
		}
		ps := NewPropertySources()
		assert.Panic(t, func() {
			ps.AddFile("./testdata/conf")
		}, "permission denied")
	})

	t.Run("reset extra dirs and files", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddFile("./testdata/conf/app.properties")
		assert.That(t, 1).Equal(len(ps.extraFiles))
		ps.AddDir("./testdata/conf")
		assert.That(t, 1).Equal(len(ps.extraDirs))
		ps.Reset()
		assert.That(t, 0).Equal(len(ps.extraFiles))
		assert.That(t, 0).Equal(len(ps.extraDirs))
	})

	t.Run("get config files without profiles", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		files, err := ps.getFiles("./conf", conf.Map(nil))
		assert.That(t, err).Nil()
		assert.That(t, files).Equal([]string{
			"conf/app.properties",
			"conf/app.yaml",
			"conf/app.yml",
			"conf/app.toml",
			"conf/app.tml",
			"conf/app.json",
		})
	})

	t.Run("get config files with profiles", func(t *testing.T) {
		t.Cleanup(clean)
		p := conf.Map(map[string]any{
			"spring.profiles.active": "dev,test",
		})
		ps := NewPropertySources()
		files, err := ps.getFiles("./conf", p)
		assert.That(t, err).Nil()
		assert.That(t, files).Equal([]string{
			"conf/app.properties",
			"conf/app.yaml",
			"conf/app.yml",
			"conf/app.toml",
			"conf/app.tml",
			"conf/app.json",
			"conf/app-dev.properties",
			"conf/app-dev.yaml",
			"conf/app-dev.yml",
			"conf/app-dev.toml",
			"conf/app-dev.tml",
			"conf/app-dev.json",
			"conf/app-test.properties",
			"conf/app-test.yaml",
			"conf/app-test.yml",
			"conf/app-test.toml",
			"conf/app-test.tml",
			"conf/app-test.json",
		})
	})

	t.Run("load files from property sources", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddFile("./testdata/conf/app.properties")
		files, err := ps.loadFiles(conf.Map(nil))
		assert.That(t, err).Nil()
		assert.That(t, 1).Equal(len(files))
	})

	t.Run("profile resolve error", func(t *testing.T) {
		t.Cleanup(clean)
		p := conf.Map(map[string]any{
			"spring.profiles.active": "${a}",
		})
		ps := NewPropertySources()
		_, err := ps.loadFiles(p)
		assert.Error(t, err).Matches(`resolve string "\${a}" error: property \"a\" not exist`)
	})

	t.Run("file path resolve error", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddFile("./testdata/conf/app-${a}.properties")
		_, err := ps.loadFiles(conf.Map(nil))
		assert.Error(t, err).Matches("property \"a\" not exist")
	})

	t.Run("config file load error", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddFile("./testdata/conf/error.json")
		_, err := ps.loadFiles(conf.Map(nil))
		assert.Error(t, err).Matches("cannot unmarshal .*")
	})

	t.Run("load files with non-existent dir", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddDir("non_existent_dir")
		files, err := ps.loadFiles(conf.Map(nil))
		assert.That(t, err).Nil()
		assert.That(t, 0).Equal(len(files))
	})

	t.Run("add multiple directories and files", func(t *testing.T) {
		t.Cleanup(clean)
		ps := NewPropertySources()
		ps.AddDir("./testdata/conf")
		ps.AddDir("./testdata/conf/remote")
		ps.AddFile("./testdata/conf/app.properties")
		ps.AddFile("./testdata/conf/remote/app.properties")
		assert.That(t, 2).Equal(len(ps.extraDirs))
		assert.That(t, 2).Equal(len(ps.extraFiles))
	})
}
