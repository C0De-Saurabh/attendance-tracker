package models

type Subject struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Credits int    `json:"credits"`
}
