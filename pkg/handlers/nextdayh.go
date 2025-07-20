package handlers

import (
	"log"
	"net/http"
	"time"
)

const (
	// DateFormat constant for date format 20060102
	DateFormat = "20060102"
)

// NextDateHandler handles GET /api/nextdate requests
func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	// Check required parameters
	if dateStr == "" {
		http.Error(w, "parameter 'date' is required", http.StatusBadRequest)
		return
	}

	if repeatStr == "" {
		http.Error(w, "parameter 'repeat' is required", http.StatusBadRequest)
		return
	}

	// Determine current date
	var now time.Time
	var err error

	if nowStr == "" {
		// If now parameter is not specified, use current date
		now = time.Now()
	} else {
		// Parse provided date
		now, err = time.Parse(DateFormat, nowStr)
		if err != nil {
			http.Error(w, "invalid 'now' date format, expected YYYYMMDD", http.StatusBadRequest)
			return
		}
	}

	// Call NextDate function
	nextDate, err := NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return next date
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(nextDate)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
