/*
 * Copyright 2025 The Go-Spring Authors.
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

package gs

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"testing"

	"github.com/go-spring/log"
	"github.com/go-spring/spring-base/util"
	"github.com/go-spring/spring-core/gs/internal/gs"
	"github.com/go-spring/spring-core/gs/internal/gs_app"
)

// started indicates whether the application has started.
var started bool

// AppStarter is a wrapper to manage the lifecycle of a Spring application.
// It handles initialization, running, graceful shutdown, and logging.
type AppStarter struct {
	app *gs_app.App
}

// NewApp creates a new AppStarter instance.
func NewApp() *AppStarter {
	started = true
	return &AppStarter{
		app: gs_app.NewApp(),
	}
}

// Configure sets the application configuration provider.
func (s *AppStarter) Configure(f func(cfg gs_app.AppConfigurer)) *AppStarter {
	s.app.Configure(f)
	return s
}

// Start starts the application asynchronously and returns a function
// that can be used to trigger shutdown from outside.
func (s *AppStarter) Start() error {

	// Print banner
	printBanner()

	// Start application
	if err := s.app.Start(); err != nil {
		err = util.WrapError(err, "start app failed")
		log.Errorf(context.Background(), log.TagAppDef, "%s", err)
		return err
	}

	return nil
}

// Stop triggers graceful shutdown of the application.
func (s *AppStarter) Stop() {
	s.app.ShutDown()
	s.app.WaitForShutdown()
}

// Run starts the application with a custom run function.
func Run() error {
	return NewApp().Run()
}

// Run starts the application, optionally runs a user-defined callback,
// and waits for termination signals (e.g., SIGTERM, Ctrl+C) to trigger graceful shutdown.
func (s *AppStarter) Run() error {

	if err := s.Start(); err != nil {
		return err
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		sig := <-ch
		signal.Stop(ch)
		close(ch)
		log.Infof(context.Background(), log.TagAppDef, "Received signal: %v", sig)
		s.app.ShutDown()
	}()

	s.app.WaitForShutdown()
	return nil
}

// RunTest runs a test function.
func RunTest(t *testing.T, f any) {
	NewApp().RunTest(t, f)
}

// RunTest runs a user-defined test function.
func (s *AppStarter) RunTest(t *testing.T, f any) {
	ft := reflect.TypeOf(f)
	obj := reflect.New(ft.In(0).Elem())

	// 提供测试对象
	s.app.Provide(obj.Interface()).
		Name("__tester__").
		Export(gs.As[any]())

	if err := s.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() { s.Stop() }()

	// 执行测试函数
	reflect.ValueOf(f).Call([]reflect.Value{obj})
}
