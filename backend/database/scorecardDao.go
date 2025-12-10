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

func (ScorecardRepository)GetAllForms() ([]models.User, error) {
	query := "SELECT id, uuid, name, email FROM users"
	users := []models.User{}

	rows, err := DB.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var user models.User
		if err := rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ScorecardRepository) UpdateForm(user models.User) error {
	query := "UPDATE users SET name=?, email=? WHERE id=?"
	res, err := DB.Db.Exec(query, user.Name, user.Email, user.Id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update user Id %d: %w", user.Id, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Updated user ID %d, rows affected: %d", user.Id, rowsAffected)
	return nil
}





func (ScorecardRepository) DeleteForm(userId int) error {
	query := "DELETE FROM users WHERE id=?"
	res, err := DB.Db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user Id %d: %w", userId, err)
	}

	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Deleted user Id %d, rows affected: %d", userId, rowsAffected)
	return nil
}