// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/provisioners/api (interfaces: Credential)

package api

import (
	gomock "github.com/golang/mock/gomock"
	// Note: this line is hand-edited as a workaround for https://github.com/golang/mock/issues/4
	context "golang.org/x/net/context"
)

// Mock of Credential interface
type MockCredential struct {
	ctrl     *gomock.Controller
	recorder *_MockCredentialRecorder
}

// Recorder for MockCredential (not exported)
type _MockCredentialRecorder struct {
	mock *MockCredential
}

func NewMockCredential(ctrl *gomock.Controller) *MockCredential {
	mock := &MockCredential{ctrl: ctrl}
	mock.recorder = &_MockCredentialRecorder{mock}
	return mock
}

func (_m *MockCredential) EXPECT() *_MockCredentialRecorder {
	return _m.recorder
}

func (_m *MockCredential) Authenticate(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Authenticate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCredentialRecorder) Authenticate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authenticate", arg0)
}

func (_m *MockCredential) ProvisionerName() string {
	ret := _m.ctrl.Call(_m, "ProvisionerName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCredentialRecorder) ProvisionerName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ProvisionerName")
}

func (_m *MockCredential) Refresh(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Refresh", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCredentialRecorder) Refresh(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Refresh", arg0)
}