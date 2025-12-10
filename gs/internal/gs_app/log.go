package gs_app

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/go-spring/log"
	"github.com/go-spring/spring-base/util"
)

// initLog initializes the application's logging system.
func (app *App) initLog() error {

	// Step 1: Refresh the global system configuration.
	p, err := app.p.SysConfig()
	if err != nil {
		return util.FormatError(err, "refresh error in source sys")
	}

	// Step 2: Load logging-related configuration parameters.
	var c struct {
		// LocalDir is the directory that contains configuration files.
		// Defaults to "./conf" if not provided.
		LocalDir string `value:"${spring.app.config.dir:=./conf}"`

		// Profiles specifies the active application profile(s),
		// such as "dev", "prod", etc.
		// Multiple profiles can be provided as a comma-separated list.
		Profiles string `value:"${spring.profiles.active:=}"`
	}
	if err = p.Bind(&c); err != nil {
		return util.FormatError(err, "bind error in source sys")
	}

	extensions := []string{".properties", ".yaml", ".yml", ".xml", ".json"}

	// Step 3: Build a list of candidate configuration files.
	var files []string
	if profiles := strings.TrimSpace(c.Profiles); profiles != "" {
		for s := range strings.SplitSeq(profiles, ",") { // NOTE: range returns index
			if s = strings.TrimSpace(s); s != "" {
				for _, ext := range extensions {
					files = append(files, filepath.Join(c.LocalDir, "log-"+s+ext))
				}
			}
		}
	}
	for _, ext := range extensions {
		files = append(files, filepath.Join(c.LocalDir, "log"+ext))
	}

	// Step 4: Detect existing configuration files.
	var logFiles []string
	for _, s := range files {
		if ok, err := util.PathExists(s); err != nil {
			return err
		} else if ok {
			logFiles = append(logFiles, s)
		}
	}

	// Step 5: Apply logging configuration or fall back to defaults.
	switch n := len(logFiles); {
	case n == 0:
		log.Infof(context.Background(), log.TagAppDef, "no log configuration file found, using default logger")
		return nil
	case n > 1:
		return util.FormatError(nil, "multiple log files found: %s", logFiles)
	default:
		return log.RefreshFile(logFiles[0])
	}
}
