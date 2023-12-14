// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package handler is a generated GoMock package.
package handler

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/taguchi-w/example-oapi-codegen/internal/service"
	api "github.com/taguchi-w/example-oapi-codegen/pkg/api"
)

// MockPetService is a mock of PetService interface.
type MockPetService struct {
	ctrl     *gomock.Controller
	recorder *MockPetServiceMockRecorder
}

// MockPetServiceMockRecorder is the mock recorder for MockPetService.
type MockPetServiceMockRecorder struct {
	mock *MockPetService
}

// NewMockPetService creates a new mock instance.
func NewMockPetService(ctrl *gomock.Controller) *MockPetService {
	mock := &MockPetService{ctrl: ctrl}
	mock.recorder = &MockPetServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPetService) EXPECT() *MockPetServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPetService) Create(ctx context.Context, req service.CreatePetRequest) (*api.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(*api.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPetServiceMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPetService)(nil).Create), ctx, req)
}

// Delete mocks base method.
func (m *MockPetService) Delete(ctx context.Context, req service.DeletePetRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPetServiceMockRecorder) Delete(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPetService)(nil).Delete), ctx, req)
}

// Get mocks base method.
func (m *MockPetService) Get(ctx context.Context, req service.GetPetRequest) (*api.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*api.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPetServiceMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPetService)(nil).Get), ctx, req)
}

// List mocks base method.
func (m *MockPetService) List(ctx context.Context, req service.GetPetsRequest) ([]*api.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, req)
	ret0, _ := ret[0].([]*api.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPetServiceMockRecorder) List(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPetService)(nil).List), ctx, req)
}

// Update mocks base method.
func (m *MockPetService) Update(ctx context.Context, req service.UpdatePetRequest) (*api.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req)
	ret0, _ := ret[0].(*api.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPetServiceMockRecorder) Update(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPetService)(nil).Update), ctx, req)
}
