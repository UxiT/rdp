package repository

import (
	"context"
	"fmt"

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

	courseInterface, err := cr.database.GetByQuery(*builder.GetQuery(), &courseInterface)

	course, ok := courseInterface.(courses.Course)

	if !ok {
		fmt.Errorf("UserInterface does not contain User struct")
	}

	return course, err
}

func (cr *courseRepository) GetByGroup(c context.Context, group_id string) ([]courses.Course, error) {
	var courseInterface interface{} = courses.Course{}

	builder := query.NewBuilder("courses")
	builder.Where("group_id", "=", group_id)

	courseInterface, err := cr.database.GetByQuery(*builder.GetQuery(), &courseInterface)

	course, ok := courseInterface.(courses.Course)

	if !ok {
		fmt.Errorf("UserInterface does not contain User struct")
	}

	return course, err
}

func (cr *courseRepository) FetchByUser(c context.Context, user_id string) ([]courses.Course, error) {
	return c.Err()
}
