// Code generated by MockGen. DO NOT EDIT.
// Source: aws/resource.go

// Package mock_aws is a generated GoMock package.
package mock_aws

import (
	gomock "github.com/golang/mock/gomock"
	aws "github.com/thomas.obenaus/terrastate/aws"
	reflect "reflect"
)

// MockResource is a mock of Resource interface
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// Id mocks base method
func (m *MockResource) Id() string {
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockResourceMockRecorder) Id() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockResource)(nil).Id))
}

// Type mocks base method
func (m *MockResource) Type() aws.Type {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(aws.Type)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockResourceMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockResource)(nil).Type))
}

// String mocks base method
func (m *MockResource) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockResourceMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockResource)(nil).String))
}
