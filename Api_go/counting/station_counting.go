package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Get Number of passenger of the Station
// @Description Check the number of passengers in the current and previous hours
// @Tags Counting
// @Accept json
// @produce json
// @Param line_id query int true "Lien id of target station"
// @Param station_id query int true "station id to look up"
// @Param record_day query string true "Date to look up"
// @Param record_hour query int true "Hour to look up"
// @Success 200 {object} StationCounting "Successfully retrieved passenger counts"
// @Failure 400 {string} badRequest "Bad request - Error decoding JSON"
// @Failure 404 {string} notFound "Not found - Station ID and record details not found"
// @Failure 500 {string} serverError "Internal server error - Error querying database or encoding JSON"
// @Router /station_counting [get]
func getStationCounting(w http.ResponseWriter, r *http.Request) {
	var req StationCountingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	query := `
	SELECT
	  (SELECT total_passengers FROM DateStationCount WHERE station_id=$1 AND record_day=$2 AND record_hour=$3) 
	      AS current_passengers,
	  (SELECT total_passengers FROM DateStationCount WHERE station_id=$1 AND record_day=$2 AND record_hour=$3 - 1) 
	      AS previous_passengers;
	`
	var currentPassengers, previousPassengers sql.NullInt64
	err = Db.QueryRow(query, req.StationID, req.RecordDay, req.RecordHour).Scan(&currentPassengers, &previousPassengers)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "station_id and record_date not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := StationCounting{
		StationID:          req.StationID,
		RecordDay:          req.RecordDay,
		RecordHour:         req.RecordHour,
		ThisTimePassengers: convertSQLNullInt64(currentPassengers),
		PrevPassengers:     convertSQLNullInt64(previousPassengers),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
