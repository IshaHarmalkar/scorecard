package database

import (
	"database/sql"
	"fmt"
	"log"
	"scorecard/models"
)

type ScorecardRepository struct {}





func (ScorecardRepository) CreateScorecard(scorecard models.Scorecard) error {
	query := "INSERT INTO scorecards (user_id, title, url, total_score) VALUES (?, ?, ?, ?)"

	//execute query
	res, err := DB.Db.Exec(query, scorecard.UserId, scorecard.Title, scorecard.Url, scorecard.TotalScore)
	if err != nil {
		return fmt.Errorf("failed to create scorecard %s into database: %w", scorecard.Title, err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last insert id: %w", err)
	}

	fmt.Printf("Scorecard '%s' successfully created in in db with Id %d.\n", scorecard.Title, lastId)
	log.Printf("Created scorecard with ID: %d", lastId)
    return nil 
}

func  (ScorecardRepository)  GetScorecardById(scorecardId int) (models.Scorecard, error) {
	var scorecard  models.Scorecard
	query := "SELECT  id, user_id, title, url, total_score, created_at, updated_at FROM scorecards WHERE id=?"
	err := DB.Db.QueryRow(query, scorecardId).Scan(&scorecard.Id, &scorecard.UserId, &scorecard.Title, &scorecard.Url, &scorecard.TotalScore, &scorecard.CreatedAt, &scorecard.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows{
		
			return scorecard, fmt.Errorf("no scorecard found for id:%d", scorecardId)
		}
        fmt.Println(err)
		return scorecard, fmt.Errorf("failed to get scorecard with  Id %d: %w", scorecard.Id, err)
	}

	return scorecard, nil
}

func (ScorecardRepository)GetAllScorecards() ([]models.Scorecard, error) {
	query := "SELECT  id, user_id, title, url, total_score, created_at, updated_at FROM scorecards"
	scorecards := []models.Scorecard{}

	rows, err := DB.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var scorecard models.Scorecard
		if err := rows.Scan(&scorecard.Id, &scorecard.UserId, &scorecard.Title, &scorecard.Url, &scorecard.TotalScore, &scorecard.CreatedAt, &scorecard.UpdatedAt); err != nil {
			return nil, err
		}
		scorecards = append(scorecards, scorecard)
	}

	return scorecards, nil
}

func (ScorecardRepository) UpdateScorecard(scorecard models.Scorecard) error {
	query := "UPDATE scorecardss SET title=?, total_score=? WHERE id=?"
	res, err := DB.Db.Exec(query, scorecard.Title, scorecard.TotalScore)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update user Id %d: %w", scorecard.Id, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Updated user ID %d, rows affected: %d", scorecard.Id, rowsAffected)
	return nil
}





func (ScorecardRepository) DeleteScorecard(scorecardId int) error {
	query := "DELETE FROM scorecards WHERE id=?"
	res, err := DB.Db.Exec(query, scorecardId)
	if err != nil {
		return fmt.Errorf("failed to delete scorecard with  Id %d: %w", scorecardId, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Deleted user Id %d, rows affected: %d", scorecardId, rowsAffected)
	return nil
}