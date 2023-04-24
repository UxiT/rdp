package profile

type Student struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	GroupId   int64  `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type StudentModel interface {
	Create()
	Read()
	Update()
	Delete()
}
