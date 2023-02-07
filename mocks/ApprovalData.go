// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	approval "timesync-be/features/approval"

	mock "github.com/stretchr/testify/mock"
)

// ApprovalData is an autogenerated mock type for the ApprovalData type
type ApprovalData struct {
	mock.Mock
}

// ApprovalDetail provides a mock function with given fields: approvalID
func (_m *ApprovalData) ApprovalDetail(approvalID uint) (approval.Core, error) {
	ret := _m.Called(approvalID)

	var r0 approval.Core
	if rf, ok := ret.Get(0).(func(uint) approval.Core); ok {
		r0 = rf(approvalID)
	} else {
		r0 = ret.Get(0).(approval.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(approvalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EmployeeApprovalRecord provides a mock function with given fields: employeeID
func (_m *ApprovalData) EmployeeApprovalRecord(employeeID uint) ([]approval.Core, error) {
	ret := _m.Called(employeeID)

	var r0 []approval.Core
	if rf, ok := ret.Get(0).(func(uint) []approval.Core); ok {
		r0 = rf(employeeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]approval.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(employeeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApproval provides a mock function with given fields:
func (_m *ApprovalData) GetApproval() ([]approval.Core, error) {
	ret := _m.Called()

	var r0 []approval.Core
	if rf, ok := ret.Get(0).(func() []approval.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]approval.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostApproval provides a mock function with given fields: employeeID, newApproval
func (_m *ApprovalData) PostApproval(employeeID uint, newApproval approval.Core) (approval.Core, error) {
	ret := _m.Called(employeeID, newApproval)

	var r0 approval.Core
	if rf, ok := ret.Get(0).(func(uint, approval.Core) approval.Core); ok {
		r0 = rf(employeeID, newApproval)
	} else {
		r0 = ret.Get(0).(approval.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, approval.Core) error); ok {
		r1 = rf(employeeID, newApproval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApproval provides a mock function with given fields: adminID, approvalID, updatedApproval
func (_m *ApprovalData) UpdateApproval(adminID uint, approvalID uint, updatedApproval approval.Core) (approval.Core, error) {
	ret := _m.Called(adminID, approvalID, updatedApproval)

	var r0 approval.Core
	if rf, ok := ret.Get(0).(func(uint, uint, approval.Core) approval.Core); ok {
		r0 = rf(adminID, approvalID, updatedApproval)
	} else {
		r0 = ret.Get(0).(approval.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, approval.Core) error); ok {
		r1 = rf(adminID, approvalID, updatedApproval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApprovalData interface {
	mock.TestingT
	Cleanup(func())
}

// NewApprovalData creates a new instance of ApprovalData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApprovalData(t mockConstructorTestingTNewApprovalData) *ApprovalData {
	mock := &ApprovalData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
