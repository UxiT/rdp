package usecase

import (
	"context"
	"time"

	task "github.com/UxiT/rdp/domain/tasks"
)

type userTaskUsecase struct {
	userTaskModel  task.UserTaskModel
	contextTimeout time.Duration
}

func NewUserTaskUsecase(userTaskModel task.UserTaskModel, timeout time.Duration) task.UserTaskModel {
	return &userTaskUsecase{
		userTaskModel:  userTaskModel,
		contextTimeout: timeout,
	}
}

func (utu *userTaskUsecase) Create(c context.Context) error {
	err := utu.userTaskModel.Create(c)

	return err
}

func (utu *userTaskUsecase) GetByCourse(c context.Context, studentId string, courseId string) ([]task.UserTask, error) {
	return utu.userTaskModel.GetByCourse(c, studentId, courseId)
}

func (utu *userTaskUsecase) GetById(c context.Context, userId string, taskId string) (task.UserTask, error) {
	return utu.userTaskModel.GetById(c, userId, taskId)
}
