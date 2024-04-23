package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func postLineDetail(w http.ResponseWriter, r *http.Request) {
	var info LineDetails

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if info.LineID == nil || info.LineDetail == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE Lines
	SET line_detail = $2
	WHERE line_id = $1`
	_, err := Db.Exec(query, *info.LineID, info.LineDetail)
	if err != nil {
		http.Error(w, "Error updating data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Line detail updated successfully")); err != nil {
		log.Printf("Error occured while writing response: %v", err)
	}
}
