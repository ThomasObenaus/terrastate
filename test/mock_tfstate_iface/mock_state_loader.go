// Code generated by MockGen. DO NOT EDIT.
// Source: tfstate/iface/state_loader.go

// Package mock_iface is a generated GoMock package.
package mock_iface

import (
	gomock "github.com/golang/mock/gomock"
	terraform "github.com/hashicorp/terraform/terraform"
	iface "github.com/thomas.obenaus/inframapper/tfstate/iface"
	reflect "reflect"
)

// MockStateLoader is a mock of StateLoader interface
type MockStateLoader struct {
	ctrl     *gomock.Controller
	recorder *MockStateLoaderMockRecorder
}

// MockStateLoaderMockRecorder is the mock recorder for MockStateLoader
type MockStateLoaderMockRecorder struct {
	mock *MockStateLoader
}

// NewMockStateLoader creates a new mock instance
func NewMockStateLoader(ctrl *gomock.Controller) *MockStateLoader {
	mock := &MockStateLoader{ctrl: ctrl}
	mock.recorder = &MockStateLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateLoader) EXPECT() *MockStateLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockStateLoader) Load(filename string) (*terraform.State, error) {
	ret := m.ctrl.Call(m, "Load", filename)
	ret0, _ := ret[0].(*terraform.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockStateLoaderMockRecorder) Load(filename interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockStateLoader)(nil).Load), filename)
}

// LoadRemoteState mocks base method
func (m *MockStateLoader) LoadRemoteState(remoteCfg iface.RemoteConfig) ([]*terraform.State, error) {
	ret := m.ctrl.Call(m, "LoadRemoteState", remoteCfg)
	ret0, _ := ret[0].([]*terraform.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadRemoteState indicates an expected call of LoadRemoteState
func (mr *MockStateLoaderMockRecorder) LoadRemoteState(remoteCfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadRemoteState", reflect.TypeOf((*MockStateLoader)(nil).LoadRemoteState), remoteCfg)
}