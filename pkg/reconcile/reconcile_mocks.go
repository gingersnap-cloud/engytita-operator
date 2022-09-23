// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gingersnap-project/operator/pkg/reconcile (interfaces: Context,ContextProvider,Handler)

// Package reconcile is a generated GoMock package.
package reconcile

import (
	context "context"
	reflect "reflect"
	time "time"

	client "github.com/gingersnap-project/operator/pkg/kubernetes/client"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockContext) Client() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockContextMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockContext)(nil).Client))
}

// Ctx mocks base method.
func (m *MockContext) Ctx() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ctx")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Ctx indicates an expected call of Ctx.
func (mr *MockContextMockRecorder) Ctx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ctx", reflect.TypeOf((*MockContext)(nil).Ctx))
}

// Log mocks base method.
func (m *MockContext) Log() logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockContextMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockContext)(nil).Log))
}

// Requeue mocks base method.
func (m *MockContext) Requeue(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Requeue", arg0)
}

// Requeue indicates an expected call of Requeue.
func (mr *MockContextMockRecorder) Requeue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Requeue", reflect.TypeOf((*MockContext)(nil).Requeue), arg0)
}

// RequeueAfter mocks base method.
func (m *MockContext) RequeueAfter(arg0 time.Duration, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequeueAfter", arg0, arg1)
}

// RequeueAfter indicates an expected call of RequeueAfter.
func (mr *MockContextMockRecorder) RequeueAfter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeueAfter", reflect.TypeOf((*MockContext)(nil).RequeueAfter), arg0, arg1)
}

// Status mocks base method.
func (m *MockContext) Status() FlowStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(FlowStatus)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockContextMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockContext)(nil).Status))
}

// StopProcessing mocks base method.
func (m *MockContext) StopProcessing(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopProcessing", arg0)
}

// StopProcessing indicates an expected call of StopProcessing.
func (mr *MockContextMockRecorder) StopProcessing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopProcessing", reflect.TypeOf((*MockContext)(nil).StopProcessing), arg0)
}

// MockContextProvider is a mock of ContextProvider interface.
type MockContextProvider struct {
	ctrl     *gomock.Controller
	recorder *MockContextProviderMockRecorder
}

// MockContextProviderMockRecorder is the mock recorder for MockContextProvider.
type MockContextProviderMockRecorder struct {
	mock *MockContextProvider
}

// NewMockContextProvider creates a new mock instance.
func NewMockContextProvider(ctrl *gomock.Controller) *MockContextProvider {
	mock := &MockContextProvider{ctrl: ctrl}
	mock.recorder = &MockContextProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextProvider) EXPECT() *MockContextProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockContextProvider) Get(arg0 interface{}) (Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockContextProviderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextProvider)(nil).Get), arg0)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockHandler) Handle(arg0 interface{}, arg1 Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", arg0, arg1)
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), arg0, arg1)
}
