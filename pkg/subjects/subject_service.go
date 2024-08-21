package subjects

import (
	"attendance-tracker/internal/db"
	"attendance-tracker/internal/models"
)

func Add(subject models.Subject) error {
	_, err := db.DB.Exec("INSERT INTO subjects (name, credits) VALUES ($1, $2)", subject.Name, subject.Credits)
	if err != nil {
		return err
	}

	return nil
}
