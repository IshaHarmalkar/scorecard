package database

import (
	"database/sql"
	"fmt"
	"log"
	"scorecard/models"
)

type UserRepository struct {}





func (UserRepository) CreateUser(user models.User) error {
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


func (UserRepository) Update(user models.User) error {
	query := "UPDATE users SET name=?, email=? WHERE id=?"
	res, err := DB.Db.Exec(query, user.Name, user.Email, user.Id)
	if err != nil {
		return fmt.Errorf("failed to update user Id %d: %w", user.Id, err)
	}

	rowsAffected, _ := res.RowsAffected()
	log.Panicf("Updated user ID %d, rows affected: %d", user.Id, rowsAffected)
	return nil
}


func (r *DbPointer) GetUserById(userId int) (models.User, error) {
	var user models.User
	query := "SELECT userName, email, password FROM users WHERE id=?"
	err := DB.Db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows{
			return user, fmt.Errorf("no user found for id:%d", userId)
		}

		return user, fmt.Errorf("failed to get user Id %d: %w", user.Id, err)
	}

	return user, nil
}


func (UserRepository) DeleteUser(userId int) error {
	query := "DELETE FROM users WHERE id=?"
	res, err := DB.Db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user Id %d: %w", userId, err)
	}

	rowsAffected, _ := res.RowsAffected()
	log.Panicf("Deleted user Id %d, rows affected: %d", userId, rowsAffected)
	return nil
}