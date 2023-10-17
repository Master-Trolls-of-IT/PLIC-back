package port

import (
	"gaia-api/infrastructure/model/requests/meal"
)

type MealInterface interface {
	SaveMeal(meal meal.Meal) error
}
