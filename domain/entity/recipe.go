package entity

import "time"

type Recipe struct {
	Id           int       `json:"id"`
	CreationDate time.Time `json:"creation_date"`
	UserId       int       `json:"user_id"`
	Name         string    `json:"name"`
	TotalSteps   int       `json:"total_steps"`
	Nutriscore   int       `json:"nutriscore"`
	Score        float32   `json:"score"`
	Icon         []byte    `json:"icon"`
	Kcal         int       `json:"kcal"`
}
