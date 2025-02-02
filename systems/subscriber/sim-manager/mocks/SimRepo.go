// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	db "github.com/ukama/ukama/systems/subscriber/sim-manager/pkg/db"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/ukama/ukama/systems/common/uuid"
)

// SimRepo is an autogenerated mock type for the SimRepo type
type SimRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: sim, nestedFunc
func (_m *SimRepo) Add(sim *db.Sim, nestedFunc func(*db.Sim, *gorm.DB) error) error {
	ret := _m.Called(sim, nestedFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.Sim, func(*db.Sim, *gorm.DB) error) error); ok {
		r0 = rf(sim, nestedFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: simID, nestedFunc
func (_m *SimRepo) Delete(simID uuid.UUID, nestedFunc func(uuid.UUID, *gorm.DB) error) error {
	ret := _m.Called(simID, nestedFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, func(uuid.UUID, *gorm.DB) error) error); ok {
		r0 = rf(simID, nestedFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: simID
func (_m *SimRepo) Get(simID uuid.UUID) (*db.Sim, error) {
	ret := _m.Called(simID)

	var r0 *db.Sim
	if rf, ok := ret.Get(0).(func(uuid.UUID) *db.Sim); ok {
		r0 = rf(simID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Sim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(simID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNetwork provides a mock function with given fields: networkID
func (_m *SimRepo) GetByNetwork(networkID uuid.UUID) ([]db.Sim, error) {
	ret := _m.Called(networkID)

	var r0 []db.Sim
	if rf, ok := ret.Get(0).(func(uuid.UUID) []db.Sim); ok {
		r0 = rf(networkID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Sim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(networkID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySubscriber provides a mock function with given fields: subscriberID
func (_m *SimRepo) GetBySubscriber(subscriberID uuid.UUID) ([]db.Sim, error) {
	ret := _m.Called(subscriberID)

	var r0 []db.Sim
	if rf, ok := ret.Get(0).(func(uuid.UUID) []db.Sim); ok {
		r0 = rf(subscriberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Sim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(subscriberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: sim, nestedFunc
func (_m *SimRepo) Update(sim *db.Sim, nestedFunc func(*db.Sim, *gorm.DB) error) error {
	ret := _m.Called(sim, nestedFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.Sim, func(*db.Sim, *gorm.DB) error) error); ok {
		r0 = rf(sim, nestedFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSimRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewSimRepo creates a new instance of SimRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSimRepo(t mockConstructorTestingTNewSimRepo) *SimRepo {
	mock := &SimRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
