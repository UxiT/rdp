package task

import (
	"context"
	"reflect"
)

type UserTask struct {
	StudentId int64   `json:"student_id"`
	TaskId    int64   `json:"task_id"`
	Score     int16   `json:"score"`
	MaxScore  int16   `json:"max_score"`
	Status    int8    `json:"status"`
	Report    string  `json:"report"`
	Comment   *string `json:"comment"`
}

type UserTaskModel interface {
	Create(c context.Context) error
	GetByCourse(c context.Context, studentId string, courseId string) ([]UserTask, error)
	GetById(c context.Context, userId string, taskId string) (UserTask, error)
}

func (ut *UserTask) GetFields() []string {
	task := reflect.TypeOf(UserTask{})

	names := make([]string, task.NumField())
	for i := range names {
		names[i] = task.Field(i).Tag.Get("json")
	}

	return names
}
