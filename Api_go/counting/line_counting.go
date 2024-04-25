package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func getLineCounting(w http.ResponseWriter, r *http.Request) {
	var req LineCountingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	query := `SELECT total_passengers, alighted_passengers FROM LinePassengerCount WHERE line_id=$1 AND record_date=$2`
	var response LineCounting
	err = Db.QueryRow(query, req.LineID, req.RecordDate).Scan(&response.TotalPassengers, &response.AlightedPassengers)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "line_id and record_date not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.LineID = req.LineID
	response.RecordDate = req.RecordDate

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
