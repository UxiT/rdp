package domain

const (
	CollectionUser = "users"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Last_Name  string `json:"last_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByLogin(login string) (User, error)
	GetById(id string) (User, error)
}
