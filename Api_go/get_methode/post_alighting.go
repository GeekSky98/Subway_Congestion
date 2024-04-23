package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

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
