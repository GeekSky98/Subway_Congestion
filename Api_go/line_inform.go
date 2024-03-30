package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Line struct {
	LineID   int    `json:"line_id"`
	LineName string `json:"line_name"`
}

type Station struct {
	StationID   int    `json:"station_id"`
	StationName string `json:"station_name"`
	TrainCount  int    `json:"train_count"`
}

func getLines(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	query := "SELECT line_id, line_name FROM Lines"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lines []Line
	for rows.Next() {
		var line Line
		if err := rows.Scan(&line.LineID, &line.LineName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lines = append(lines, line)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lines)
}

func getStationsOfLine(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	lineID := r.URL.Query().Get("line_id")
	if lineID == "" {
		http.Error(w, "line_id is required", http.StatusBadRequest)
		return
	}

	query := `SELECT station_id, station_name, train_count FROM Stations WHERE line_id = $1`
	rows, err := db.Query(query, lineID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stations []Station
	for rows.Next() {
		var station Station
		if err := rows.Scan(&station.StationID, &station.StationName, &station.TrainCount); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stations = append(stations, station)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stations)
}

func main() {
	http.HandleFunc("/lines", getLines)
	http.HandleFunc("/statinsOfLine", getStationsOfLine)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
