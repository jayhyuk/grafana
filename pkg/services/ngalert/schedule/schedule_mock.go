// Code generated by mockery v2.12.1. DO NOT EDIT.

package schedule

import (
	context "context"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"

	url "net/url"
)

// FakeScheduleService is an autogenerated mock type for the ScheduleService type
type FakeScheduleService struct {
	mock.Mock
}

// AlertmanagersFor provides a mock function with given fields: orgID
func (_m *FakeScheduleService) AlertmanagersFor(orgID int64) []*url.URL {
	ret := _m.Called(orgID)

	var r0 []*url.URL
	if rf, ok := ret.Get(0).(func(int64) []*url.URL); ok {
		r0 = rf(orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*url.URL)
		}
	}

	return r0
}

// DeleteAlertRule provides a mock function with given fields: key
func (_m *FakeScheduleService) DeleteAlertRule(key models.AlertRuleKey) {
	_m.Called(key)
}

// DroppedAlertmanagersFor provides a mock function with given fields: orgID
func (_m *FakeScheduleService) DroppedAlertmanagersFor(orgID int64) []*url.URL {
	ret := _m.Called(orgID)

	var r0 []*url.URL
	if rf, ok := ret.Get(0).(func(int64) []*url.URL); ok {
		r0 = rf(orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*url.URL)
		}
	}

	return r0
}

// Run provides a mock function with given fields: _a0
func (_m *FakeScheduleService) Run(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAlertRule provides a mock function with given fields: key
func (_m *FakeScheduleService) UpdateAlertRule(key models.AlertRuleKey) {
	_m.Called(key)
}

// evalApplied provides a mock function with given fields: _a0, _a1
func (_m *FakeScheduleService) evalApplied(_a0 models.AlertRuleKey, _a1 time.Time) {
	_m.Called(_a0, _a1)
}

// overrideCfg provides a mock function with given fields: cfg
func (_m *FakeScheduleService) overrideCfg(cfg SchedulerCfg) {
	_m.Called(cfg)
}

// stopApplied provides a mock function with given fields: _a0
func (_m *FakeScheduleService) stopApplied(_a0 models.AlertRuleKey) {
	_m.Called(_a0)
}

// NewFakeScheduleService creates a new instance of FakeScheduleService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeScheduleService(t testing.TB) *FakeScheduleService {
	mock := &FakeScheduleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
