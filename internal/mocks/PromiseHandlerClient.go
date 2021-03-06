// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	promise "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server/protocol"
)

// PromiseHandlerClient is an autogenerated mock type for the PromiseHandlerClient type
type PromiseHandlerClient struct {
	mock.Mock
}

// CreatePromise provides a mock function with given fields: ctx, in, opts
func (_m *PromiseHandlerClient) CreatePromise(ctx context.Context, in *promise.CreateRequest, opts ...grpc.CallOption) (*promise.Promise, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(context.Context, *promise.CreateRequest, ...grpc.CallOption) *promise.Promise); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *promise.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribePromise provides a mock function with given fields: ctx, in, opts
func (_m *PromiseHandlerClient) DescribePromise(ctx context.Context, in *promise.UUID, opts ...grpc.CallOption) (*promise.Promise, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(context.Context, *promise.UUID, ...grpc.CallOption) *promise.Promise); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *promise.UUID, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPromises provides a mock function with given fields: ctx, in, opts
func (_m *PromiseHandlerClient) ListPromises(ctx context.Context, in *promise.ListPromisesRequest, opts ...grpc.CallOption) (*promise.ListPromisesRequestResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *promise.ListPromisesRequestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *promise.ListPromisesRequest, ...grpc.CallOption) *promise.ListPromisesRequestResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.ListPromisesRequestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *promise.ListPromisesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePromise provides a mock function with given fields: ctx, in, opts
func (_m *PromiseHandlerClient) RemovePromise(ctx context.Context, in *promise.UUID, opts ...grpc.CallOption) (*promise.SuccessMessage, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *promise.SuccessMessage
	if rf, ok := ret.Get(0).(func(context.Context, *promise.UUID, ...grpc.CallOption) *promise.SuccessMessage); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.SuccessMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *promise.UUID, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePromise provides a mock function with given fields: ctx, in, opts
func (_m *PromiseHandlerClient) UpdatePromise(ctx context.Context, in *promise.UpdatePromiseRequest, opts ...grpc.CallOption) (*promise.Promise, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(context.Context, *promise.UpdatePromiseRequest, ...grpc.CallOption) *promise.Promise); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *promise.UpdatePromiseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
