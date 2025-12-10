package gs_init

import (
	"testing"

	"github.com/go-spring/spring-core/conf"
	"github.com/go-spring/spring-core/gs/internal/gs"
	"github.com/go-spring/spring-core/gs/internal/gs_bean"
	"github.com/go-spring/spring-core/gs/internal/gs_cond"
)

var (
	beans   []*gs_bean.BeanDefinition
	modules []Module
)

// BeanProvider defines the API for registering beans in the IoC container.
type BeanProvider interface {
	Provide(objOrCtor any, args ...gs.Arg) *gs_bean.BeanDefinition
}

// ModuleFunc defines the signature of a module function.
type ModuleFunc func(r BeanProvider, p conf.Properties) error

// Module represents a module that can register additional beans
// when certain conditions are met.
type Module struct {
	ModuleFunc ModuleFunc
	Condition  gs.Condition
}

// Beans returns all registered beans.
func Beans() []*gs_bean.BeanDefinition {
	if !testing.Testing() {
		return beans
	}
	var ret []*gs_bean.BeanDefinition
	for _, b := range beans {
		ret = append(ret, b.Clone())
	}
	return ret
}

// Modules returns all registered modules.
func Modules() []Module {
	return modules
}

// Clear clears all registered beans and modules.
func Clear() {
	beans = nil
	modules = nil
}

// Provide registers a bean definition.
func Provide(objOrCtor any, args ...gs.Arg) *gs_bean.BeanDefinition {
	b := gs_bean.NewBean(objOrCtor, args...)
	beans = append(beans, b)
	return b
}

// AddModule registers a module function.
func AddModule(conditions []gs_cond.ConditionOnProperty, fn ModuleFunc) {
	var arr []gs.Condition
	for _, cond := range conditions {
		arr = append(arr, cond)
	}
	modules = append(modules, Module{
		ModuleFunc: fn,
		Condition:  gs_cond.And(arr...),
	})
}
