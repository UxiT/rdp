package repository

import (
	"github.com/UxiT/rdp/db"
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
	// var fields = []string{"name", "last_name", "login", "password"}
	// var values =

	_, err := ur.database.InsertOne(fields, values, "users")

	return err
}

func (ur *userRepository) GetByLogin(login string) (domain.User, error) {
	var user domain.User
	err := ur.database.GetRecordByField("users", "login", login, &user)

	return user, err
}

func (ur *userRepository) GetById(login string) (domain.User, error) {
	var user domain.User
	err := ur.database.GetRecordByField("users", "id", login, &user)

	return user, err
}
