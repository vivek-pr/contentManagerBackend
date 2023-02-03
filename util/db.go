package util

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=root host=localhost port=5432 dbname=machax sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
