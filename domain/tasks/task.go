package task

import "encoding/json"

type Task struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Course_Id   string          `json:"cource_id"`
	Theory_File string          `json:"theory_file"`
	Rdp_Config  string          `json:"rdp_config"`
	Description string          `json:"description"`
	Extra_Files json.RawMessage `json:"extra_files"`
	Created_At  string          `json:"created_at"`
	Updated_At  string          `json:"updated_at"`
}

type TaskRepository interface {
	Create()
	Update()
	Delete()
	UploadFile()
	FetchByCourse()
}
