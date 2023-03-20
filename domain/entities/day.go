package entities

import "time"

type Day struct {
	Id         int       `json:"id"`
	User_id    string    `json:"user_id"`
	Date       time.Time `json:"date"`
	Kcal_count int       `json:"calories"`
	EcoScore   float32   `json:"ecoScore"`
}
