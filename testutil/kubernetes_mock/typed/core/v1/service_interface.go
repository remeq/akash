// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/api/core/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, service, opts
func (_m *ServiceInterface) Create(ctx context.Context, service *v1.Service, opts metav1.CreateOptions) (*v1.Service, error) {
	ret := _m.Called(ctx, service, opts)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Service, metav1.CreateOptions) *v1.Service); ok {
		r0 = rf(ctx, service, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.Service, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, service, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ServiceInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ServiceInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Service, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *v1.Service); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *ServiceInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *v1.ServiceList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *v1.ServiceList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ServiceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ServiceInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.Service, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *v1.Service); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProxyGet provides a mock function with given fields: scheme, name, port, path, params
func (_m *ServiceInterface) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	ret := _m.Called(scheme, name, port, path, params)

	var r0 rest.ResponseWrapper
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) rest.ResponseWrapper); ok {
		r0 = rf(scheme, name, port, path, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.ResponseWrapper)
		}
	}

	return r0
}

// Update provides a mock function with given fields: ctx, service, opts
func (_m *ServiceInterface) Update(ctx context.Context, service *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	ret := _m.Called(ctx, service, opts)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Service, metav1.UpdateOptions) *v1.Service); ok {
		r0 = rf(ctx, service, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.Service, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, service, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: ctx, service, opts
func (_m *ServiceInterface) UpdateStatus(ctx context.Context, service *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	ret := _m.Called(ctx, service, opts)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Service, metav1.UpdateOptions) *v1.Service); ok {
		r0 = rf(ctx, service, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.Service, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, service, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ServiceInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
