package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func postPassengerBoarding(w http.ResponseWriter, r *http.Request) {
	var info BoardingInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if info.EncryptedCardID == "" || info.BoardingLine == nil || info.BoardingStation == nil || info.BoardingTime.IsZero() {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO PassengerInfo (encrypted_card_id, boarding_line, boarding_station, boarding_time) 
		VALUES ($1, $2, $3, $4)`
	_, err := Db.Exec(query, info.EncryptedCardID, *info.BoardingLine, *info.BoardingStation, info.BoardingTime)
	if err != nil {
		http.Error(w, "Error inserting data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Boarding record created successfully")); err != nil {
		log.Fatalf("Error occured while writing response: %v", err)
	}
}
