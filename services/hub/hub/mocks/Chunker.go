// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	semver "github.com/Masterminds/semver/v3"
)

// Chunker is an autogenerated mock type for the Chunker type
type Chunker struct {
	mock.Mock
}

// Chunk provides a mock function with given fields: name, ver, fileStorageUrl
func (_m *Chunker) Chunk(name string, ver *semver.Version, fileStorageUrl string) error {
	ret := _m.Called(name, ver, fileStorageUrl)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *semver.Version, string) error); ok {
		r0 = rf(name, ver, fileStorageUrl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}