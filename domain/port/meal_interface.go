package port

import "gaia-api/domain/entity"

type MealInterface interface {
	SaveMeal(meal entity.Meal) error
}
