// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/tetrafolium/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// AuditStore is an autogenerated mock type for the AuditStore type
type AuditStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: user_id, offset, limit
func (_m *AuditStore) Get(user_id string, offset int, limit int) (model.Audits, error) {
	ret := _m.Called(user_id, offset, limit)

	var r0 model.Audits
	if rf, ok := ret.Get(0).(func(string, int, int) model.Audits); ok {
		r0 = rf(user_id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Audits)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(user_id, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *AuditStore) PermanentDeleteByUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: audit
func (_m *AuditStore) Save(audit *model.Audit) error {
	ret := _m.Called(audit)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Audit) error); ok {
		r0 = rf(audit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
