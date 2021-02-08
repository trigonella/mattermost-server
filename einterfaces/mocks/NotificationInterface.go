// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/trigonella/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// NotificationInterface is an autogenerated mock type for the NotificationInterface type
type NotificationInterface struct {
	mock.Mock
}

// CheckLicense provides a mock function with given fields:
func (_m *NotificationInterface) CheckLicense() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetNotificationMessage provides a mock function with given fields: ack, userId
func (_m *NotificationInterface) GetNotificationMessage(ack *model.PushNotificationAck, userId string) (*model.PushNotification, *model.AppError) {
	ret := _m.Called(ack, userId)

	var r0 *model.PushNotification
	if rf, ok := ret.Get(0).(func(*model.PushNotificationAck, string) *model.PushNotification); ok {
		r0 = rf(ack, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PushNotification)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.PushNotificationAck, string) *model.AppError); ok {
		r1 = rf(ack, userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}
