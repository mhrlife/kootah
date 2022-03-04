// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	domain "kootah/domain"

	mock "github.com/stretchr/testify/mock"
)

// UIDRepository is an autogenerated mock type for the UIDRepository type
type UIDRepository struct {
	mock.Mock
}

// GetUID provides a mock function with given fields:
func (_m *UIDRepository) GetUID() (domain.UID, error) {
	ret := _m.Called()

	var r0 domain.UID
	if rf, ok := ret.Get(0).(func() domain.UID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.UID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}