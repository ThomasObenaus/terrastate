// Code generated by MockGen. DO NOT EDIT.
// Source: aws/infra.go

// Package mock_aws is a generated GoMock package.
package mock_aws

import (
	gomock "github.com/golang/mock/gomock"
	aws "github.com/thomas.obenaus/inframapper/aws"
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
func (m *MockInfra) FindById(id string) aws.Resource {
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(aws.Resource)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockInfraMockRecorder) FindById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockInfra)(nil).FindById), id)
}

// FindVpc mocks base method
func (m *MockInfra) FindVpc(id string) *aws.Vpc {
	ret := m.ctrl.Call(m, "FindVpc", id)
	ret0, _ := ret[0].(*aws.Vpc)
	return ret0
}

// FindVpc indicates an expected call of FindVpc
func (mr *MockInfraMockRecorder) FindVpc(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindVpc", reflect.TypeOf((*MockInfra)(nil).FindVpc), id)
}

// Vpcs mocks base method
func (m *MockInfra) Vpcs() []*aws.Vpc {
	ret := m.ctrl.Call(m, "Vpcs")
	ret0, _ := ret[0].([]*aws.Vpc)
	return ret0
}

// Vpcs indicates an expected call of Vpcs
func (mr *MockInfraMockRecorder) Vpcs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vpcs", reflect.TypeOf((*MockInfra)(nil).Vpcs))
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

// Region mocks base method
func (m *MockInfra) Region() string {
	ret := m.ctrl.Call(m, "Region")
	ret0, _ := ret[0].(string)
	return ret0
}

// Region indicates an expected call of Region
func (mr *MockInfraMockRecorder) Region() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Region", reflect.TypeOf((*MockInfra)(nil).Region))
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
