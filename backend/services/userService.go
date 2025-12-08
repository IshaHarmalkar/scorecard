package services

import (
	"errors"
	"scorecard/database"
	"scorecard/models"
)

type UserService struct {
	Repo database.UserRepository
}

func (s UserService) CreateUser(user models.User) error {
	if user.Email == "" || user.Name == "" {
		return errors.New("name and email required")
	}

	return  s.Repo.CreateUser(user)
}