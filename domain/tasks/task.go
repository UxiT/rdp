package task

import (
	"context"
	"database/sql"
)

type Task struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Course_Id   string         `json:"cource_id"`
	Theory_File string         `json:"theory_file"`
	Rdp_Config  string         `json:"rdp_config"`
	Description string         `json:"description_text"`
	Extra_Files sql.NullString `json:"extra_file"`
}

type TaskModel interface {
	Read(c context.Context, userId string, taskId string) (Task, error)
}
