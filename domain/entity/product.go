package entity

import "encoding/json"

type Product struct {
	ID               int              `json:"id"`
	Brand            string           `json:"brand"`
	Name             string           `json:"name"`
	Nutrients        Nutrients        `json:"nutrients"`
	Nutrients100g    Nutrients100g    `json:"nutrients_100g"`
	NutrientsValue   NutrientsValue   `json:"nutrients_value"`
	NutrientsServing NutrientsServing `json:"nutrients_serving"`
	//TODO: ajouter nutrientsUnit aux données du produit si besoin
	//NutrientsUnit    NutrientsUnit    `json:"nutrients_unit"`
	ImageURL        string      `json:"image_url"`
	NutriScore      NutriScore  `json:"nutriscore"`
	EcoScore        string      `json:"ecoscore"`
	IsWater         bool        `json:"isWater"`
	Quantity        string      `json:"quantity"`
	ServingQuantity json.Number ` json:"serving_quantity"`
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

type Nutrients100g struct {
	Nutrients
}

type NutrientsValue struct {
	Nutrients
}

type NutrientsServing struct {
	Nutrients
}

//TODO: ajouter nutrientsUnit aux données du produit si besoin
//type NutrientsUnit struct {
//	EnergyKj      string `json:"energyKj"`
//	EnergyKcal    string `json:"energyKcal"`
//	Fat           string `json:"fat"`
//	SaturatedFat  string `json:"saturatedFat"`
//	Carbohydrates string `json:"carbohydrates"`
//	Sugar         string `json:"sugar"`
//	Fiber         string `json:"fiber"`
//	Proteins      string `json:"proteins"`
//	Salt          string `json:"salt"`
//}

type NutriScore struct {
	Score float64 `json:"score"`
	Grade string  `json:"grade"`
}
