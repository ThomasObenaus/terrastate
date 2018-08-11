// Code generated by MockGen. DO NOT EDIT.
// Source: terraform/infra.go

// Package mock_terraform is a generated GoMock package.
package mock_terraform

import (
	gomock "github.com/golang/mock/gomock"
	terraform "github.com/thomasobenaus/inframapper/terraform"
	reflect "reflect"
)

// MockInfra is a mock of Infra interface
type MockInfra struct {
	ctrl     *gomock.Controller
	recorder *MockInfraMockRecorder
}

// MockInfraMockRecorder is the mock recorder for MockInfra
type MockInfraMockRecorder struct {
	mock *MockInfra
}

// NewMockInfra creates a new mock instance
func NewMockInfra(ctrl *gomock.Controller) *MockInfra {
	mock := &MockInfra{ctrl: ctrl}
	mock.recorder = &MockInfraMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfra) EXPECT() *MockInfraMockRecorder {
	return m.recorder
}

// FindById mocks base method
func (m *MockInfra) FindById(id string) terraform.Resource {
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(terraform.Resource)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockInfraMockRecorder) FindById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockInfra)(nil).FindById), id)
}

// FindByName mocks base method
func (m *MockInfra) FindByName(id string) terraform.Resource {
	ret := m.ctrl.Call(m, "FindByName", id)
	ret0, _ := ret[0].(terraform.Resource)
	return ret0
}

// FindByName indicates an expected call of FindByName
func (mr *MockInfraMockRecorder) FindByName(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockInfra)(nil).FindByName), id)
}

// NumResources mocks base method
func (m *MockInfra) NumResources() int {
	ret := m.ctrl.Call(m, "NumResources")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumResources indicates an expected call of NumResources
func (mr *MockInfraMockRecorder) NumResources() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumResources", reflect.TypeOf((*MockInfra)(nil).NumResources))
}

// String mocks base method
func (m *MockInfra) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockInfraMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockInfra)(nil).String))
}
