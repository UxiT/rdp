package repository

import (
	"context"

	"github.com/UxiT/rdp/db"
	"github.com/UxiT/rdp/domain"
)

type taskRepository struct {
	database   db.Database
	collection string
}

func NewTaskRepository(db db.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	var fields = []string{"name", "last_name", "login", "password"}
	var values = []string{}

	data := make([]interface{}, len(values))

	for i, s := range values {
		data[i] = s
	}

	err := tr.database.InsertOne(fields, data, "tasks")

	return err
}

// func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
// 	collection := tr.database.Collection(tr.collection)

// 	var tasks []domain.Task

// 	idHex, err := primitive.ObjectIDFromHex(userID)
// 	if err != nil {
// 		return tasks, err
// 	}

// 	cursor, err := collection.Find(c, bson.M{"userID": idHex})
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cursor.All(c, &tasks)
// 	if tasks == nil {
// 		return []domain.Task{}, err
// 	}

// 	return tasks, err
// }
