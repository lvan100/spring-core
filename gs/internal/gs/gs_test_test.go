// Code generated by MockGen. DO NOT EDIT.
// Source: gs.go
//
// Generated by this command:
//
//	mockgen -build_flags="-mod=mod" -package=gs -source=gs.go -destination=gs_test_test.go -exclude_interfaces=BeanSelector,Condition,CondBean,Arg,ReadySignal,BeanInitFunc,BeanDestroyFunc
//

// Package gs is a generated GoMock package.
package gs

import (
	"context"
	"reflect"

	"go.uber.org/mock/gomock"
)

// MockCondContext is a mock of CondContext interface.
type MockCondContext struct {
	ctrl     *gomock.Controller
	recorder *MockCondContextMockRecorder
	isgomock struct{}
}

// MockCondContextMockRecorder is the mock recorder for MockCondContext.
type MockCondContextMockRecorder struct {
	mock *MockCondContext
}

// NewMockCondContext creates a new mock instance.
func NewMockCondContext(ctrl *gomock.Controller) *MockCondContext {
	mock := &MockCondContext{ctrl: ctrl}
	mock.recorder = &MockCondContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCondContext) EXPECT() *MockCondContextMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockCondContext) Find(s BeanSelector) ([]CondBean, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", s)
	ret0, _ := ret[0].([]CondBean)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCondContextMockRecorder) Find(s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCondContext)(nil).Find), s)
}

// Has mocks base method.
func (m *MockCondContext) Has(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockCondContextMockRecorder) Has(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCondContext)(nil).Has), key)
}

// Prop mocks base method.
func (m *MockCondContext) Prop(key string, def ...string) string {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Prop", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Prop indicates an expected call of Prop.
func (mr *MockCondContextMockRecorder) Prop(key any, def ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prop", reflect.TypeOf((*MockCondContext)(nil).Prop), varargs...)
}

// MockArgContext is a mock of ArgContext interface.
type MockArgContext struct {
	ctrl     *gomock.Controller
	recorder *MockArgContextMockRecorder
	isgomock struct{}
}

// MockArgContextMockRecorder is the mock recorder for MockArgContext.
type MockArgContextMockRecorder struct {
	mock *MockArgContext
}

// NewMockArgContext creates a new mock instance.
func NewMockArgContext(ctrl *gomock.Controller) *MockArgContext {
	mock := &MockArgContext{ctrl: ctrl}
	mock.recorder = &MockArgContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArgContext) EXPECT() *MockArgContextMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockArgContext) Bind(v reflect.Value, tag string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", v, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockArgContextMockRecorder) Bind(v, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockArgContext)(nil).Bind), v, tag)
}

// Check mocks base method.
func (m *MockArgContext) Check(c Condition) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", c)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockArgContextMockRecorder) Check(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockArgContext)(nil).Check), c)
}

// Wire mocks base method.
func (m *MockArgContext) Wire(v reflect.Value, tag string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wire", v, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wire indicates an expected call of Wire.
func (mr *MockArgContextMockRecorder) Wire(v, tag any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wire", reflect.TypeOf((*MockArgContext)(nil).Wire), v, tag)
}

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
	isgomock struct{}
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockRunner) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockRunnerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run))
}

// MockJob is a mock of Job interface.
type MockJob struct {
	ctrl     *gomock.Controller
	recorder *MockJobMockRecorder
	isgomock struct{}
}

// MockJobMockRecorder is the mock recorder for MockJob.
type MockJobMockRecorder struct {
	mock *MockJob
}

// NewMockJob creates a new mock instance.
func NewMockJob(ctrl *gomock.Controller) *MockJob {
	mock := &MockJob{ctrl: ctrl}
	mock.recorder = &MockJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJob) EXPECT() *MockJobMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockJob) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockJobMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockJob)(nil).Run), ctx)
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
	isgomock struct{}
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// ListenAndServe mocks base method.
func (m *MockServer) ListenAndServe(sig ReadySignal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe", sig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockServerMockRecorder) ListenAndServe(sig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockServer)(nil).ListenAndServe), sig)
}

// Shutdown mocks base method.
func (m *MockServer) Shutdown(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockServerMockRecorder) Shutdown(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServer)(nil).Shutdown), ctx)
}

// MockBeanRegistration is a mock of BeanRegistration interface.
type MockBeanRegistration struct {
	ctrl     *gomock.Controller
	recorder *MockBeanRegistrationMockRecorder
	isgomock struct{}
}

// MockBeanRegistrationMockRecorder is the mock recorder for MockBeanRegistration.
type MockBeanRegistrationMockRecorder struct {
	mock *MockBeanRegistration
}

