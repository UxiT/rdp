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

type CreateCourseRequest struct {
	Title string `json:"title"`
}

type AttachTaskRequest struct {
	CourseId int64   `json:"course_id"`
	TaskIds  []int64 `json:"task_ids"`
}

type CourseModel interface {
	Create(c context.Context, course CreateCourseRequest) error
	GetByID(c context.Context, id string) (Course, error)
	GetByGroup(c context.Context, group_id string) ([]Course, error)
	FetchByUser(c context.Context, user_id string) ([]Course, error)
}

func (c *Course) GetTable() string {
	return "courses"
}

func (c *Course) GetFillableFields(course Course) []*string {
	fields := []*string{&course.Title}

	return fields
}
