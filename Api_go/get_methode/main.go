package main

import (
	"log"
	"net/http"
)

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
