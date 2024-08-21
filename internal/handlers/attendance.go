package handlers

import (
	"attendance-tracker/internal/models"
	"attendance-tracker/pkg/attendance"
	"encoding/json"
	"net/http"
)

func TrackAttendance(w http.ResponseWriter, r *http.Request) {
	var record models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := attendance.Track(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
