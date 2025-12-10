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

func (s ScorecardService) GetScorecard(id int) (models.Scorecard, error){
	return s.Repo.GetScorecardById(id)
}

func (s ScorecardService) GetAllScorecards() ([]models.User, error){
	return s.Repo.GetAllForms()
}


func (s ScorecardService) UpdateScorecard(user models.User) ( error){
	return s.Repo.UpdateForm(user)
}

func (s ScorecardService)DeleteScorecard(id int) error {
	return s.Repo.DeleteForm(id)
}
