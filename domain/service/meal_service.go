package service

import (
	ports "gaia-api/domain/port"
)

type MealService struct {
	MealRepo *ports.MealInterface
}

func NewMealService(mealInterface ports.MealInterface) *MealService {
	return &MealService{
		MealRepo: &mealInterface,
	}
}
