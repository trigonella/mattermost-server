// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/trigonella/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// EmojiStore is an autogenerated mock type for the EmojiStore type
type EmojiStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: emoji, time
func (_m *EmojiStore) Delete(emoji *model.Emoji, time int64) error {
	ret := _m.Called(emoji, time)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Emoji, int64) error); ok {
		r0 = rf(emoji, time)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id, allowFromCache
func (_m *EmojiStore) Get(id string, allowFromCache bool) (*model.Emoji, error) {
	ret := _m.Called(id, allowFromCache)

	var r0 *model.Emoji
	if rf, ok := ret.Get(0).(func(string, bool) *model.Emoji); ok {
		r0 = rf(id, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(id, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name, allowFromCache
func (_m *EmojiStore) GetByName(name string, allowFromCache bool) (*model.Emoji, error) {
	ret := _m.Called(name, allowFromCache)

	var r0 *model.Emoji
	if rf, ok := ret.Get(0).(func(string, bool) *model.Emoji); ok {
		r0 = rf(name, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(name, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: offset, limit, sort
func (_m *EmojiStore) GetList(offset int, limit int, sort string) ([]*model.Emoji, error) {
	ret := _m.Called(offset, limit, sort)

	var r0 []*model.Emoji
	if rf, ok := ret.Get(0).(func(int, int, string) []*model.Emoji); ok {
		r0 = rf(offset, limit, sort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(offset, limit, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMultipleByName provides a mock function with given fields: names
func (_m *EmojiStore) GetMultipleByName(names []string) ([]*model.Emoji, error) {
	ret := _m.Called(names)

	var r0 []*model.Emoji
	if rf, ok := ret.Get(0).(func([]string) []*model.Emoji); ok {
		r0 = rf(names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(names)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: emoji
func (_m *EmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {
	ret := _m.Called(emoji)

	var r0 *model.Emoji
	if rf, ok := ret.Get(0).(func(*model.Emoji) *model.Emoji); ok {
		r0 = rf(emoji)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Emoji) error); ok {
		r1 = rf(emoji)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: name, prefixOnly, limit
func (_m *EmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {
	ret := _m.Called(name, prefixOnly, limit)

	var r0 []*model.Emoji
	if rf, ok := ret.Get(0).(func(string, bool, int) []*model.Emoji); ok {
		r0 = rf(name, prefixOnly, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, int) error); ok {
		r1 = rf(name, prefixOnly, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
