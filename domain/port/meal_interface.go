package port

type MealInterface interface {
	SaveMeal(productsBarcodes []string, meal string, email string, isFavourite bool) error
}
