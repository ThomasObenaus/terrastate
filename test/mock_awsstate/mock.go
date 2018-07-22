// Code generated by MockGen. DO NOT EDIT.
// Source: awsstate/resource_loader.go

// Package mock_awsstate is a generated GoMock package.
package mock_awsstate

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResourceLoader is a mock of ResourceLoader interface
type MockResourceLoader struct {
	ctrl     *gomock.Controller
	recorder *MockResourceLoaderMockRecorder
}

// MockResourceLoaderMockRecorder is the mock recorder for MockResourceLoader
type MockResourceLoaderMockRecorder struct {
	mock *MockResourceLoader
}

// NewMockResourceLoader creates a new mock instance
func NewMockResourceLoader(ctrl *gomock.Controller) *MockResourceLoader {
	mock := &MockResourceLoader{ctrl: ctrl}
	mock.recorder = &MockResourceLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceLoader) EXPECT() *MockResourceLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockResourceLoader) Load() error {
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockResourceLoaderMockRecorder) Load() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockResourceLoader)(nil).Load))
}

// Validate mocks base method
func (m *MockResourceLoader) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockResourceLoaderMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockResourceLoader)(nil).Validate))
}