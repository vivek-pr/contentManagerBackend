package service

import (
	"contentManagerBackend/model"
	"contentManagerBackend/repository"
)

func Signup(user model.User) error {


	// Save the user to the database
	var err error
	err = repository.SaveUser(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

