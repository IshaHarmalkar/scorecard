package services

import (
	"fmt"
	"scorecard/database"
	"scorecard/models"
)

type SectionService struct {
	Repo database.SectionRepository
}



func (s SectionService) CreateSection(section models.Section) error {
	fmt.Println("sectionService recieves input: ", section)

	return  s.Repo.CreateSection(section)
}

func (s SectionService) GetSection(scorecardId int, sectionId int) (models.Section, error){
	return s.Repo.GetSectionById(scorecardId, sectionId)
}

func (s SectionService) GetAllSections() ([]models.Section, error){
	return s.Repo.GetAllSections()
}


func (s SectionService) UpdateSection(section models.Section) ( error){
	return s.Repo.UpdateSection(section)
}

func (s SectionService)DeleteSection(id int) error {
	return s.Repo.DeleteSection(id)
}
