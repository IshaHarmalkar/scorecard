package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	hostname = "localhost"
	dbname   = "scorecard"
)

func Dsn(dbName string) string {

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

}

func Connect(){
	db, err := sql.Open("mysql", Dsn(""))
	if err != nil {
		log.Panicf("Error %s when opening DB\n", err)
		return

	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Panicf("Error %s when creating DB\n", err)
		return
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Panicf("Error %s when fetching rows", err)
		return
		
	}

	log.Printf("rows affected: %d\n", no)
	db.Close()

	db, err = sql.Open("mysql", Dsn(dbname))
	if err != nil {
		log.Panicf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(20)
	db.SetConnMaxIdleTime(time.Minute * 5)


	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pining DB", err)
		return
	}

	log.Printf("Connected to DB %s successfully\n", dbname)
}