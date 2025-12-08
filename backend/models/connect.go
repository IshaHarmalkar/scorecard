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

func Connect()(*sql.DB, error){
	db, err := sql.Open("mysql", Dsn(""))
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %w", err)

	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		
		return nil, fmt.Errorf("error creating DB: %w", err)
	}
	db.Close()

	

	db, err = sql.Open("mysql", Dsn(dbname))
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %w", err)
	}


	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(20)
	db.SetConnMaxIdleTime(time.Minute * 5)


	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging DB: %w", err)
	}

	log.Printf("Connected to DB %s successfully", dbname)
	return db, nil
}