// NewMockBeanRegistration creates a new mock instance.
func NewMockBeanRegistration(ctrl *gomock.Controller) *MockBeanRegistration {
	mock := &MockBeanRegistration{ctrl: ctrl}
	mock.recorder = &MockBeanRegistrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeanRegistration) EXPECT() *MockBeanRegistrationMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockBeanRegistration) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBeanRegistrationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBeanRegistration)(nil).Name))
}

// OnProfiles mocks base method.
func (m *MockBeanRegistration) OnProfiles(profiles string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnProfiles", profiles)
}

// OnProfiles indicates an expected call of OnProfiles.
func (mr *MockBeanRegistrationMockRecorder) OnProfiles(profiles any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnProfiles", reflect.TypeOf((*MockBeanRegistration)(nil).OnProfiles), profiles)
}

// SetCaller mocks base method.
func (m *MockBeanRegistration) SetCaller(skip int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCaller", skip)
}

// SetCaller indicates an expected call of SetCaller.
func (mr *MockBeanRegistrationMockRecorder) SetCaller(skip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCaller", reflect.TypeOf((*MockBeanRegistration)(nil).SetCaller), skip)
}

// SetCondition mocks base method.
func (m *MockBeanRegistration) SetCondition(c ...Condition) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range c {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetCondition", varargs...)
}

// SetCondition indicates an expected call of SetCondition.
func (mr *MockBeanRegistrationMockRecorder) SetCondition(c ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCondition", reflect.TypeOf((*MockBeanRegistration)(nil).SetCondition), c...)
}

// SetConfiguration mocks base method.
func (m *MockBeanRegistration) SetConfiguration(c ...Configuration) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range c {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetConfiguration", varargs...)
}

// SetConfiguration indicates an expected call of SetConfiguration.
func (mr *MockBeanRegistrationMockRecorder) SetConfiguration(c ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfiguration", reflect.TypeOf((*MockBeanRegistration)(nil).SetConfiguration), c...)
}

// SetDependsOn mocks base method.
func (m *MockBeanRegistration) SetDependsOn(selectors ...BeanSelector) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range selectors {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetDependsOn", varargs...)
}

// SetDependsOn indicates an expected call of SetDependsOn.
func (mr *MockBeanRegistrationMockRecorder) SetDependsOn(selectors ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDependsOn", reflect.TypeOf((*MockBeanRegistration)(nil).SetDependsOn), selectors...)
}

// SetDestroy mocks base method.
func (m *MockBeanRegistration) SetDestroy(fn BeanDestroyFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDestroy", fn)
}

// SetDestroy indicates an expected call of SetDestroy.
func (mr *MockBeanRegistrationMockRecorder) SetDestroy(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDestroy", reflect.TypeOf((*MockBeanRegistration)(nil).SetDestroy), fn)
}

// SetDestroyMethod mocks base method.
func (m *MockBeanRegistration) SetDestroyMethod(method string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDestroyMethod", method)
}

// SetDestroyMethod indicates an expected call of SetDestroyMethod.
func (mr *MockBeanRegistrationMockRecorder) SetDestroyMethod(method any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDestroyMethod", reflect.TypeOf((*MockBeanRegistration)(nil).SetDestroyMethod), method)
}

// SetExport mocks base method.
func (m *MockBeanRegistration) SetExport(exports ...reflect.Type) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range exports {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetExport", varargs...)
}

// SetExport indicates an expected call of SetExport.
func (mr *MockBeanRegistrationMockRecorder) SetExport(exports ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExport", reflect.TypeOf((*MockBeanRegistration)(nil).SetExport), exports...)
}

// SetInit mocks base method.
func (m *MockBeanRegistration) SetInit(fn BeanInitFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInit", fn)
}

// SetInit indicates an expected call of SetInit.
func (mr *MockBeanRegistrationMockRecorder) SetInit(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInit", reflect.TypeOf((*MockBeanRegistration)(nil).SetInit), fn)
}

// SetInitMethod mocks base method.
func (m *MockBeanRegistration) SetInitMethod(method string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInitMethod", method)
}

// SetInitMethod indicates an expected call of SetInitMethod.
func (mr *MockBeanRegistrationMockRecorder) SetInitMethod(method any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitMethod", reflect.TypeOf((*MockBeanRegistration)(nil).SetInitMethod), method)
}

// SetName mocks base method.
func (m *MockBeanRegistration) SetName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", name)
}

// SetName indicates an expected call of SetName.
func (mr *MockBeanRegistrationMockRecorder) SetName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockBeanRegistration)(nil).SetName), name)
}

// Type mocks base method.
func (m *MockBeanRegistration) Type() reflect.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(reflect.Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockBeanRegistrationMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockBeanRegistration)(nil).Type))
}

// Value mocks base method.
func (m *MockBeanRegistration) Value() reflect.Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(reflect.Value)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockBeanRegistrationMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockBeanRegistration)(nil).Value))
}
