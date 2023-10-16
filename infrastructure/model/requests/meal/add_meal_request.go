package meal

type MealProduct struct {
	Barcode  string `json:"barcode"`
	Quantity string `json:"quantity"`
}

type Meal struct {
	Title     string        `json:"title"`
	UserEmail string        `json:"email"`
	Products  []MealProduct `json:"products"`
}
