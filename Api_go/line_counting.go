package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Counting struct {
	TotalPassengers    int `json:"total_passengers"`
	AlightedPassengers int `json:"alighted_passengers"`
}

func getLineCounting(w http.ResponseWriter, r *http.Request) {
	lineID := r.URL.Query().Get("line_id")
	if lineID == "" {
		http.Error(w, "line_id is required", http.StatusBadRequest)
		return
	}

	recordDate := r.URL.Query().Get("today")
	if recordDate == "" {
		http.Error(w, "record_date is required", http.StatusBadRequest)
		return
	}

	query := `SELECT total_passengers, alighted_passengers FROM LinePassengerCount WHERE line_id=$1 AND record_date=$2`
	rows, err := Db.Query(query, lineID, recordDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var line_count Counting
	if err := rows.Scan(&line_count.TotalPassengers, &line_count.AlightedPassengers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(line_count); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	ConnectDB()
	defer Db.Close()

	http.HandleFunc("/line_counting", getLineCounting)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
