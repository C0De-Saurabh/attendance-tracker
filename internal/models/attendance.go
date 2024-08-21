package models

type Attendance struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	SubjectID int    `json:"subject_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}
