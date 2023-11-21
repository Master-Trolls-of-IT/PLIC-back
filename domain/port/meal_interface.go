package port

import (
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
)

type MealInterface interface {
	SaveMeal(meal request.Meal) (*response.Meal, error)
	GetMeals(userEmail string) ([]response.Meal, error)
	DeleteMeal(mealID int) error
	ConsumeMeal(meal request.Meal) ([]response.ConsumedProduct, error)
}
