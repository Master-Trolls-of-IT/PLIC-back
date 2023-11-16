package port

import (
	"gaia-api/domain/entity"
	"gaia-api/domain/entity/requests/mealRequest"
	"gaia-api/domain/entity/responses/meal"
)

type MealInterface interface {
	SaveMeal(meal mealRequest.Meal) (*response.Meal, error)
	GetMeals(userEmail string) ([]response.Meal, error)
	DeleteMeal(mealID int) error
	ConsumeMeal(meal response.Meal) ([]entity.ConsumedProduct, error)
}
