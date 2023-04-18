package usecase

import (
	"context"
	"time"

	courses "github.com/UxiT/rdp/domain/course"
)

type coursesUsecase struct {
	courseModel    courses.CourseModel
	contextTimeout time.Duration
}

func NewCoursesUsecase(couseModel courses.CourseModel, timeout time.Duration) courses.CourseModel {
	return &coursesUsecase{
		courseModel:    couseModel,
		contextTimeout: timeout,
	}
}

func (cu *coursesUsecase) Create(c context.Context, request courses.CreateCourseRequest) error {
	err := cu.courseModel.Create(c, request)

	return err
}

func (cu *coursesUsecase) GetByID(c context.Context, id string) (courses.Course, error) {
	course, err := cu.courseModel.GetByID(c, id)

	return course, err
}

func (cu *coursesUsecase) GetByGroup(c context.Context, group_id string) ([]courses.Course, error) {
	return cu.courseModel.GetByGroup(c, group_id)
}

func (cu *coursesUsecase) FetchByUser(c context.Context, user_id string) ([]courses.Course, error) {
	return cu.courseModel.FetchByUser(c, user_id)
}
