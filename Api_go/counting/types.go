package main

type LineCountingRequest struct {
	LineID     int    `json:"line_id"`
	RecordDate string `json:"record_date"`
}

type LineCounting struct {
	LineID             int    `json:"line_id"`
	RecordDate         string `json:"record_date"`
	TotalPassengers    int    `json:"total_passengers"`
	AlightedPassengers int    `json:"alighted_passengers"`
}

type StationCountingRequest struct {
	LineID     int    `json:"line_id"`
	StationID  int    `json:"station_id"`
	RecordDay  string `json:"record_day"`
	RecordHour int    `json:"record_hour"`
}

type StationCounting struct {
	StationID          int    `json:"station_id"`
	RecordDay          string `json:"record_day"`
	RecordHour         int    `json:"record_hour"`
	PrevPassengers     *int   `json:"prev_passengers"`
	ThisTimePassengers *int   `json:"this_time_passengers"`
}

type StationDayAverRequest struct {
	LineID    int    `json:"line_id"`
	StationID int    `json:"station_id"`
	TodayDate string `json:"today_date"`
	DayOfWeek int    `json:"day_of_week"`
	Hour      int    `json:"hour"`
}

type StationDayAver struct {
	LineID    int  `json:"line_id"`
	StationID int  `json:"station_id"`
	DayAver   *int `json:"day_aver"`
	HourAver  *int `json:"hour_aver"`
}
