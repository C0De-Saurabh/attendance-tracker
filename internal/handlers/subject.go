package handlers

import (
	"attendance-tracker/internal/models"
	"attendance-tracker/pkg/subjects"
	"encoding/json"
	"net/http"
)

func AddSubject(w http.ResponseWriter, r *http.Request) {
	var subject models.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := subjects.Add(subject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
