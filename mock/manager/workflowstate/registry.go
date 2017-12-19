// Code generated by MockGen. DO NOT EDIT.
// Source: ../manager/workflowstate/registry.go

// Package mock_workflowstate is a generated GoMock package.
package mock_workflowstate

import (
	model "github.com/Attsun1031/jobnetes/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWorkflowExecutionStateProcessorRegistry is a mock of WorkflowExecutionStateProcessorRegistry interface
type MockWorkflowExecutionStateProcessorRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowExecutionStateProcessorRegistryMockRecorder
}

// MockWorkflowExecutionStateProcessorRegistryMockRecorder is the mock recorder for MockWorkflowExecutionStateProcessorRegistry
type MockWorkflowExecutionStateProcessorRegistryMockRecorder struct {
	mock *MockWorkflowExecutionStateProcessorRegistry
}

// NewMockWorkflowExecutionStateProcessorRegistry creates a new mock instance
func NewMockWorkflowExecutionStateProcessorRegistry(ctrl *gomock.Controller) *MockWorkflowExecutionStateProcessorRegistry {
	mock := &MockWorkflowExecutionStateProcessorRegistry{ctrl: ctrl}
	mock.recorder = &MockWorkflowExecutionStateProcessorRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkflowExecutionStateProcessorRegistry) EXPECT() *MockWorkflowExecutionStateProcessorRegistryMockRecorder {
	return m.recorder
}

// GetProcessor mocks base method
func (m *MockWorkflowExecutionStateProcessorRegistry) GetProcessor(arg0 *model.WorkflowExecution) (WorkflowStateProcessor, error) {
	ret := m.ctrl.Call(m, "GetProcessor", arg0)
	ret0, _ := ret[0].(WorkflowStateProcessor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessor indicates an expected call of GetProcessor
func (mr *MockWorkflowExecutionStateProcessorRegistryMockRecorder) GetProcessor(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessor", reflect.TypeOf((*MockWorkflowExecutionStateProcessorRegistry)(nil).GetProcessor), arg0)
}
