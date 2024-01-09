// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

// MockTodoAdapter is a mock of TodoAdapter interface.
type MockTodoAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockTodoAdapterMockRecorder
}

// MockTodoAdapterMockRecorder is the mock recorder for MockTodoAdapter.
type MockTodoAdapterMockRecorder struct {
	mock *MockTodoAdapter
}

// NewMockTodoAdapter creates a new mock instance.
func NewMockTodoAdapter(ctrl *gomock.Controller) *MockTodoAdapter {
	mock := &MockTodoAdapter{ctrl: ctrl}
	mock.recorder = &MockTodoAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoAdapter) EXPECT() *MockTodoAdapterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoAdapter) Create(ctx context.Context, req CreateTodoRequest) (*api.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(*api.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoAdapterMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoAdapter)(nil).Create), ctx, req)
}

// Delete mocks base method.
func (m *MockTodoAdapter) Delete(ctx context.Context, req DeleteTodoRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoAdapterMockRecorder) Delete(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoAdapter)(nil).Delete), ctx, req)
}

// Get mocks base method.
func (m *MockTodoAdapter) Get(ctx context.Context, req GetTodoRequest) (*api.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*api.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTodoAdapterMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodoAdapter)(nil).Get), ctx, req)
}

// List mocks base method.
func (m *MockTodoAdapter) List(ctx context.Context, req GetTodosRequest) ([]*api.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, req)
	ret0, _ := ret[0].([]*api.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTodoAdapterMockRecorder) List(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTodoAdapter)(nil).List), ctx, req)
}

// Update mocks base method.
func (m *MockTodoAdapter) Update(ctx context.Context, req UpdateTodoRequest) (*api.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req)
	ret0, _ := ret[0].(*api.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTodoAdapterMockRecorder) Update(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoAdapter)(nil).Update), ctx, req)
}
