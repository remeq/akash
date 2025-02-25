// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	fields "k8s.io/apimachinery/pkg/fields"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	runtime "k8s.io/apimachinery/pkg/runtime"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/core/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// EventInterface is an autogenerated mock type for the EventInterface type
type EventInterface struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Apply(ctx context.Context, event *v1.EventApplyConfiguration, opts metav1.ApplyOptions) (*corev1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(context.Context, *v1.EventApplyConfiguration, metav1.ApplyOptions) *corev1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.EventApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Create(ctx context.Context, event *corev1.Event, opts metav1.CreateOptions) (*corev1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Event, metav1.CreateOptions) *corev1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Event, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWithEventNamespace provides a mock function with given fields: event
func (_m *EventInterface) CreateWithEventNamespace(event *corev1.Event) (*corev1.Event, error) {
	ret := _m.Called(event)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(*corev1.Event) *corev1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*corev1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *EventInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *EventInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *EventInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Event, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *corev1.Event); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
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

// GetFieldSelector provides a mock function with given fields: involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID
func (_m *EventInterface) GetFieldSelector(involvedObjectName *string, involvedObjectNamespace *string, involvedObjectKind *string, involvedObjectUID *string) fields.Selector {
	ret := _m.Called(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID)

	var r0 fields.Selector
	if rf, ok := ret.Get(0).(func(*string, *string, *string, *string) fields.Selector); ok {
		r0 = rf(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fields.Selector)
		}
	}

	return r0
}

// List provides a mock function with given fields: ctx, opts
func (_m *EventInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.EventList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *corev1.EventList
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *corev1.EventList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.EventList)
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
func (_m *EventInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1.Event, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *corev1.Event); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
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

// PatchWithEventNamespace provides a mock function with given fields: event, data
func (_m *EventInterface) PatchWithEventNamespace(event *corev1.Event, data []byte) (*corev1.Event, error) {
	ret := _m.Called(event, data)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(*corev1.Event, []byte) *corev1.Event); ok {
		r0 = rf(event, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*corev1.Event, []byte) error); ok {
		r1 = rf(event, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: scheme, objOrRef
func (_m *EventInterface) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*corev1.EventList, error) {
	ret := _m.Called(scheme, objOrRef)

	var r0 *corev1.EventList
	if rf, ok := ret.Get(0).(func(*runtime.Scheme, runtime.Object) *corev1.EventList); ok {
		r0 = rf(scheme, objOrRef)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.EventList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*runtime.Scheme, runtime.Object) error); ok {
		r1 = rf(scheme, objOrRef)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Update(ctx context.Context, event *corev1.Event, opts metav1.UpdateOptions) (*corev1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(context.Context, *corev1.Event, metav1.UpdateOptions) *corev1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *corev1.Event, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWithEventNamespace provides a mock function with given fields: event
func (_m *EventInterface) UpdateWithEventNamespace(event *corev1.Event) (*corev1.Event, error) {
	ret := _m.Called(event)

	var r0 *corev1.Event
	if rf, ok := ret.Get(0).(func(*corev1.Event) *corev1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*corev1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*corev1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *EventInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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
