// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/shared/acquirers_gateway.shared.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAcquiresGateway is a mock of AcquiresGateway interface.
type MockAcquiresGateway struct {
	ctrl     *gomock.Controller
	recorder *MockAcquiresGatewayMockRecorder
}

// MockAcquiresGatewayMockRecorder is the mock recorder for MockAcquiresGateway.
type MockAcquiresGatewayMockRecorder struct {
	mock *MockAcquiresGateway
}

// NewMockAcquiresGateway creates a new mock instance.
func NewMockAcquiresGateway(ctrl *gomock.Controller) *MockAcquiresGateway {
	mock := &MockAcquiresGateway{ctrl: ctrl}
	mock.recorder = &MockAcquiresGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAcquiresGateway) EXPECT() *MockAcquiresGatewayMockRecorder {
	return m.recorder
}

// CieloAuth mocks base method.
func (m *MockAcquiresGateway) CieloAuth(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CieloAuth", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// CieloAuth indicates an expected call of CieloAuth.
func (mr *MockAcquiresGatewayMockRecorder) CieloAuth(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CieloAuth", reflect.TypeOf((*MockAcquiresGateway)(nil).CieloAuth), key)
}

// CieloSendTransaction mocks base method.
func (m *MockAcquiresGateway) CieloSendTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CieloSendTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// CieloSendTransaction indicates an expected call of CieloSendTransaction.
func (mr *MockAcquiresGatewayMockRecorder) CieloSendTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CieloSendTransaction", reflect.TypeOf((*MockAcquiresGateway)(nil).CieloSendTransaction))
}

// RedeAuth mocks base method.
func (m *MockAcquiresGateway) RedeAuth(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeAuth", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// RedeAuth indicates an expected call of RedeAuth.
func (mr *MockAcquiresGatewayMockRecorder) RedeAuth(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeAuth", reflect.TypeOf((*MockAcquiresGateway)(nil).RedeAuth), key)
}

// RedeSendTransaction mocks base method.
func (m *MockAcquiresGateway) RedeSendTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeSendTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// RedeSendTransaction indicates an expected call of RedeSendTransaction.
func (mr *MockAcquiresGatewayMockRecorder) RedeSendTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeSendTransaction", reflect.TypeOf((*MockAcquiresGateway)(nil).RedeSendTransaction))
}

// StoneAuth mocks base method.
func (m *MockAcquiresGateway) StoneAuth(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoneAuth", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// StoneAuth indicates an expected call of StoneAuth.
func (mr *MockAcquiresGatewayMockRecorder) StoneAuth(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoneAuth", reflect.TypeOf((*MockAcquiresGateway)(nil).StoneAuth), key)
}

// StoneSendTransaction mocks base method.
func (m *MockAcquiresGateway) StoneSendTransaction() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoneSendTransaction")
	ret0, _ := ret[0].(error)
	return ret0
}

// StoneSendTransaction indicates an expected call of StoneSendTransaction.
func (mr *MockAcquiresGatewayMockRecorder) StoneSendTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoneSendTransaction", reflect.TypeOf((*MockAcquiresGateway)(nil).StoneSendTransaction))
}
