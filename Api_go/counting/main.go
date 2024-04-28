package main

import (
	"log"
	"net/http"
)

// @Title Subway Congestion GO-API
// @Description This service is a Go based API for counting
// @Version 0.0.0
// @Host localhost:8082
// @BasePath /app/v1
func main() {
	ConnectDB()
	defer CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/line_counting", getLineCounting)
	mux.HandleFunc("/station_counting", getStationCounting)
	mux.HandleFunc("/passenger_count_aver", getDayCountAverage)

	log.Fatal(http.ListenAndServe(":8082", mux))
}
