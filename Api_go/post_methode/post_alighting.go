package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Post Alighting Passenger
// @Description Post information of Alighting passenger
// @Tags Post
// @Accept json
// @produce json
// @Param encrypted_card_id query string true "Encrypted Card ID of Passenger"
// @Param alighting_line query int true "Alighting Line"
// @Param alighting_station query int true "Alighting Station"
// @Param alighting_time query string true "Alighting Time - ISO type"
// @Success 201 {string} Posting "Alighting record created successfully"
// @Failure 400 {string} badRequest "Bad request - Error decoding JSON or missing required fields"
// @Failure 500 {string} serverError "Internal server error - Error updating data"
// @Router /post_alighting [post]
func postPassengerAlighting(w http.ResponseWriter, r *http.Request) {
	var info AlightingInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if info.EncryptedCardID == "" || info.AlightingLine == nil || info.AlightingStation == nil || info.AlightingTime.IsZero() {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE PassengerInfo
	SET alighting_line = $2, alighting_station = $3, alighting_time = $4
	WHERE encrypted_card_id = $1 AND boarding_time = (
	    SELECT boarding_time
	    FROM PassengerInfo
	    WHERE encrypted_card_id = $1
	    ORDER BY boarding_time DESC
	    LIMIT 1
	)`
	_, err := Db.Exec(query, info.EncryptedCardID, *info.AlightingLine, *info.AlightingStation, info.AlightingTime)
	if err != nil {
		http.Error(w, "Error updating data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Alighting record created successfully")); err != nil {
		log.Fatalf("Error occured while writing response: %v", err)
	}
}
