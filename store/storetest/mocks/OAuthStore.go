// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/trigonella/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// OAuthStore is an autogenerated mock type for the OAuthStore type
type OAuthStore struct {
	mock.Mock
}

// DeleteApp provides a mock function with given fields: id
func (_m *OAuthStore) DeleteApp(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccessData provides a mock function with given fields: token
func (_m *OAuthStore) GetAccessData(token string) (*model.AccessData, error) {
	ret := _m.Called(token)

	var r0 *model.AccessData
	if rf, ok := ret.Get(0).(func(string) *model.AccessData); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccessDataByRefreshToken provides a mock function with given fields: token
func (_m *OAuthStore) GetAccessDataByRefreshToken(token string) (*model.AccessData, error) {
	ret := _m.Called(token)

	var r0 *model.AccessData
	if rf, ok := ret.Get(0).(func(string) *model.AccessData); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccessDataByUserForApp provides a mock function with given fields: userId, clientId
func (_m *OAuthStore) GetAccessDataByUserForApp(userId string, clientId string) ([]*model.AccessData, error) {
	ret := _m.Called(userId, clientId)

	var r0 []*model.AccessData
	if rf, ok := ret.Get(0).(func(string, string) []*model.AccessData); ok {
		r0 = rf(userId, clientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, clientId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApp provides a mock function with given fields: id
func (_m *OAuthStore) GetApp(id string) (*model.OAuthApp, error) {
	ret := _m.Called(id)

	var r0 *model.OAuthApp
	if rf, ok := ret.Get(0).(func(string) *model.OAuthApp); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppByUser provides a mock function with given fields: userId, offset, limit
func (_m *OAuthStore) GetAppByUser(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	ret := _m.Called(userId, offset, limit)

	var r0 []*model.OAuthApp
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.OAuthApp); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(userId, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetApps provides a mock function with given fields: offset, limit
func (_m *OAuthStore) GetApps(offset int, limit int) ([]*model.OAuthApp, error) {
	ret := _m.Called(offset, limit)

	var r0 []*model.OAuthApp
	if rf, ok := ret.Get(0).(func(int, int) []*model.OAuthApp); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthData provides a mock function with given fields: code
func (_m *OAuthStore) GetAuthData(code string) (*model.AuthData, error) {
	ret := _m.Called(code)

	var r0 *model.AuthData
	if rf, ok := ret.Get(0).(func(string) *model.AuthData); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AuthData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorizedApps provides a mock function with given fields: userId, offset, limit
func (_m *OAuthStore) GetAuthorizedApps(userId string, offset int, limit int) ([]*model.OAuthApp, error) {
	ret := _m.Called(userId, offset, limit)

	var r0 []*model.OAuthApp
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.OAuthApp); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(userId, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreviousAccessData provides a mock function with given fields: userId, clientId
func (_m *OAuthStore) GetPreviousAccessData(userId string, clientId string) (*model.AccessData, error) {
	ret := _m.Called(userId, clientId)

	var r0 *model.AccessData
	if rf, ok := ret.Get(0).(func(string, string) *model.AccessData); ok {
		r0 = rf(userId, clientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, clientId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteAuthDataByUser provides a mock function with given fields: userId
func (_m *OAuthStore) PermanentDeleteAuthDataByUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAccessData provides a mock function with given fields: token
func (_m *OAuthStore) RemoveAccessData(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllAccessData provides a mock function with given fields:
func (_m *OAuthStore) RemoveAllAccessData() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAuthData provides a mock function with given fields: code
func (_m *OAuthStore) RemoveAuthData(code string) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAccessData provides a mock function with given fields: accessData
func (_m *OAuthStore) SaveAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	ret := _m.Called(accessData)

	var r0 *model.AccessData
	if rf, ok := ret.Get(0).(func(*model.AccessData) *model.AccessData); ok {
		r0 = rf(accessData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.AccessData) error); ok {
		r1 = rf(accessData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveApp provides a mock function with given fields: app
func (_m *OAuthStore) SaveApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	ret := _m.Called(app)

	var r0 *model.OAuthApp
	if rf, ok := ret.Get(0).(func(*model.OAuthApp) *model.OAuthApp); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.OAuthApp) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAuthData provides a mock function with given fields: authData
func (_m *OAuthStore) SaveAuthData(authData *model.AuthData) (*model.AuthData, error) {
	ret := _m.Called(authData)

	var r0 *model.AuthData
	if rf, ok := ret.Get(0).(func(*model.AuthData) *model.AuthData); ok {
		r0 = rf(authData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AuthData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.AuthData) error); ok {
		r1 = rf(authData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccessData provides a mock function with given fields: accessData
func (_m *OAuthStore) UpdateAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	ret := _m.Called(accessData)

	var r0 *model.AccessData
	if rf, ok := ret.Get(0).(func(*model.AccessData) *model.AccessData); ok {
		r0 = rf(accessData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.AccessData) error); ok {
		r1 = rf(accessData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApp provides a mock function with given fields: app
func (_m *OAuthStore) UpdateApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	ret := _m.Called(app)

	var r0 *model.OAuthApp
	if rf, ok := ret.Get(0).(func(*model.OAuthApp) *model.OAuthApp); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OAuthApp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.OAuthApp) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
