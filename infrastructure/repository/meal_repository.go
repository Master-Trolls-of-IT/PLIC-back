package repository

import (
	"gaia-api/infrastructure/model/requests/meal"
	"github.com/lib/pq"
)

type MealRepo struct {
	data *Database
}

func NewMealRepository(db *Database) *MealRepo {
	return &MealRepo{data: db}
}

func (mealRepo *MealRepo) SaveMeal(meal meal.Meal) error {
	var database = mealRepo.data.DB
	var barcodes string[]

	query := `WITH meal_insert AS ( 
				INSERT INTO meal (title, user_email) VALUES ($1, $2) RETURNING id)


			INSERT INTO meal_product (meal_id, product_id)
			SELECT meal_insert.id, product.id from meal_insert join product on product.barcode = ANY($3::text[])`

	_, err := database.Exec(query, meal.Title, meal.UserEmail, pq.Array(meal.Products))
	if err != nil {
		return err
	}
	return nil
}
