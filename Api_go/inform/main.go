package main

import (
	"log"
	"net/http"
)

func main() {
	ConnectDB()
	defer CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/lines", getLines)
	mux.HandleFunc("/statinsOfLine", getStationsOfLine)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
