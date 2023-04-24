package usecase

import (
	"context"
	"time"

	task "github.com/UxiT/rdp/domain/tasks"
)

type taskUsecase struct {
	taskModel      task.TaskModel
	contextTimeout time.Duration
}

func NewTaskUsecase(taskModel task.TaskModel, timeout time.Duration) task.TaskModel {
	return &taskUsecase{
		taskModel:      taskModel,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Read(c context.Context, userId string, taskId string) (task.Task, error) {
	return tu.taskModel.Read(c, userId, taskId)
}
