// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Source: nvidia_gpu_manager.go in package gpu
// Code generated by MockGen. DO NOT EDIT.

// Package gpu is a generated GoMock package.
package gpu

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGPUManager is a mock of GPUManager interface
type MockGPUManager struct {
	ctrl     *gomock.Controller
	recorder *MockGPUManagerMockRecorder
}

// MockGPUManagerMockRecorder is the mock recorder for MockGPUManager
type MockGPUManagerMockRecorder struct {
	mock *MockGPUManager
}

// NewMockGPUManager creates a new mock instance
func NewMockGPUManager(ctrl *gomock.Controller) *MockGPUManager {
	mock := &MockGPUManager{ctrl: ctrl}
	mock.recorder = &MockGPUManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGPUManager) EXPECT() *MockGPUManagerMockRecorder {
	return m.recorder
}

// Setup mocks base method
func (m *MockGPUManager) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup
func (mr *MockGPUManagerMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockGPUManager)(nil).Setup))
}

// Initialize mocks base method
func (m *MockGPUManager) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockGPUManagerMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockGPUManager)(nil).Initialize))
}

// Shutdown mocks base method
func (m *MockGPUManager) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockGPUManagerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockGPUManager)(nil).Shutdown))
}

// GetGPUDeviceIDs mocks base method
func (m *MockGPUManager) GetGPUDeviceIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGPUDeviceIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGPUDeviceIDs indicates an expected call of GetGPUDeviceIDs
func (mr *MockGPUManagerMockRecorder) GetGPUDeviceIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGPUDeviceIDs", reflect.TypeOf((*MockGPUManager)(nil).GetGPUDeviceIDs))
}

// GetDriverVersion mocks base method
func (m *MockGPUManager) GetDriverVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverVersion indicates an expected call of GetDriverVersion
func (mr *MockGPUManagerMockRecorder) GetDriverVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverVersion", reflect.TypeOf((*MockGPUManager)(nil).GetDriverVersion))
}

// DetectGPUDevices mocks base method
func (m *MockGPUManager) DetectGPUDevices() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectGPUDevices")
	ret0, _ := ret[0].(error)
	return ret0
}

// DetectGPUDevices indicates an expected call of DetectGPUDevices
func (mr *MockGPUManagerMockRecorder) DetectGPUDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectGPUDevices", reflect.TypeOf((*MockGPUManager)(nil).DetectGPUDevices))
}

// SaveGPUState mocks base method
func (m *MockGPUManager) SaveGPUState() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveGPUState")
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveGPUState indicates an expected call of SaveGPUState
func (mr *MockGPUManagerMockRecorder) SaveGPUState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveGPUState", reflect.TypeOf((*MockGPUManager)(nil).SaveGPUState))
}
