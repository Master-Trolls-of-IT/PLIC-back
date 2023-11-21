package service

import (
	ports "gaia-api/domain/port"
)

type OpenFoodFactsService struct {
	ProductRepo *ports.ProductInterface
	MealRepo    *ports.MealInterface
	RecipeRepo  *ports.RecipeInterface
}

func NewOpenFoodFactsService(productRepo ports.ProductInterface, mealRepo ports.MealInterface, recipeRepo ports.RecipeInterface) *OpenFoodFactsService {
	return &OpenFoodFactsService{ProductRepo: &productRepo, MealRepo: &mealRepo, RecipeRepo: &recipeRepo}
}
