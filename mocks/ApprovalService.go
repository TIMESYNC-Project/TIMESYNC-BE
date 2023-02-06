// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	approval "timesync-be/features/approval"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// ApprovalService is an autogenerated mock type for the ApprovalService type
type ApprovalService struct {
	mock.Mock
}

// GetApproval provides a mock function with given fields:
func (_m *ApprovalService) GetApproval() ([]approval.Core, error) {
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

// PostApproval provides a mock function with given fields: token, fileData, newApproval
func (_m *ApprovalService) PostApproval(token interface{}, fileData multipart.FileHeader, newApproval approval.Core) (approval.Core, error) {
	ret := _m.Called(token, fileData, newApproval)

	var r0 approval.Core
	if rf, ok := ret.Get(0).(func(interface{}, multipart.FileHeader, approval.Core) approval.Core); ok {
		r0 = rf(token, fileData, newApproval)
	} else {
		r0 = ret.Get(0).(approval.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, multipart.FileHeader, approval.Core) error); ok {
		r1 = rf(token, fileData, newApproval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApproval provides a mock function with given fields: token, approvalID, updatedApproval
func (_m *ApprovalService) UpdateApproval(token interface{}, approvalID uint, updatedApproval approval.Core) (approval.Core, error) {
	ret := _m.Called(token, approvalID, updatedApproval)

	var r0 approval.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, approval.Core) approval.Core); ok {
		r0 = rf(token, approvalID, updatedApproval)
	} else {
		r0 = ret.Get(0).(approval.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, approval.Core) error); ok {
		r1 = rf(token, approvalID, updatedApproval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApprovalService interface {
	mock.TestingT
	Cleanup(func())
}

// NewApprovalService creates a new instance of ApprovalService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApprovalService(t mockConstructorTestingTNewApprovalService) *ApprovalService {
	mock := &ApprovalService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
