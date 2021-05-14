// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	"github.com/americanas-go/serverless/cloudevents"
	event "github.com/cloudevents/sdk-go/v2/event"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// After provides a mock function with given fields: ctx, in, out, err
func (_m *Middleware) After(ctx context.Context, in event.Event, out *event.Event, err error) (context.Context, error) {
	ret := _m.Called(ctx, in, out, err)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, event.Event, *event.Event, error) context.Context); ok {
		r0 = rf(ctx, in, out, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, event.Event, *event.Event, error) error); ok {
		r1 = rf(ctx, in, out, err)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AfterAll provides a mock function with given fields: ctx, inout
func (_m *Middleware) AfterAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	ret := _m.Called(ctx, inout)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, []*cloudevents.InOut) context.Context); ok {
		r0 = rf(ctx, inout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*cloudevents.InOut) error); ok {
		r1 = rf(ctx, inout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Before provides a mock function with given fields: ctx, in
func (_m *Middleware) Before(ctx context.Context, in *event.Event) (context.Context, error) {
	ret := _m.Called(ctx, in)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) context.Context); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *event.Event) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeforeAll provides a mock function with given fields: ctx, inout
func (_m *Middleware) BeforeAll(ctx context.Context, inout []*cloudevents.InOut) (context.Context, error) {
	ret := _m.Called(ctx, inout)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, []*cloudevents.InOut) context.Context); ok {
		r0 = rf(ctx, inout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*cloudevents.InOut) error); ok {
		r1 = rf(ctx, inout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
