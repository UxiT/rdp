package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/UxiT/rdp/db/query"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

type Database struct {
	db *sql.DB
}

// Open db connection
func NewDatabase(ConnString string) (*Database, error) {
	db, err := sql.Open("postgres", ConnString)

	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// Close the db connection
func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetByQuery(query query.Query, model *any) (interface{}, error) {
	rows, err := d.db.Query(query.QueryString, query.Bindings...)

	if err != nil {
		return nil, fmt.Errorf("Query error")
	}

	pointers := parseStruct(model)

	err = rows.Scan(pointers...)

	return model, err
}

// Update updates an existing entity in the database
func (d *Database) CreateUpdateDelete(query query.Query) error {
	result, err := d.db.Exec(query.QueryString)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Error changing table %s", query.Table)
	}

	return nil
}

func (db *Database) InsertOne(fields []string, data []interface{}, table string) error {
	query := `INSERT INTO users (%s) VALUES (%s) RETURNING id`

	var placeholders []string
	for i := 0; i < len(fields); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	// Construct the final SQL query string with the column names and placeholders
	query = fmt.Sprintf(query, strings.Join(fields, ", "), strings.Join(placeholders, ", "))

	// Execute the query and retrieve the inserted ID
	_, err := db.db.Exec(query, data...)

	if err != nil {
		return fmt.Errorf("failed to insert: %v", err)
	}

	return nil
}

func parseStruct(model interface{}) []interface{} {
	structType := reflect.TypeOf(model).Elem()
	structValue := reflect.ValueOf(model).Elem()

	// Create pointers to the struct fields
	pointers := make([]interface{}, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		pointers[i] = structValue.FieldByName(field.Name).Addr().Interface()
	}

	return pointers
}
