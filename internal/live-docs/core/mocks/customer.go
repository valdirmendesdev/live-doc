// Code generated by MockGen. DO NOT EDIT.
// Source: internal/live-docs/core/customer/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	types "github.com/valdirmendesdev/live-doc/internal/utils/types"
)

// MockCustomer is a mock of Repository interface.
type MockCustomer struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerMockRecorder
}

// MockCustomerMockRecorder is the mock recorder for MockCustomer.
type MockCustomerMockRecorder struct {
	mock *MockCustomer
}

// NewMockCustomer creates a new mock instance.
func NewMockCustomer(ctrl *gomock.Controller) *MockCustomer {
	mock := &MockCustomer{ctrl: ctrl}
	mock.recorder = &MockCustomerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomer) EXPECT() *MockCustomerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCustomer) Add(c *entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", c)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCustomerMockRecorder) Add(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCustomer)(nil).Add), c)
}

// All mocks base method.
func (m *MockCustomer) All(limit, page int) ([]entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", limit, page)
	ret0, _ := ret[0].([]entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockCustomerMockRecorder) All(limit, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockCustomer)(nil).All), limit, page)
}

// GetById mocks base method.
func (m *MockCustomer) GetById(id types.ID) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCustomerMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCustomer)(nil).GetById), id)
}
