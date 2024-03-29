package response

import "encoding/json"

type Product struct {
	ID              int         `json:"id"`
	Brand           string      `json:"brand"`
	Name            string      `json:"name"`
	Barcode         string      `json:"barcode"`
	Nutrients       Nutrients   `json:"nutrients"`
	ImageURL        string      `json:"image_url"`
	NutriScore      NutriScore  `json:"nutriscore"`
	EcoScore        string      `json:"ecoscore"`
	IsWater         bool        `json:"isWater"`
	Quantity        string      `json:"quantity"`
	ServingQuantity json.Number `json:"serving_quantity"`
	ServingSize     string      `json:"serving_size"`
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
	Score int    `json:"score"`
	Grade string `json:"grade"`
}
