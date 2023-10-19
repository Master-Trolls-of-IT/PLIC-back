package port

import (
	request "gaia-api/infrastructure/model/requests/meal"
	response "gaia-api/infrastructure/model/responses/meal"
)

type MealInterface interface {
	SaveMeal(meal request.Meal) error
	GetMeals(userEmail string) ([]response.Meal, error)
}
