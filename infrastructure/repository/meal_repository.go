package repository

import (
	"fmt"
	"gaia-api/infrastructure/model/requests/meal"
	"github.com/lib/pq"
)

type MealRepo struct {
	data *Database
}

func NewMealRepository(db *Database) *MealRepo {
	return &MealRepo{data: db}
}

func (mealRepo *MealRepo) SaveMeal(myMeal meal.Meal) error {
	var database = mealRepo.data.DB
	var mealID string

	var barcodes []string
	var tag_labels []string

	for _, product := range myMeal.Products {
		barcodes = append(barcodes, product.Barcode)
	}

	for _, tag := range myMeal.Tags {
		tag_labels = append(tag_labels, tag.Label)
	}

	// Insert product quantity for each meal product
	productQuantityInsertQuery := `INSERT INTO product_quantity (quantity, product_id)
									SELECT $1::numeric, product.id
									FROM product
									WHERE product.barcode = $2`

	stmt, err := database.Prepare(productQuantityInsertQuery)
	if err != nil {
		return err
	}

	for _, product := range myMeal.Products {
		fmt.Print(product.Quantity, " ", product.Barcode)
		_, err := stmt.Exec(product.Quantity, product.Barcode)
		if err != nil {
			return err
		}
	}

	//Inserts the meal and retrieves the meal id
	mealInsertQuery := `INSERT INTO meal (title, user_email) VALUES ($1, $2) RETURNING id`

	//Associates the meal id with the corresponding product_id  / product_quantity_id
	mealProductQuantitiesInsert := `INSERT INTO meal_product_quantity (meal_id, product_quantity_id)
										SELECT $1, product_quantity.id  FROM product_quantity JOIN product ON product_quantity.product_id = product.id WHERE product.barcode = ANY($2::text[])`

	//Associates the  meal id with the corresponding tags
	mealTagsInsert := `
    INSERT INTO meal_tag (meal_id, tag_id)
    SELECT $1, tag.id
    FROM tag
	WHERE tag.label = ANY($2::TEXT[])
`

	err = database.QueryRow(mealInsertQuery, myMeal.Title, myMeal.UserEmail).Scan(&mealID)
	if err != nil {
		return err
	}

	_, err = database.Exec(mealProductQuantitiesInsert, mealID, pq.Array(barcodes))
	if err != nil {
		return err
	}

	_, err = database.Exec(mealTagsInsert, mealID, pq.Array(tag_labels))
	if err != nil {
		return err
	}

	return nil
}
