package entity

import "time"

type Day struct {
	Id        int       `json:"id"`
	UserId    string    `json:"user_id"`
	Date      time.Time `json:"date"`
	KcalCount int       `json:"calories"`
	EcoScore  float32   `json:"ecoScore"`
}
