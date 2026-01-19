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
	"github.com/go-spring/spring-core/gs/internal/gs"
	"github.com/go-spring/spring-core/gs/internal/gs_app"
	"github.com/go-spring/spring-core/gs/internal/gs_bean"
	"github.com/go-spring/stdlib/errutil"
	"github.com/go-spring/stdlib/goutil"
)

// started indicates whether the application has started.
var started bool

// App defines the interface for a Go-Spring application instance.
// It allows setting properties and providing beans to the IoC container.
type App interface {
	// Property sets a key-value property in the application configuration.
	Property(key string, val string)
	// Provide registers an object or constructor as a bean in the application.
	Provide(objOrCtor any, args ...gs.Arg) *gs_bean.BeanDefinition
}

// AppStarter wraps a gs_app.App and manages its lifecycle.
// It provides methods for initialization, configuration, starting,
// stopping, running, and testing the application.
type AppStarter struct {
	app *gs_app.App
	cfg func(App)
}

// NewApp creates a new AppStarter instance.
func NewApp() *AppStarter {
	started = true
	return &AppStarter{
		app: gs_app.NewApp(),
	}
}

// Configure sets the configuration function that will be applied to the application
// before it starts. It returns the AppStarter instance for chaining.
func (s *AppStarter) Configure(f func(App)) *AppStarter {
	s.cfg = f
	return s
}

// Start initializes and starts the application. It prints the banner,
// applies the configuration function if provided, and starts the underlying gs_app.App.
// Returns an error if the application fails to start.
func (s *AppStarter) Start() error {
	// Print banner
	printBanner()

	// Apply user configuration
	if s.cfg != nil {
		s.cfg(s.app)
	}

	// Start application
	if err := s.app.Start(); err != nil {
		err = errutil.Stack(err, "start app failed")
		log.Errorf(context.Background(), log.TagAppDef, "%s", err)
		return err
	}

	return nil
}

// Stop triggers a graceful shutdown of the application and waits
// until all resources and goroutines have completed.
func (s *AppStarter) Stop() {
	s.app.ShutDown()
	s.app.WaitForShutdown()
}

// Run creates and starts a new application using default settings.
func Run() error {
	return NewApp().Run()
}

// Run starts the application, applies configuration, and waits for
// termination signals (e.g., SIGTERM, Ctrl+C) to trigger a graceful shutdown.
// If no servers are running, the application stops immediately.
func (s *AppStarter) Run() error {
	if err := s.Start(); err != nil {
		return err
	}

	// If no servers are running, stop immediately
	if len(s.app.Servers) == 0 {
		s.Stop()
		return nil
	}

	// Listen for termination signals in a separate goroutine
	goutil.Go(s.app.Context(), func(ctx context.Context) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		sig := <-ch
		signal.Stop(ch)
		close(ch)
		log.Infof(ctx, log.TagAppDef, "Received signal: %v", sig)
		s.app.ShutDown()
	}, false)

	// Wait for shutdown to complete
	s.app.WaitForShutdown()
	return nil
}

// RunTest runs a test function using a new application instance.
func RunTest(t *testing.T, f any) {
	NewApp().RunTest(t, f)
}

// RunTest runs a user-defined test function with a provided test object.
// It initializes the application, registers the test object as a bean,
// starts the application, executes the test, and ensures graceful shutdown.
func (s *AppStarter) RunTest(t *testing.T, f any) {
	ft := reflect.TypeOf(f)
	obj := reflect.New(ft.In(0).Elem())

	// Provide the test object as a bean
	s.app.Provide(obj.Interface()).
		Name("__tester__").
		Export(gs.As[any]())

	// Start the application
	if err := s.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() { s.Stop() }()

	// Execute the test function
	reflect.ValueOf(f).Call([]reflect.Value{obj})
}
