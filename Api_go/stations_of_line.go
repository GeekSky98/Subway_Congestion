package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Station struct {
	StationID   int    `json:"station_id"`
	StationName string `json:"station_name"`
	TrainCount  int    `json:"train_count"`
}

func getStationsOfLine(w http.ResponseWriter, r *http.Request) {
	lineID := r.URL.Query().Get("line_id")
	if lineID == "" {
		http.Error(w, "line_id is required", http.StatusBadRequest)
		return
	}

	query := `SELECT station_id, station_name, train_count FROM Stations WHERE line_id = $1`
	rows, err := Db.Query(query, lineID)
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
	if err := json.NewEncoder(w).Encode(stations); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	ConnectDB()
	defer Db.Close()

	http.HandleFunc("/statinsOfLine", getStationsOfLine)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
