package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	database = "temp_subway"
	user     = "geeksky"
	passwd   = "geeksky"
)

var Db *sql.DB

func ConnectDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, passwd, database)
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	fmt.Println("Successfully connected to database")
}
