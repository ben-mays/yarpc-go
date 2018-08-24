// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/v2/yarpctransport (interfaces: Router,RouteTable)

// Package yarpctransporttest is a generated GoMock package.
package yarpctransporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	yarpctransport "go.uber.org/yarpc/v2/yarpctransport"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouter) Choose(arg0 context.Context, arg1 *yarpctransport.Request) (yarpctransport.HandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(yarpctransport.HandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouter)(nil).Choose), arg0, arg1)
}

// Procedures mocks base method
func (m *MockRouter) Procedures() []yarpctransport.Procedure {
	ret := m.ctrl.Call(m, "Procedures")
	ret0, _ := ret[0].([]yarpctransport.Procedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMockRecorder) Procedures() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouter)(nil).Procedures))
}

// MockRouteTable is a mock of RouteTable interface
type MockRouteTable struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTableMockRecorder
}

// MockRouteTableMockRecorder is the mock recorder for MockRouteTable
type MockRouteTableMockRecorder struct {
	mock *MockRouteTable
}

// NewMockRouteTable creates a new mock instance
func NewMockRouteTable(ctrl *gomock.Controller) *MockRouteTable {
	mock := &MockRouteTable{ctrl: ctrl}
	mock.recorder = &MockRouteTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouteTable) EXPECT() *MockRouteTableMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouteTable) Choose(arg0 context.Context, arg1 *yarpctransport.Request) (yarpctransport.HandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1)
	ret0, _ := ret[0].(yarpctransport.HandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouteTableMockRecorder) Choose(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouteTable)(nil).Choose), arg0, arg1)
}

// Procedures mocks base method
func (m *MockRouteTable) Procedures() []yarpctransport.Procedure {
	ret := m.ctrl.Call(m, "Procedures")
	ret0, _ := ret[0].([]yarpctransport.Procedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouteTableMockRecorder) Procedures() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouteTable)(nil).Procedures))
}

// Register mocks base method
func (m *MockRouteTable) Register(arg0 []yarpctransport.Procedure) {
	m.ctrl.Call(m, "Register", arg0)
}

// Register indicates an expected call of Register
func (mr *MockRouteTableMockRecorder) Register(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRouteTable)(nil).Register), arg0)
}