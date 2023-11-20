package port

import (
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
)

type MealInterface interface {
	SaveMeal(meal request.Meal) (*response.Meal, error)
	GetMeals(userEmail string) ([]response.Meal, error)
	DeleteMeal(mealID int) error
	ConsumeMeal(meal response.Meal) ([]mapping.ConsumedProduct, error)
}
