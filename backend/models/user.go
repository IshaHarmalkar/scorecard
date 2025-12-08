package models

//maybe add binding req later

type User struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}