// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/brunno98/voucher-manager/domain/code (interfaces: Repository)

// Package mock_code is a generated GoMock package.
package code

import (
        reflect "reflect"
        time "time"

        gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
        ctrl     *gomock.Controller
        recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
        mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
        mock := &MockRepository{ctrl: ctrl}
        mock.recorder = &MockRepositoryMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
        return m.recorder
}

// GetLastRecoveredDates mocks base method.
func (m *MockRepository) GetLastRecoveredDates(arg0 string, arg1 int) []time.Time {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetLastRecoveredDates", arg0, arg1)
        ret0, _ := ret[0].([]time.Time)
        return ret0
}

// GetLastRecoveredDates indicates an expected call of GetLastRecoveredDates.
func (mr *MockRepositoryMockRecorder) GetLastRecoveredDates(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRecoveredDates", reflect.TypeOf((*MockRepository)(nil).GetLastRecoveredDates), arg0, arg1)
}