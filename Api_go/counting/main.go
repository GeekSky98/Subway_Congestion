package main

import (
	"log"
	"net/http"
)

func main() {
	ConnectDB()
	defer Db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/line_counting", getLineCounting)
	mux.HandleFunc("/station_counting", getStationCounting)
	mux.HandleFunc("/passenge_count_aver", getDayCountAverage)

	log.Fatal(http.ListenAndServe(":8082", mux))
}
