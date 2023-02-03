package repository

import (
	"contentManagerBackend/util"
)

func SaveUser(username, email, password string) error {
	// Connect to the database
	db, err := util.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Hash the password
	hashedPassword := util.HashPassword(password)

	// Execute the insert statement
	_, err = stmt.Exec(username, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
