package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

type Entity struct {
	ID        int64
	CreatedAt string
	UpdatedAt string
}

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

func (d *Database) GetRecordByField(table string, field string, value interface{}, record interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, field)
	row := d.db.QueryRow(query, value)

	// Get the struct type and value using reflection
	structType := reflect.TypeOf(record).Elem()
	structValue := reflect.ValueOf(record).Elem()

	// Create pointers to the struct fields
	pointers := make([]interface{}, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		pointers[i] = structValue.FieldByName(field.Name).Addr().Interface()
	}

	// Scan the row and populate the struct fields
	err := row.Scan(pointers...)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("record not found")
		}
		return fmt.Errorf("failed to get record: %v", err)
	}

	return nil
}

// GetByID retrieves a single entity by ID
func (d *Database) GetByID(id int64, table string) (Entity, error) {
	var entity Entity
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", table)
	row := d.db.QueryRow(query, id)
	if err := row.Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity, fmt.Errorf("%s with id %d not found", table, id)
		}
		return entity, err
	}
	return entity, nil
}

// Update updates an existing entity in the database
func (d *Database) Update(entity *Entity, table string) error {
	query := fmt.Sprintf("UPDATE %s SET updated_at=NOW() WHERE id=$1", table)
	result, err := d.db.Exec(query, entity.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s with id %d not found", table, entity.ID)
	}
	return nil
}

// Delete deletes an existing entity from the database
func (d *Database) Delete(entity *Entity, table string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", table)
	result, err := d.db.Exec(query, entity.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s with id %d not found", table, entity.ID)
	}
	return nil
}

func (db *Database) InsertOne(fields []string, data interface{}, table string) (int64, error) {
	query := `INSERT INTO users (%s) VALUES (%s) RETURNING id`

	values := reflect.ValueOf(data).Elem()

	if values.NumField() != len(fields) {
		log.Fatal("Incorect insertion data")
	}

	var placeholders []string
	for i := 0; i < values.NumField(); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	// Construct the final SQL query string with the column names and placeholders
	query = fmt.Sprintf(query, strings.Join(fields, ", "), strings.Join(placeholders, ", "))

	// Execute the query and retrieve the inserted ID
	var id int64
	res, err := db.db.Exec(query, values)

	if err != nil {
		return 0, fmt.Errorf("failed to insert: %v", err)
	}

	id, err = res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("failed to fetch inserted id: %v", err)
	}

	return id, nil
}
