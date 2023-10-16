package repository

import "github.com/lib/pq"

type MealRepo struct {
	data *Database
}

func NewMealRepository(db *Database) *MealRepo {
	return &MealRepo{data: db}
}

func (mealRepo *MealRepo) SaveMeal(productsBarcodes []string, title string, email string, isFavourite bool) error {
	var database = mealRepo.data.DB

	query := `WITH meal_insert AS ( 
				INSERT INTO meal (title, user_email, is_favourite) VALUES ($1, $2, $3) RETURNING id)


			INSERT INTO meal_product (meal_id, product_id)
			SELECT meal_insert.id, product.id from meal_insert join product on product.barcode = ANY($4::text[])`

	_, err := database.Exec(query, title, email, isFavourite, pq.Array(productsBarcodes))
	if err != nil {
		return err
	}
	return nil
}
