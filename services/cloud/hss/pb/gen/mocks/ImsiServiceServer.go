// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gen "github.com/ukama/ukama/services/cloud/hss/pb/gen"

	testing "testing"
)

// ImsiServiceServer is an autogenerated mock type for the ImsiServiceServer type
type ImsiServiceServer struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) Add(_a0 context.Context, _a1 *gen.AddImsiRequest) (*gen.AddImsiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.AddImsiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.AddImsiRequest) *gen.AddImsiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.AddImsiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.AddImsiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddGuti provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) AddGuti(_a0 context.Context, _a1 *gen.AddGutiRequest) (*gen.AddGutiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.AddGutiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.AddGutiRequest) *gen.AddGutiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.AddGutiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.AddGutiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) Delete(_a0 context.Context, _a1 *gen.DeleteImsiRequest) (*gen.DeleteImsiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.DeleteImsiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.DeleteImsiRequest) *gen.DeleteImsiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.DeleteImsiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.DeleteImsiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) Get(_a0 context.Context, _a1 *gen.GetImsiRequest) (*gen.GetImsiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.GetImsiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.GetImsiRequest) *gen.GetImsiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.GetImsiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.GetImsiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) Update(_a0 context.Context, _a1 *gen.UpdateImsiRequest) (*gen.UpdateImsiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpdateImsiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpdateImsiRequest) *gen.UpdateImsiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpdateImsiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpdateImsiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTai provides a mock function with given fields: _a0, _a1
func (_m *ImsiServiceServer) UpdateTai(_a0 context.Context, _a1 *gen.UpdateTaiRequest) (*gen.UpdateTaiResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *gen.UpdateTaiResponse
	if rf, ok := ret.Get(0).(func(context.Context, *gen.UpdateTaiRequest) *gen.UpdateTaiResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.UpdateTaiResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gen.UpdateTaiRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedImsiServiceServer provides a mock function with given fields:
func (_m *ImsiServiceServer) mustEmbedUnimplementedImsiServiceServer() {
	_m.Called()
}

// NewImsiServiceServer creates a new instance of ImsiServiceServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewImsiServiceServer(t testing.TB) *ImsiServiceServer {
	mock := &ImsiServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}