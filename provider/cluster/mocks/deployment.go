// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	manifest "github.com/ovrclk/akash/manifest"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ovrclk/akash/x/market/types"
)

// Deployment is an autogenerated mock type for the Deployment type
type Deployment struct {
	mock.Mock
}

// LeaseID provides a mock function with given fields:
func (_m *Deployment) LeaseID() types.LeaseID {
	ret := _m.Called()

	var r0 types.LeaseID
	if rf, ok := ret.Get(0).(func() types.LeaseID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.LeaseID)
	}

	return r0
}

// ManifestGroup provides a mock function with given fields:
func (_m *Deployment) ManifestGroup() manifest.Group {
	ret := _m.Called()

	var r0 manifest.Group
	if rf, ok := ret.Get(0).(func() manifest.Group); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(manifest.Group)
	}

	return r0
}
