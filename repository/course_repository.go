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

func (cr *courseRepository) Create(c context.Context, course courses.Course) error {
	err := cr.database.InsertOne(course.GetFillableFields(), course.GetValues(), course.GetTable())

	return err
}

func (cr *courseRepository) GetByID(c context.Context, courseId string) (courses.Course, error) {
	var courseInterface interface{} = courses.Course{}

	builder := query.NewBuilder("courses")
	builder.Where("id", "=", courseId)

	courseInterface, err := cr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(courses.Course{}))

	course, ok := courseInterface.(courses.Course)

	if !ok {
		fmt.Errorf("userInterface does not contain User struct")
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
	builder.Where("user_id", "=", user_id)

	courseInterfaces, err := cr.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(courses.Course{}))

	for _, c := range courseInterfaces {
		course, ok := c.(courses.Course)

		if !ok {
			fmt.Errorf("UserInterface does not contain User struct")
		} else {
			coursesByUser = append(coursesByUser, course)
		}
	}

	return coursesByUser, err
}
