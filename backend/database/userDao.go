package database

import (
	"fmt"
	"scorecard/models"
)

func (r *UserRespository)createUser(user models.User){
	query := "INSERT INTO users (name, email, password) VALUES (?, ?)"

	fmt.Println(query)
 
}