package response

import (
	"gaia-api/domain/entity"
)

type MealProduct struct {
	ProductInfo entity.Product `json:"productInfo"`
	Quantity    int            `json:"quantity"`
}

type MealTag struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Meal struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	UserEmail   string        `json:"email"`
	Products    []MealProduct `json:"products"`
	Tags        []MealTag     `json:"tags"`
	IsFavourite bool          `json:"isFavourite"`
}
