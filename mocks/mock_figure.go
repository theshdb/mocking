// Code generated by MockGen. DO NOT EDIT.
// Source: learning/interface/shape (interfaces: Figure)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFigure is a mock of Figure interface.
type MockFigure struct {
	ctrl     *gomock.Controller
	recorder *MockFigureMockRecorder
}

// MockFigureMockRecorder is the mock recorder for MockFigure.
type MockFigureMockRecorder struct {
	mock *MockFigure
}

// NewMockFigure creates a new mock instance.
func NewMockFigure(ctrl *gomock.Controller) *MockFigure {
	mock := &MockFigure{ctrl: ctrl}
	mock.recorder = &MockFigureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFigure) EXPECT() *MockFigureMockRecorder {
	return m.recorder
}

// Area mocks base method.
func (m *MockFigure) Area() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Area")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Area indicates an expected call of Area.
func (mr *MockFigureMockRecorder) Area() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Area", reflect.TypeOf((*MockFigure)(nil).Area))
}