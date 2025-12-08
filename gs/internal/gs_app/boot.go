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

package gs_app

import (
	"github.com/go-spring/spring-core/conf"
	"github.com/go-spring/spring-core/gs/internal/gs"
	"github.com/go-spring/spring-core/gs/internal/gs_conf"
	"github.com/go-spring/spring-core/gs/internal/gs_core"
)

// Boot defines the interface for application bootstrapping.
type Boot interface {
	// Config returns the boot configuration.
	Config() *gs_conf.BootConfig
	// Provide registers a bean definition.
	Provide(objOrCtor any, args ...gs.Arg) *gs.RegisteredBean
}

// BootImpl is the concrete implementation of the Boot interface.
// It manages the application's bootstrapping process.
type BootImpl struct {
	c *gs_core.Container  // The IoC container
	p *gs_conf.BootConfig // The boot configuration

	// flag indicates whether the bootstrapper has been used.
	flag bool

	Runners []gs.Runner `autowire:"${spring.boot.runners:=?}"`
}

// NewBoot creates and returns a new BootImpl instance.
func NewBoot() Boot {
	return &BootImpl{
		c: gs_core.New(),
		p: gs_conf.NewBootConfig(),
	}
}

// Config returns the boot configuration.
func (b *BootImpl) Config() *gs_conf.BootConfig {
	return b.p
}

// Root registers a root bean definition in the container.
func (b *BootImpl) Root(x *gs.RegisteredBean) {
	b.c.Root(x)
}

// Provide registers a bean definition.
func (b *BootImpl) Provide(objOrCtor any, args ...gs.Arg) *gs.RegisteredBean {
	b.flag = true
	return b.c.Provide(objOrCtor, args...).Caller(1)
}

// Run executes the application's boot process.
func (b *BootImpl) Run() error {
	// If no beans were registered, there’s nothing to run.
	if !b.flag {
		return nil
	}
	b.c.Root(b.c.Provide(b))

	var p conf.Properties

	// Refresh the boot configuration.
	{
		var err error
		if p, err = b.p.Refresh(); err != nil {
			return err
		}
	}

	// Refresh the container.
	if err := b.c.Refresh(p); err != nil {
		return err
	}

	// Execute all registered runners.
	for _, r := range b.Runners {
		if err := r.Run(); err != nil {
			return err
		}
	}

	b.c.Close()
	return nil
}
