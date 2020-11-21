// Code generated by mockery v1.0.0. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/api/authentication/v1beta1"
)

// TokenReviewInterface is an autogenerated mock type for the TokenReviewInterface type
type TokenReviewInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, tokenReview, opts
func (_m *TokenReviewInterface) Create(ctx context.Context, tokenReview *v1beta1.TokenReview, opts v1.CreateOptions) (*v1beta1.TokenReview, error) {
	ret := _m.Called(ctx, tokenReview, opts)

	var r0 *v1beta1.TokenReview
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.TokenReview, v1.CreateOptions) *v1beta1.TokenReview); ok {
		r0 = rf(ctx, tokenReview, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.TokenReview)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.TokenReview, v1.CreateOptions) error); ok {
		r1 = rf(ctx, tokenReview, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
