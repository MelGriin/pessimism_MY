// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/engine/heuristic (interfaces: Heuristic)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/base-org/pessimism/internal/core"
	heuristic "github.com/base-org/pessimism/internal/engine/heuristic"
	gomock "github.com/golang/mock/gomock"
)

// MockHeuristic is a mock of Heuristic interface.
type MockHeuristic struct {
	ctrl     *gomock.Controller
	recorder *MockHeuristicMockRecorder
}

// MockHeuristicMockRecorder is the mock recorder for MockHeuristic.
type MockHeuristicMockRecorder struct {
	mock *MockHeuristic
}

// NewMockHeuristic creates a new mock instance.
func NewMockHeuristic(ctrl *gomock.Controller) *MockHeuristic {
	mock := &MockHeuristic{ctrl: ctrl}
	mock.recorder = &MockHeuristicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeuristic) EXPECT() *MockHeuristicMockRecorder {
	return m.recorder
}

// Assess mocks base method.
func (m *MockHeuristic) Assess(arg0 core.TransitData) (*heuristic.ActivationSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Assess", arg0)
	ret0, _ := ret[0].(*heuristic.ActivationSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Assess indicates an expected call of Assess.
func (mr *MockHeuristicMockRecorder) Assess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assess", reflect.TypeOf((*MockHeuristic)(nil).Assess), arg0)
}

// InputType mocks base method.
func (m *MockHeuristic) InputType() core.RegisterType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputType")
	ret0, _ := ret[0].(core.RegisterType)
	return ret0
}

// InputType indicates an expected call of InputType.
func (mr *MockHeuristicMockRecorder) InputType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputType", reflect.TypeOf((*MockHeuristic)(nil).InputType))
}

// SUUID mocks base method.
func (m *MockHeuristic) SUUID() core.SUUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SUUID")
	ret0, _ := ret[0].(core.SUUID)
	return ret0
}

// SUUID indicates an expected call of SUUID.
func (mr *MockHeuristicMockRecorder) SUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUUID", reflect.TypeOf((*MockHeuristic)(nil).SUUID))
}

// SetSUUID mocks base method.
func (m *MockHeuristic) SetSUUID(arg0 core.SUUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSUUID", arg0)
}

// SetSUUID indicates an expected call of SetSUUID.
func (mr *MockHeuristicMockRecorder) SetSUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSUUID", reflect.TypeOf((*MockHeuristic)(nil).SetSUUID), arg0)
}

// ValidateInput mocks base method.
func (m *MockHeuristic) ValidateInput(arg0 core.TransitData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateInput", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateInput indicates an expected call of ValidateInput.
func (mr *MockHeuristicMockRecorder) ValidateInput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateInput", reflect.TypeOf((*MockHeuristic)(nil).ValidateInput), arg0)
}
