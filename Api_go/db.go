package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	db     = "temp_subway"
	user   = "geeksky"
	passwd = "geeksky"
)

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host = %s user = %s password = %s dbname = %s sslmode = disable", host, user, passwd, db)
	db, err := sql.Open("Postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success connection DB")
	return db
}
