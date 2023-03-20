package entities

import "time"

type Recipe struct {
	Id            int       `json:"id"`
	Creation_date time.Time `json:"creation_date"`
	User_id       int       `json:"user_id"`
	Name          string    `json:"name"`
	Total_steps   int       `json:"total_steps"`
	Nutriscore    int       `json:"nutriscore"`
	Score         float32   `json:"score"`
	Icon          []byte    `json:"image"`
	Kcal          int       `json:"kcal"`
}
