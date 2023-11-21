package mapping

type OpenFoodFactsProduct struct {
	Product struct {
		Brand           string                 `json:"brands"`
		Name            string                 `json:"product_name"`
		Barcode         string                 `json:"code"`
		Nutrients       OpenFoodFactsNutrients `json:"nutriments"`
		ImageURL        string                 `json:"image_front_url"`
		Nutriscore      string                 `json:"nutriscore_grade"`
		Ecoscore        float64                `json:"ecoscore_score"`
		CategoriesTags  []string               `json:"categories_tags"`
		ServingQuantity string                 `json:"serving_quantity"`
	} `json:"product"`
}

type OpenFoodFactsNutrients struct {
	EnergyKj      float64 `json:"energy-kj_100g"`
	EnergyKcal    float64 `json:"energy-kcal_100g"`
	Fat           float64 `json:"fat_100g"`
	SaturatedFat  float64 `json:"saturated-fat_100g"`
	Carbohydrates float64 `json:"carbohydrates_100g"`
	Sugar         float64 `json:"sugars_100g"`
	Fiber         float64 `json:"fiber_100g"`
	Proteins      float64 `json:"proteins_100g"`
	Salt          float64 `json:"salt_100g"`
}
