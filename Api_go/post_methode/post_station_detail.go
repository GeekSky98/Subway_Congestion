package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Post Station details
// @Description Post description of a particular station
// @Tags Post
// @Accept json
// @produce json
// @Param line_id query int true "Line id to modify the description"
// @Param station_id query int true "Station id to modify the description"
// @Param station_detail query string true "Description of the station"
// @Success 201 {string} Posting "Station detail record modified successfully"
// @Failure 400 {string} badRequest "Bad request - Error decoding JSON or missing required fields"
// @Failure 500 {string} serverError "Internal server error - Error updating data"
// @Router /post_station_detail [post]
func postStationDetail(w http.ResponseWriter, r *http.Request) {
	var info StationDetails

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if info.LineID == nil || info.StationID == nil || info.StationDetail == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE Stations
	SET station_detail = $3
	WHERE line_id = $1 AND station_id = $2`
	_, err := Db.Exec(query, *info.LineID, *info.StationID, info.StationDetail)
	if err != nil {
		http.Error(w, "Error updating data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Station detail updated successfully")); err != nil {
		log.Printf("Error occured while writing response: %v", err)
	}
}
