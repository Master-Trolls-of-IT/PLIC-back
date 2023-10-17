package service

import (
	ports "gaia-api/domain/port"
)

type OpenFoodFactsService struct {
	ProductRepo *ports.ProductInterface
	MealRepo    *ports.MealInterface
}

func NewOpenFoodFactsService(productRepo ports.ProductInterface, mealRepo ports.MealInterface) *OpenFoodFactsService {
	return &OpenFoodFactsService{ProductRepo: &productRepo, MealRepo: &mealRepo}
}
