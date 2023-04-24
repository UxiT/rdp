package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/db/query"
	"github.com/UxiT/rdp/domain/profile"
	task "github.com/UxiT/rdp/domain/tasks"
)

type userTaskRepository struct {
	database db.Database
}

func NewUserTaskRepository(db db.Database) task.UserTaskModel {
	return &userTaskRepository{
		database: db,
	}
}

func (utr *userTaskRepository) Create(c context.Context) error {
	return nil
}

func (utr *userTaskRepository) GetByCourse(c context.Context, userId string, courseId string) ([]task.UserTask, error) {
	var userTasks []task.UserTask
	studentId, err := getStudentId(utr, userId)

	if err != nil {
		return nil, fmt.Errorf("error matchinf user_id to student: %s", err.Error())
	}

	builder := query.NewBuilder("users_tasks")
	builder.Select([]string{"student_id", "task_id", "score", "max_score", "status", "report", "comment"})
	builder.Join("tasks as t", "t.id", "=", "task_id")
	builder.Where("student_id", "=", fmt.Sprintf("%d", *studentId))
	builder.Where("t.course_id", "=", courseId)
	builder.Read()

	userTaskInerface, err := utr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(task.UserTask{}))

	for _, c := range userTaskInerface {
		task, ok := c.(task.UserTask)

		if !ok {
			return nil, errors.New("userInterface does not contain User struct")
		} else {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks, err
}

func (utr *userTaskRepository) GetById(c context.Context, userId string, taskId string) (task.UserTask, error) {
	var userTask task.UserTask

	studentId, err := getStudentId(utr, userId)

	if err != nil {
		return userTask, fmt.Errorf("error matchinf user_id to student: %s", err.Error())
	}

	builder := query.NewBuilder("users_tasks")
	builder.Select([]string{"student_id", "task_id", "score", "max_score", "status", "report", "comment"})
	builder.Where("task_id", "=", taskId)
	builder.Where("student_id", "=", fmt.Sprintf("%d", *studentId))
	builder.Read()

	taskInterface, err := utr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(task.UserTask{}))

	if err != nil || len(taskInterface) == 0 {
		return userTask, errors.New(err.Error())
	}

	task, ok := taskInterface[0].(task.UserTask)

	if !ok {
		return userTask, errors.New("invalid struct")
	}

	userTask = task

	return userTask, err
}

func getStudentId(utr *userTaskRepository, userId string) (*int64, error) {
	studentBuilder := query.NewBuilder("students")
	studentBuilder.Where("user_id", "=", userId)
	studentBuilder.Read()

	studentInterface, err := utr.database.GetByQuery(*studentBuilder.GetQuery(), reflect.TypeOf(profile.Student{}))

	if err != nil || len(studentInterface) == 0 {
		return nil, errors.New(err.Error())
	}

	student, ok := studentInterface[0].(profile.Student)

	if !ok {
		return nil, errors.New("not a student type")
	} else {
		return &student.Id, nil
	}
}
