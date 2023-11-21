package mapping

import "encoding/json"

type Product struct {
	ID              int         `db:"id"`
	Name            string      `db:"name"`
	ImageURL        string      `db:"image_url"`
	NutriScoreScore int         `db:"nutriscore_score"`
	NutriScoreGrade string      `db:"nutriscore_grade"`
	EcoScore        string      `db:"ecoscore"`
	Barcode         string      `db:"barcode"`
	IsWater         bool        `db:"iswater"`
	Brand           string      `db:"brand"`
	Quantity        string      `db:"quantity"`
	ServingQuantity json.Number `db:"serving_quantity"`
	ServingSize     string      `db:"serving_size"`
	NutrientsID     int         `db:"nutrients_id"`
}

type Nutrients struct {
	EnergyKj      float64 `db:"energy_kj"`
	EnergyKcal    float64 `db:"energy_kcal"`
	Fat           float64 `db:"fat"`
	SaturatedFat  float64 `db:"saturated_fat"`
	Carbohydrates float64 `db:"carbohydrates"`
	Sugar         float64 `db:"sugar"`
	Fiber         float64 `db:"fiber"`
	Proteins      float64 `db:"proteins"`
	Salt          float64 `db:"salt"`
}
