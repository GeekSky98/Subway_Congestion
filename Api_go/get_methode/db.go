package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	database = os.Getenv("POSTGRES_DB")
	user     = os.Getenv("POSTGRES_USER")
	passwd   = os.Getenv("POSTGRES_PASSWORD")
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

func CloseDB() {
	if err := Db.Close(); err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
}
