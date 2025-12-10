package database

import (
	"database/sql"
	"fmt"
	"log"
	"scorecard/models"
)

type ScorecardRepository struct {}





func (ScorecardRepository) CreateForm(user models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	//execute query
	res, err := DB.Db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user %s into database: %w", user.Name, err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch last insert id: %w", err)
	}

	fmt.Printf("User '%s' successfully created in in db with Id %d.\n", user.Name, lastId)
	log.Printf("Inserted user with ID: %d", lastId)
    return nil 
}

func  (ScorecardRepository)  GetFormById(userId int) (models.User, error) {
	var user models.User
	query := "SELECT  id, name, email FROM users WHERE id=?"
	err := DB.Db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows{
		
			return user, fmt.Errorf("no user found for id:%d", userId)
		}
        fmt.Println(err)
		return user, fmt.Errorf("failed to get user Id %d: %w", user.Id, err)
	}

	return user, nil
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