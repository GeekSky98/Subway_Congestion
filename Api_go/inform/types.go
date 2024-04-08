package main

type Station struct {
	StationID   int    `json:"station_id"`
	StationName string `json:"station_name"`
	TrainCount  int    `json:"train_count"`
}

type Line struct {
	LineID   int    `json:"line_id"`
	LineName string `json:"line_name"`
}