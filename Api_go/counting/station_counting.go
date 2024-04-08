package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func getStationCounting(w http.ResponseWriter, r *http.Request) {
	StationID := r.URL.Query().Get("station_id")
	if StationID == "" {
		http.Error(w, "station_id is required", http.StatusBadRequest)
		return
	}

	recordDate := r.URL.Query().Get("today")
	if recordDate == "" {
		http.Error(w, "record_date is required", http.StatusBadRequest)
		return
	}

	query := `SELECT total_passengers, alighted_passengers FROM StqationPassengerCount WHERE line_id=$1 AND record_date=$2`
	rows, err := Db.Query(query, StationID, recordDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var station_count Counting
	if err := rows.Scan(&station_count.TotalPassengers, &station_count.AlightedPassengers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(station_count); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
