// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// NnsWriter is an autogenerated mock type for the NnsWriter type
type NnsWriter struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, nodeId
func (_m *NnsWriter) Delete(ctx context.Context, nodeId string) error {
	ret := _m.Called(ctx, nodeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nodeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: c, nodeId, ip
func (_m *NnsWriter) Set(c context.Context, nodeId string, ip string) error {
	ret := _m.Called(c, nodeId, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, nodeId, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}