package main

import "time"

type BoardingInfo struct {
	EncryptedCardID string    `json:"encrypted_card_id"`
	BoardingLine    *int      `json:"boarding_line"`
	BoardingStation *int      `json:"boarding_station"`
	BoardingTime    time.Time `json:"boarding_time"`
}

type AlightingInfo struct {
	EncryptedCardID  string    `json:"encrypted_card_id"`
	AlightingLine    *int      `json:"alighting_line"`
	AlightingStation *int      `json:"alighting_station"`
	AlightingTime    time.Time `json:"alighting_time"`
}

type LineDetails struct {
	LineID     *int   `json:"line_id"`
	LineDetail string `json:"line_detail"`
}

type StationDetails struct {
	LineID        *int   `json:"line_id"`
	StationID     *int   `json:"station_id"`
	StationDetail string `json:"station_detail"`
}
