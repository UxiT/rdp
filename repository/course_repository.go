package repository

import (
	"context"

	"github.com/UxiT/rdp/db"
	cources "github.com/UxiT/rdp/domain/courses"
)

type courseRepository struct {
	database   db.Database
	collection string
}

func NewCourseRepository(db db.Database, collection string) cources.CourseModel {
	return &courseRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *courseRepository) Create(c context.Context, course cources.Course) error {
	err := cr.database.InsertOne(course.GetFillableFields(), course.GetValues(), course.GetTable())

	return err
}

func (cr *courseRepository) GetByGroup(c context.Context, group_id int64) (error, db.Entity) {
	entity, err := cr.database.GetByID(group_id, "groups")

	return err, entity
}

func (cr *courseRepository) FetchByUser(c context.Context, user_id int64) error {
	return c.Err()
}
