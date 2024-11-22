// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/nestybox/sysbox-fs/domain"
	mock "github.com/stretchr/testify/mock"
)

// MountServiceIface is an autogenerated mock type for the MountServiceIface type
type MountServiceIface struct {
	mock.Mock
}

// MountHelper provides a mock function with given fields:
func (_m *MountServiceIface) MountHelper() domain.MountHelperIface {
	ret := _m.Called()

	var r0 domain.MountHelperIface
	if rf, ok := ret.Get(0).(func() domain.MountHelperIface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MountHelperIface)
		}
	}

	return r0
}

// NewMountHelper provides a mock function with given fields:
func (_m *MountServiceIface) NewMountHelper() domain.MountHelperIface {
	ret := _m.Called()

	var r0 domain.MountHelperIface
	if rf, ok := ret.Get(0).(func() domain.MountHelperIface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MountHelperIface)
		}
	}

	return r0
}

// NewMountInfoParser provides a mock function with given fields: c, process, launchParser, fetchOptions, fetchInodes
func (_m *MountServiceIface) NewMountInfoParser(c domain.ContainerIface, process domain.ProcessIface, launchParser bool, fetchOptions bool, fetchInodes bool) (domain.MountInfoParserIface, error) {
	ret := _m.Called(c, process, launchParser, fetchOptions, fetchInodes)

	var r0 domain.MountInfoParserIface
	if rf, ok := ret.Get(0).(func(domain.ContainerIface, domain.ProcessIface, bool, bool, bool) domain.MountInfoParserIface); ok {
		r0 = rf(c, process, launchParser, fetchOptions, fetchInodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.MountInfoParserIface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.ContainerIface, domain.ProcessIface, bool, bool, bool) error); ok {
		r1 = rf(c, process, launchParser, fetchOptions, fetchInodes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Setup provides a mock function with given fields: css, hds, prs, nss
func (_m *MountServiceIface) Setup(css domain.ContainerStateServiceIface, hds domain.HandlerServiceIface, prs domain.ProcessServiceIface, nss domain.NSenterServiceIface) {
	_m.Called(css, hds, prs, nss)
}