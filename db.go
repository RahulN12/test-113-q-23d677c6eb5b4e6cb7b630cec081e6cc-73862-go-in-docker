package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func GetDB() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/db")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot connect to db")
	}
	database = db
	return database, nil
}
