package services

import (
	"errors"
	"scorecard/database"
	"scorecard/models"
)

type ScorecardService struct {
	Repo database.ScorecardRepository
}



func (s ScorecardService) CreateForm(user models.User) error {
	if user.Email == "" || user.Name == "" {
		return errors.New("name and email required")
	}

	return  s.Repo.CreateForm(user)
}

func (s ScorecardService) GetForm(id int) (models.User, error){
	return s.Repo.GetFormById(id)
}

func (s ScorecardService) GetAllForms() ([]models.User, error){
	return s.Repo.GetAllForms()
}


func (s ScorecardService) UpdateForm(user models.User) ( error){
	return s.Repo.UpdateForm(user)
}

func (s ScorecardService)DeleteForm(id int) error {
	return s.Repo.DeleteForm(id)
}
