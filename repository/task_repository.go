package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/db/query"
	"github.com/UxiT/rdp/domain"
	"github.com/UxiT/rdp/domain/profile"
	task "github.com/UxiT/rdp/domain/tasks"
)

type taskRepository struct {
	database db.Database
}

func NewTaskRepository(db db.Database) task.TaskModel {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) Read(c context.Context, userId string, taskId string) (task.Task, error) {
	var userTask task.Task

	studentBuilder := query.NewBuilder("students")
	studentBuilder.Where("user_id", "=", userId)
	studentBuilder.Read()

	studentInterface, err := tr.database.GetByQuery(*studentBuilder.GetQuery(), reflect.TypeOf(profile.Student{}))

	if err != nil || len(studentInterface) == 0 {
		return userTask, errors.New(err.Error())
	}

	student, ok := studentInterface[0].(profile.Student)

	if !ok {
		return userTask, errors.New("not a student type")
	}

	if err != nil {
		return userTask, fmt.Errorf("error matchinf user_id to student: %s", err.Error())
	}

	builder := query.NewBuilder("tasks")
	builder.Select([]string{"id", "title", "course_id", "theory_file", "rdp_config", "description_text", "extra_file"})
	builder.Join("users_tasks as ut", "ut.task_id", "=", "id")
	builder.Where("id", "=", taskId)
	builder.Where("ut.student_id", "=", fmt.Sprintf("%d", student.Id))
	builder.Read()

	taskInterface, err := tr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(task.Task{}))

	if err != nil || len(taskInterface) == 0 {
		return userTask, errors.New(err.Error())
	}

	task, ok := taskInterface[0].(task.Task)

	if !ok {
		return userTask, errors.New("invalid struct")
	}

	return task, err
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	var fields = []string{"name", "last_name", "login", "password"}
	var values = []string{}

	data := make([]interface{}, len(values))

	for i, s := range values {
		data[i] = s
	}

	err := tr.database.InsertOne(fields, data, "tasks")

	return err
}
