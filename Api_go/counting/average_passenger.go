package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @Summary Get Average number of passengers at a particular time
// @Description Look up average pssengers of particular time
// @Tags Average
// @Accept json
// @produce json
// @Param line_id query int true "Line id to look up"
// @Param station_id query int true "Station id to look up"
// @Param today_date query int true "Today's date to look up"
// @Param day_of_week query int true "Day of week to look up"
// @Param hour query int true "Hour to look up"
// @Success 200 {object} StationDayAver "Successfully calculated and retrieved average passenger counts"
// @Failure 400 {string} badRequest "Bad request - Error decoding JSON"
// @Failure 404 {string} notFound "Not found - No data available for the given parameters"
// @Failure 500 {string} serverError "Internal server error - Error querying database or encoding JSON"
// @Router /day_count_average [get]
func getDayCountAverage(w http.ResponseWriter, r *http.Request) {
	var req StationDayAverRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	query := `
	WITH RelevantDays AS (
	    SELECT record_day
	    FROM DateStationCount
	    WHERE EXTRACT(DOW FROM record_day) = $4 AND record_day < $3 AND holiday_check = FALSE
	    ORDER BY record_day DESC
	    LIMIT 4
	)
	SELECT 
	    AVG(total_passengers) AS avg_day,
		AVG(CASE WHEN record_hour = $5 THEN total_passengers END) AS avg_hour
	FROM DateStationCount
	WHERE line_id = $1 AND station_id = $2 AND record_day IN (SELECT record_day FROM RelevantDays)
	`
	var AvgDay, AvgHour sql.NullInt64
	err = Db.QueryRow(query, req.LineID, req.StationID, req.TodayDate, req.DayOfWeek, req.Hour).Scan(&AvgDay, &AvgHour)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error querying database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := StationDayAver{
		LineID:    req.LineID,
		StationID: req.StationID,
		DayAver:   convertSQLNullInt64(AvgDay),
		HourAver:  convertSQLNullInt64(AvgHour),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
