// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FuseServerIface is an autogenerated mock type for the FuseServerIface type
type FuseServerIface struct {
	mock.Mock
}

// Create provides a mock function with given fields:
func (_m *FuseServerIface) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Destroy provides a mock function with given fields:
func (_m *FuseServerIface) Destroy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitWait provides a mock function with given fields:
func (_m *FuseServerIface) InitWait() {
	_m.Called()
}

// IsCntrRegCompleted provides a mock function with given fields:
func (_m *FuseServerIface) IsCntrRegCompleted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MountPoint provides a mock function with given fields:
func (_m *FuseServerIface) MountPoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *FuseServerIface) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCntrRegComplete provides a mock function with given fields:
func (_m *FuseServerIface) SetCntrRegComplete() {
	_m.Called()
}

// Unmount provides a mock function with given fields:
func (_m *FuseServerIface) Unmount() {
	_m.Called()
}