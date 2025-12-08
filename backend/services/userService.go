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

func (s UserService) GetUser(id int) (models.User, error){
	return s.Repo.GetUserById(id)
}

func (s UserService) GetAllUsers() ([]models.User, error){
	return s.Repo.GetAllUsers()
}


func (s UserService) UpdateUser(user models.User) ( error){
	return s.Repo.UpdateUser(user)
}

func (s UserService)DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
