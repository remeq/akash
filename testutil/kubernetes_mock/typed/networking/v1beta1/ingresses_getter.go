// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// IngressesGetter is an autogenerated mock type for the IngressesGetter type
type IngressesGetter struct {
	mock.Mock
}

// Ingresses provides a mock function with given fields: namespace
func (_m *IngressesGetter) Ingresses(namespace string) v1beta1.IngressInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.IngressInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.IngressInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.IngressInterface)
		}
	}

	return r0
}
