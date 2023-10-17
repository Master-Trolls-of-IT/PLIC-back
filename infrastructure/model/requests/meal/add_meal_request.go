package meal

import "encoding/json"

type MealProduct struct {
	Barcode  string      `json:"barcode"`
	Quantity json.Number `json:"quantity"`
}

type MealTags struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Meal struct {
	Title     string        `json:"title"`
	UserEmail string        `json:"email"`
	Products  []MealProduct `json:"products"`
	Tags      []MealTags    `json:"tags"`
}
