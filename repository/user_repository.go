package repository

import (
	"errors"
	"reflect"

	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/db/query"
	"github.com/UxiT/rdp/domain"
)

type userRepository struct {
	database   db.Database
	collection string
}

func NewUserRepository(db db.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(user *domain.User) error {
	var fields = []string{"name", "last_name", "login", "password"}
	var values = []string{user.Name, user.Last_Name, user.Login, user.Password}

	data := make([]interface{}, len(values))

	for i, s := range values {
		data[i] = s
	}

	err := ur.database.InsertOne(fields, data, "users")

	return err
}

func (ur *userRepository) GetByField(column string, value string) ([]domain.User, error) {
	var userByField []domain.User

	builder := query.NewBuilder("users")
	builder.Where(column, "=", value)
	builder.Read()

	userInterfaces, err := ur.database.GetByQuery(*builder.GetQuery(), reflect.TypeOf(domain.User{}))

	for _, c := range userInterfaces {
		course, ok := c.(domain.User)

		if !ok {
			return nil, errors.New("UserInterface does not contain User struct")
		} else {
			userByField = append(userByField, course)
		}
	}

	return userByField, err

}
