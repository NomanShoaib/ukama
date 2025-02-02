// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NetworkInfoClient is an autogenerated mock type for the NetworkInfoClient type
type NetworkInfoClient struct {
	mock.Mock
}

// ValidateNetwork provides a mock function with given fields: networkId, orgId
func (_m *NetworkInfoClient) ValidateNetwork(networkId string, orgId string) error {
	ret := _m.Called(networkId, orgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(networkId, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewNetworkInfoClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetworkInfoClient creates a new instance of NetworkInfoClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkInfoClient(t mockConstructorTestingTNewNetworkInfoClient) *NetworkInfoClient {
	mock := &NetworkInfoClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
