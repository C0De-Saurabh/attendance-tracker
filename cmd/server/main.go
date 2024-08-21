package main

import (
	"attendance-tracker/internal/db"
	"attendance-tracker/internal/handlers"
	"attendance-tracker/internal/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	db.Connect()

	// Connect to Redis
	redis.Connect()

	// Set up router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/subjects", handlers.AddSubject).Methods("POST")
	r.HandleFunc("/attendance", handlers.TrackAttendance).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
