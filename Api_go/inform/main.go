package main

import (
	"log"
	"net/http"
)

// @Title Subway Congestion GO-API
// @Description This service is a Go based API for elementary information
// @Version 0.0.0
// @Host localhost:8081
// @BasePath /app/v1
func main() {
	ConnectDB()
	defer CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/lines", getLines)
	mux.HandleFunc("/statinsOfLine", getStationsOfLine)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
