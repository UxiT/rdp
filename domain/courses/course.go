package cources

import (
	"context"
	"github.com/UxiT/rdp/db"
)

type Course struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Created_At string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CourseModel interface {
	Create(c context.Context, course Course) error
	GetByGroup(c context.Context, group_id int64) (error, db.Entity)
	FetchByUser(c context.Context, user_id int64) error
}

func (c *Course) GetTable() string {
	return "courses"
}

func (c *Course) GetFillableFields() []string {
	fields := []string{"title"}

	return fields
}

func (c *Course) GetValues() []interface{} {
	data := make([]interface{}, len(g.GetFillableFields()))

	data = append(data, c.Title)

	return data
}
