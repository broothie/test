// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/broothie/test (interfaces: TestingT)
//
// Generated by this command:
//
//	mockgen --destination mocks/testing_t.go --package mocks . TestingT
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTestingT is a mock of TestingT interface.
type MockTestingT struct {
	ctrl     *gomock.Controller
	recorder *MockTestingTMockRecorder
	isgomock struct{}
}

// MockTestingTMockRecorder is the mock recorder for MockTestingT.
type MockTestingTMockRecorder struct {
	mock *MockTestingT
}

// NewMockTestingT creates a new mock instance.
func NewMockTestingT(ctrl *gomock.Controller) *MockTestingT {
	mock := &MockTestingT{ctrl: ctrl}
	mock.recorder = &MockTestingTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestingT) EXPECT() *MockTestingTMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockTestingT) Error(arg0 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockTestingTMockRecorder) Error(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockTestingT)(nil).Error), arg0...)
}

// FailNow mocks base method.
func (m *MockTestingT) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MockTestingTMockRecorder) FailNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MockTestingT)(nil).FailNow))
}

// Helper mocks base method.
func (m *MockTestingT) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MockTestingTMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockTestingT)(nil).Helper))
}
