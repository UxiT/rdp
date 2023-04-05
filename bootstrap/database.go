package bootstrap

import (
	"fmt"
	"log"

	"github.com/UxiT/rdp/db"
)

func NewDatabase(env *Env) db.Database {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/lab?sslmode=disable", dbUser, dbPass, dbHost, dbPort)

	db, err := db.NewDatabase(connString)

	if err != nil {
		log.Fatal(err)
	}

	return *db
}

func closeDatabaseConnection(db *db.Database) {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to DB closed")
}
