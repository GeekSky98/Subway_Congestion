package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Get station informations of line
// @Description Query information about all stations on the line
// @Tags Informations
// @Accept json
// @Produce json
// @Param line_id query int true "Line ID for query"
// @Success 200 {array} Station "Successfully retrieved list of subway stations"
// @Failure 500 {string} serverError "Internal server error - Error querying database or encoding JSON"
// @Router /statinsOfLine [get]
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
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

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
