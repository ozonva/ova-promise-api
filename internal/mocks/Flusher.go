// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ozonva/ova-promise-api/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// Flusher is an autogenerated mock type for the Flusher type
type Flusher struct {
	mock.Mock
}

// Flush provides a mock function with given fields: ctx, promises
func (_m *Flusher) Flush(ctx context.Context, promises []domain.Promise) []domain.Promise {
	ret := _m.Called(ctx, promises)

	var r0 []domain.Promise
	if rf, ok := ret.Get(0).(func(context.Context, []domain.Promise) []domain.Promise); ok {
		r0 = rf(ctx, promises)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Promise)
		}
	}

	return r0
}
