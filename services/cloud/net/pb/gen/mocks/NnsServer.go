// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gen "github.com/ukama/ukama/services/cloud/net/pb/gen"
)

// NnsServer is an autogenerated mock type for the NnsServer type
type NnsServer struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *NnsServer) Delete(_a0 context.Context, _a1 *gen.DeleteRequest) (*gen.DeleteResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.DeleteResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeleteRequest) *gen.DeleteResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.DeleteResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.DeleteRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *NnsServer) Get(_a0 context.Context, _a1 *gen.GetRequest) (*gen.GetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetRequest) *gen.GetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *NnsServer) List(_a0 context.Context, _a1 *gen.ListRequest) (*gen.ListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.ListRequest) *gen.ListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.ListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *NnsServer) Set(_a0 context.Context, _a1 *gen.SetRequest) (*gen.SetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.SetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.SetRequest) *gen.SetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.SetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.SetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedNnsServer provides a mock function with given fields:
func (_m *NnsServer) mustEmbedUnimplementedNnsServer() {
	_m.Called()
}