package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/db/query"
	courses "github.com/UxiT/rdp/domain/course"
)

type courseRepository struct {
	database   db.Database
	collection string
}

func NewCourseRepository(db db.Database, collection string) courses.CourseModel {
	return &courseRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *courseRepository) Create(c context.Context, request courses.CreateCourseRequest) error {
	builder := query.NewBuilder("courses")
	var columns [][]string
	columns = append(columns, []string{request.Title})

	builder.Create([]string{"title"}, columns)

	err := cr.database.CreateUpdateDelete(*builder.GetQuery())

	return err
}

func (cr *courseRepository) GetByID(c context.Context, courseId string) (courses.Course, error) {
	var courseInterface interface{} = courses.Course{}

	builder := query.NewBuilder("courses")
	builder.Where("id", "=", courseId)

	courseInterface, err := cr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(courses.Course{}))

	course, ok := courseInterface.(courses.Course)

	if !ok {
		return courses.Course{}, fmt.Errorf("userInterface does not contain User struct")
	}

	return course, err
}

func (cr *courseRepository) GetByGroup(c context.Context, group_id string) ([]courses.Course, error) {
	var coursesByGroup []courses.Course

	builder := query.NewBuilder("courses")
	builder.Where("group_id", "=", group_id)

	courseInterfaces, err := cr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(courses.Course{}))

	for _, c := range courseInterfaces {
		course, ok := c.(courses.Course)

		if !ok {
			return nil, errors.New("userInterface does not contain User struct")
		} else {
			coursesByGroup = append(coursesByGroup, course)
		}
	}

	return coursesByGroup, err
}

func (cr *courseRepository) FetchByUser(c context.Context, user_id string) ([]courses.Course, error) {
	var coursesByUser []courses.Course

	builder := query.NewBuilder("courses")
	builder.Select([]string{"courses.id", "courses.title", "courses.created_at", "courses.updated_at"})
	builder.Join("groups_courses gc", "gc.course_id", "=", "courses.id")
	builder.Join("groups", "groups.id", "=", "gc.group_id")
	builder.Join("students", "students.group_id", "=", "groups.id")
	builder.Where("students.user_id", "=", user_id)
	builder.Read()

	courseInterfaces, err := cr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(courses.Course{}))

	for _, c := range courseInterfaces {
		course, ok := c.(courses.Course)

		if !ok {
			return []courses.Course{}, fmt.Errorf("userInterface does not contain User struct")
		} else {
			coursesByUser = append(coursesByUser, course)
		}
	}

	return coursesByUser, err
}
