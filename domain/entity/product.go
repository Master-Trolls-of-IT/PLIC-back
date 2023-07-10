package entity

type Product struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Nutrients  Nutrients  `json:"nutrients"`
	ImageURL   string     `json:"image_url"`
	NutriScore NutriScore `json:"nutriscore"`
	EcoScore   string     `json:"ecoscore"`
	IsWater    bool       `json:"iswater"`
}

type Nutrients struct {
	EnergyKj      float64 `json:"energyKj"`
	EnergyKcal    float64 `json:"energyKcal"`
	Fat           float64 `json:"fat"`
	SaturatedFat  float64 `json:"saturatedFat"`
	Carbohydrates float64 `json:"carbohydrates"`
	Sugar         float64 `json:"sugar"`
	Fiber         float64 `json:"fiber"`
	Proteins      float64 `json:"proteins"`
	Salt          float64 `json:"salt"`
}

type NutriScore struct {
	Score float64 `json:"score"`
	Grade string  `json:"grade"`
}
