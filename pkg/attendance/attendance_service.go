package attendance

import (
	"attendance-tracker/internal/db"
	"attendance-tracker/internal/models"
	"errors"
	"time"
)

func Track(record models.Attendance) error {
	// Check if record already exists
	var existingRecord int
	err := db.DB.QueryRow("SELECT id FROM attendance WHERE user_id = $1 AND subject_id = $2 AND date = $3", record.UserID, record.SubjectID, record.Date).Scan(&existingRecord)
	if err == nil {
		return errors.New("attendance record already exists")
	}

	// Insert new attendance record
	_, err = db.DB.Exec("INSERT INTO attendance (user_id, subject_id, date, status) VALUES ($1, $2, $3, $4)", record.UserID, record.SubjectID, record.Date, record.Status)
	if err != nil {
		return err
	}

	// Optionally, update attendance stats for the user and subject

	return nil
}
