package users

import (
	"attendance-tracker/internal/db"
	"attendance-tracker/internal/models"
	"errors"
)

func Register(user models.User) error {
	// Check if email already exists
	var existingEmail string
	err := db.DB.QueryRow("SELECT email FROM users WHERE email = $1", user.Email).Scan(&existingEmail)
	if err == nil {
		return errors.New("email already exists")
	}

	// Insert new user
	_, err = db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
