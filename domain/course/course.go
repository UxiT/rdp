package courses

import (
	"context"
)

type Course struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Created_At string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CourseModel interface {
	Create(c context.Context, course Course) error
	GetByID(c context.Context, id string) (Course, error)
	GetByGroup(c context.Context, group_id string) ([]Course, error)
	FetchByUser(c context.Context, user_id string) ([]Course, error)
}

func (c *Course) GetTable() string {
	return "courses"
}

func (c *Course) GetFillableFields() []string {
	fields := []string{"title"}

	return fields
}

func (c *Course) GetValues() []interface{} {
	data := make([]interface{}, len(c.GetFillableFields()))

	data = append(data, c.Title)

	return data
}
