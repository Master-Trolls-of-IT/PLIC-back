package mapping

import "time"

type Recipe struct {
	Id           int       `json:"id"`
	CreationDate time.Time `json:"creationDate"`
	UserId       int       `json:"userId"`
	Name         string    `json:"name"`
	TotalSteps   int       `json:"totalSteps"`
	Nutriscore   int       `json:"nutriscore"`
	Score        float32   `json:"score"`
	Icon         []byte    `json:"icon"`
	Kcal         int       `json:"kcal"`
}
