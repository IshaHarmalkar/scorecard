package services

import (
	"scorecard/database"
	"scorecard/models"
	"scorecard/utils"
)

type ScorecardService struct {
	Repo database.ScorecardRepository
}



func (s ScorecardService) CreateScorecard(scorecard models.Scorecard) error {

	//generate slug for scorecard
	slug, err := utils.GenerateSlug(scorecard.Title)
	if err != nil{
		return err
	}
	scorecard.Url = slug
	

	return  s.Repo.CreateScorecard(scorecard)
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
