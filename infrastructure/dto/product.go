package dto

type ProductInfo struct {
	Name       string `json:"product_name"`
	NameFr     string `json:"product_name_fr"`
	Nutrients  ProductNutrients
	ImageURL   string `json:"image_url"`
	NutriScore NutriScore
}

type ProductNutrients struct {
	EnergyKj      int     `json:"energy-kj_100g"`
	EnergyKcal    int     `json:"energy-kcal_100g"`
	Fat           float64 `json:"fat_100g"`
	SaturatedFat  float64 `json:"saturated-fat_100g"`
	Carbohydrates float64 `json:"carbohydrates_100g"`
	Sugar         float64 `json:"sugars_100g"`
	Fiber         float64 `json:"fiber_100g"`
	Proteins      float64 `json:"proteins_100g"`
	Salt          float64 `json:"salt_100g"`
}

type NutriScore struct {
	Score int    `json:"nutriscore_score"`
	Grade string `json:"nutriscore_grade"`
}
