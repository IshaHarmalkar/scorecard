package database

import (
	"database/sql"
	"fmt"
	"log"
)

type DbPointer struct {
	Db *sql.DB
}

var DB DbPointer


func Connect() {
	dsn := "root@tcp(localhost)/scorecard?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Cannot open DB: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot ping DB: ", err)
	}

	DB = DbPointer{Db: db}

	fmt.Println("Connected to MySql db successfully.")



}