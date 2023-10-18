package response

import (
	"encoding/json"
	"gaia-api/domain/entity"
)

type MealProduct struct {
	ProductInfo entity.Product `json:"productInfo"`
	Quantity    json.Number    `json:"quantity"`
}

type MealTags struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Meal struct {
	Title       string        `json:"title"`
	UserEmail   string        `json:"email"`
	Products    []MealProduct `json:"products"`
	Tags        []MealTags    `json:"tags"`
	IsFavourite bool          `json:"isFavourite"`
}
