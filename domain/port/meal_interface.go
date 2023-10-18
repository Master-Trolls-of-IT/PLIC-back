package port

import (
	request "gaia-api/infrastructure/model/requests/meal"
	response "gaia-api/infrastructure/model/responses/meal"
)

type MealInterface interface {
	SaveMeal(meal request.Meal) error
	GetMeal(userID int) (response.Meal, error)
}
