package repository

import (
	"gaia-api/infrastructure/model/requests/meal"
	response "gaia-api/infrastructure/model/responses/meal"
	"github.com/lib/pq"
)

type MealRepo struct {
	data *Database
}

func NewMealRepository(db *Database) *MealRepo {
	return &MealRepo{data: db}
}

func (mealRepo *MealRepo) SaveMeal(myMeal request.Meal) error {
	var database = mealRepo.data.DB
	var mealID string
	var tag_labels []string

	for _, tag := range myMeal.Tags {
		tag_labels = append(tag_labels, tag.Label)
	}

	//Inserts the meal and retrieves the meal id
	mealInsertQuery := `INSERT INTO meal (title, user_email) VALUES ($1, $2) RETURNING id`

	//Associates the meal id with the corresponding product_id  and quantity
	mealProductInsert := `INSERT INTO meal_product (meal_id, product_id, quantity)
							SELECT $1, product.id, $2 FROM  product ON WHERE product.barcode = $3`

	//Associates the  meal id with the corresponding tags
	mealTagsInsert := `INSERT INTO meal_tag (meal_id, tag_id)
    						SELECT $1, tag.id FROM tag WHERE tag.label = ANY($2::TEXT[])`

	err := database.QueryRow(mealInsertQuery, myMeal.Title, myMeal.UserEmail).Scan(&mealID)
	if err != nil {
		return err
	}

	for _, product := range myMeal.Products {
		_, err = database.Exec(mealProductInsert, mealID, product.Quantity, product.Barcode)
		if err != nil {
			return err
		}
	}

	_, err = database.Exec(mealTagsInsert, mealID, pq.Array(tag_labels))
	if err != nil {
		return err
	}

	return nil
}

func (mealRepo *MealRepo) GetMeal(userID int) (response.Meal, error) {

	//var Meals []response.Meal
	//
	//getMealsQuery := "SELECT id , title, is_favourite FROM meal WHERE user_email = $1"
	//
	//getTagIDSByIDQuery := "SELECT  label, color FROM  meal_tag mt join tag t on mt.tag_id = t.id where meal_id = $1"
	return response.Meal{}, nil
}
