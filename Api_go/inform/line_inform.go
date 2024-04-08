package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func getLines(w http.ResponseWriter, r *http.Request) {
	query := "SELECT line_id, line_name FROM Lines"
	rows, err := Db.Query(query)
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
	if err := json.NewEncoder(w).Encode(lines); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
