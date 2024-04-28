package main

import (
	"log"
	"net/http"
)

// @Title Subway Congestion GO-API
// @Description This service is a Go based API for post routes
// @Version 0.0.0
// @Host localhost:8080
// @BasePath /app/v1
func main() {
	ConnectDB()
	defer CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/post_boarding", postPassengerBoarding)
	mux.HandleFunc("/post_alighting", postPassengerAlighting)
	mux.HandleFunc("/post_line_detali", postLineDetail)
	mux.HandleFunc("/post_station_detail", postStationDetail)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
