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

func (d *Database) GetByQuery(query query.Query, destType reflect.Type) ([]interface{}, error) {
	rows, err := d.db.Query(query.QueryString, query.Bindings...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Get the slice type for the destination struct type.
	destSliceType := reflect.SliceOf(destType)

	// Create a new slice of the destination type.
	destSlice := reflect.MakeSlice(destSliceType, 0, 0)

	// Loop over the rows and scan each row into a new instance of the destination struct type.
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		pointers := getPointersToFields(dest)
		err := rows.Scan(pointers...)
		if err != nil {
			return nil, err
		}
		destSlice = reflect.Append(destSlice, dest)
	}

	// If there was an error scanning rows, return it.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert the slice of reflect.Values to a slice of interfaces.
	destSliceInterfaces := make([]interface{}, destSlice.Len())
	for i := 0; i < destSlice.Len(); i++ {
		destSliceInterfaces[i] = destSlice.Index(i).Interface()
	}

	return destSliceInterfaces, nil
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
		return fmt.Errorf("error changing table %s", query.Table)
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

func getPointersToFields(value reflect.Value) []interface{} {
	numFields := value.NumField()
	pointers := make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		pointers[i] = value.Field(i).Addr().Interface()
	}

	return pointers
}
