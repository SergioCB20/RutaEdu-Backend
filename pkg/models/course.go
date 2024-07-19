package models

type Course struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Credits float64 `json:"credits"`
	Requisites []string `json:"requisites"`
	Level int `json:"level"`
}