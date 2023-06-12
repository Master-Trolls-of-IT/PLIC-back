package entities

type ProductInfo struct {
	Name       string           `json:"name"`
	Nutrients  ProductNutrients `json:"nutrients"`
	ImageURL   string           `json:"image_url"`
	NutriScore NutriScore       `json:"nutriscore"`
}

type ProductNutrients struct {
	EnergyKj      int     `json:"energyKj"`
	EnergyKcal    int     `json:"energyKcal"`
	Fat           float64 `json:"fat"`
	SaturatedFat  float64 `json:"saturatedFat"`
	Carbohydrates float64 `json:"carbohydrates"`
	Sugar         float64 `json:"sugar"`
	Fiber         float64 `json:"fiber"`
	Proteins      float64 `json:"proteins"`
	Salt          float64 `json:"salt"`
}

type NutriScore struct {
	Score int    `json:"score"`
	Grade string `json:"grade"`
}
