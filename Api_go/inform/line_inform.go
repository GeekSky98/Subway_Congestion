package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Get Line informations
// @Description Query all information about the line
// @Tags Informations
// @Produce json
// @Success 200 {array} Line "Successfully retrieved list of subway lines"
// @Failure 500 {string} string "Internal server error - Error querying database or encoding JSON"
// @Router /lines [get]
func getLines(w http.ResponseWriter, _ *http.Request) {
	query := "SELECT line_id, line_name FROM Lines"
	rows, err := Db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

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
