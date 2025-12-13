package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"scorecard/models"
)

type SectionRepository struct {}





func (SectionRepository) CreateSection(section models.Section) error {
	fmt.Println("sectionDao recieves input: ", section)
	sectionDataJson, err := json.Marshal(section.SectionData)
	if err != nil {
		return fmt.Errorf("failed to marshal section data: %w", err)
	}
	
	query := "INSERT INTO sections (scorecard_id, section_title, section_data, total_score) VALUES (?, ?, ?, ?)"


	//execute query
	res, err := DB.Db.Exec(query,section.ScorecardId, section.SectionTitle, sectionDataJson, section.TotalScore)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("failed to create section %s into database: %w", section.SectionTitle, err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last insert id: %w", err)
	}

	fmt.Printf("Section '%s' successfully created in in db with Id %d.\n", section.SectionTitle, lastId)
	log.Printf("Created section with ID: %d", lastId)
    return nil 
}

func  (SectionRepository)  GetSectionById(scorecardId int, sectionId int) (models.Section, error) {
	var section models.Section
	var sectionDataRaw []byte
	query := "SELECT  id, scorecard_id, section_title, section_data,  total_score, created_at, updated_at FROM sections WHERE id=? and scorecard_id=?"

	err := DB.Db.QueryRow(query, sectionId, scorecardId).Scan(&section.Id, &section.ScorecardId, &section.SectionTitle, &sectionDataRaw, &section.TotalScore, &section.CreatedAt, &section.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows{
		
			return section, fmt.Errorf("no section found scoredcardId:%d and sectionId:%d", scorecardId, sectionId)
		}
        fmt.Println(err)
		return section, fmt.Errorf("failed to get scorecard with  Id %d: %w", section.Id, err)
	}

	//unmarshal
	if err = json.Unmarshal(sectionDataRaw, &section.SectionData); err != nil {
		return section, fmt.Errorf("failed to unmarshal section data: %w", err)
	}

	return section, nil
}

func (SectionRepository)GetAllSections() ([]models.Section, error) {

	query := "SELECT  id, scorecard_id, section_title, section_data,  total_score, created_at, updated_at FROM sections"	
	sections := []models.Section{}

	rows, err := DB.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var section models.Section
		var sectionDataRaw []byte

		if err := rows.Scan(&section.Id, &section.ScorecardId, &section.SectionTitle, &sectionDataRaw, &section.TotalScore, &section.CreatedAt, &section.UpdatedAt); err != nil {
			return nil, err
		}
		  
		//unmarshal
		if err = json.Unmarshal(sectionDataRaw, &section.SectionData); err != nil {
		return sections, fmt.Errorf("failed to unmarshal section data: %w", err)
		}
	

		sections = append(sections, section)
		}

	return sections, nil
}

func (SectionRepository) UpdateSection(section models.Section) error {
	sectionDataJson, err := json.Marshal(section.SectionData)
	if err != nil {
		return fmt.Errorf("failed to marshal section data: %w", err)
	}
	query := "UPDATE sections SET scorecard_id=?, section_title=?, section_data=?, total_score=? WHERE id=?"
	res, err := DB.Db.Exec(query, section.ScorecardId, section.SectionTitle, sectionDataJson, section.TotalScore, section.Id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update section Id %d: %w", section.Id, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Updated user ID %d, rows affected: %d", section.Id, rowsAffected)
	return nil
}


func (SectionRepository) DeleteSection(sectionId int) error {
	query := "DELETE FROM sections WHERE id=?"
	res, err := DB.Db.Exec(query, sectionId)
	if err != nil {
		return fmt.Errorf("failed to delete section with  Id %d: %w", sectionId, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Deleted section Id %d, rows affected: %d", sectionId, rowsAffected)
	return nil
}