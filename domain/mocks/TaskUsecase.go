// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import domain "github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
import mock "github.com/stretchr/testify/mock"

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, task
func (_m *TaskUsecase) Create(c context.Context, task *domain.Task) error {
	ret := _m.Called(c, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByUserID provides a mock function with given fields: c, userID
func (_m *TaskUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ret := _m.Called(c, userID)

	var r0 []domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Task); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
