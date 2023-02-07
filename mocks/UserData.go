// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	user "timesync-be/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// Csv provides a mock function with given fields: newUserBatch
func (_m *UserData) Csv(newUserBatch []user.Core) error {
	ret := _m.Called(newUserBatch)

	var r0 error
	if rf, ok := ret.Get(0).(func([]user.Core) error); ok {
		r0 = rf(newUserBatch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: adminID, employeeID
func (_m *UserData) Delete(adminID uint, employeeID uint) error {
	ret := _m.Called(adminID, employeeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(adminID, employeeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllEmployee provides a mock function with given fields:
func (_m *UserData) GetAllEmployee() ([]user.Core, error) {
	ret := _m.Called()

	var r0 []user.Core
	if rf, ok := ret.Get(0).(func() []user.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
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

// Login provides a mock function with given fields: nip
func (_m *UserData) Login(nip string) (user.Core, error) {
	ret := _m.Called(nip)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(nip)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Profile provides a mock function with given fields: userID
func (_m *UserData) Profile(userID uint) (user.Core, error) {
	ret := _m.Called(userID)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(uint) user.Core); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: adminID, newUser
func (_m *UserData) Register(adminID uint, newUser user.Core) (user.Core, error) {
	ret := _m.Called(adminID, newUser)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(uint, user.Core) user.Core); ok {
		r0 = rf(adminID, newUser)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, user.Core) error); ok {
		r1 = rf(adminID, newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: adminID, quote
func (_m *UserData) Search(adminID uint, quote string) ([]user.Core, error) {
	ret := _m.Called(adminID, quote)

	var r0 []user.Core
	if rf, ok := ret.Get(0).(func(uint, string) []user.Core); ok {
		r0 = rf(adminID, quote)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(adminID, quote)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: employeeID, updateData
func (_m *UserData) Update(employeeID uint, updateData user.Core) (user.Core, error) {
	ret := _m.Called(employeeID, updateData)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(uint, user.Core) user.Core); ok {
		r0 = rf(employeeID, updateData)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, user.Core) error); ok {
		r1 = rf(employeeID, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
