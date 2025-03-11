// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/utils/printer/debug.go

// Package mock_printer is a generated GoMock package.
package mock_printer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDebugInterface is a mock of DebugInterface interface.
type MockDebugInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDebugInterfaceMockRecorder
}

// MockDebugInterfaceMockRecorder is the mock recorder for MockDebugInterface.
type MockDebugInterfaceMockRecorder struct {
	mock *MockDebugInterface
}

// NewMockDebugInterface creates a new mock instance.
func NewMockDebugInterface(ctrl *gomock.Controller) *MockDebugInterface {
	mock := &MockDebugInterface{ctrl: ctrl}
	mock.recorder = &MockDebugInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDebugInterface) EXPECT() *MockDebugInterfaceMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockDebugInterface) Debug(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockDebugInterfaceMockRecorder) Debug(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockDebugInterface)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockDebugInterface) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockDebugInterfaceMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockDebugInterface)(nil).Debugf), varargs...)
}

// MockPrinterInterface is a mock of PrinterInterface interface.
type MockPrinterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPrinterInterfaceMockRecorder
}

// MockPrinterInterfaceMockRecorder is the mock recorder for MockPrinterInterface.
type MockPrinterInterfaceMockRecorder struct {
	mock *MockPrinterInterface
}

// NewMockPrinterInterface creates a new mock instance.
func NewMockPrinterInterface(ctrl *gomock.Controller) *MockPrinterInterface {
	mock := &MockPrinterInterface{ctrl: ctrl}
	mock.recorder = &MockPrinterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrinterInterface) EXPECT() *MockPrinterInterfaceMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockPrinterInterface) Error(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockPrinterInterfaceMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockPrinterInterface)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockPrinterInterface) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockPrinterInterfaceMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockPrinterInterface)(nil).Errorf), varargs...)
}

// Exit mocks base method.
func (m *MockPrinterInterface) Exit(exitCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exit", exitCode)
}

// Exit indicates an expected call of Exit.
func (mr *MockPrinterInterfaceMockRecorder) Exit(exitCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockPrinterInterface)(nil).Exit), exitCode)
}

// Fatal mocks base method.
func (m *MockPrinterInterface) Fatal(exitCode int, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{exitCode}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockPrinterInterfaceMockRecorder) Fatal(exitCode interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{exitCode}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockPrinterInterface)(nil).Fatal), varargs...)
}

// Fatalf mocks base method.
func (m *MockPrinterInterface) Fatalf(exitCode int, format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{exitCode, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockPrinterInterfaceMockRecorder) Fatalf(exitCode, format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{exitCode, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockPrinterInterface)(nil).Fatalf), varargs...)
}

// Print mocks base method.
func (m *MockPrinterInterface) Print(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Print", varargs...)
}

// Print indicates an expected call of Print.
func (mr *MockPrinterInterfaceMockRecorder) Print(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPrinterInterface)(nil).Print), args...)
}

// Printf mocks base method.
func (m *MockPrinterInterface) Printf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockPrinterInterfaceMockRecorder) Printf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockPrinterInterface)(nil).Printf), varargs...)
}

// Println mocks base method.
func (m *MockPrinterInterface) Println(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Println", varargs...)
}

// Println indicates an expected call of Println.
func (mr *MockPrinterInterfaceMockRecorder) Println(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Println", reflect.TypeOf((*MockPrinterInterface)(nil).Println), args...)
}
