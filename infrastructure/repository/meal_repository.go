package repository

import (
	"fmt"
	"gaia-api/infrastructure/model/requests/meal"
	response "gaia-api/infrastructure/model/responses/meal"
	"github.com/lib/pq"
)

type MealRepo struct {
	data        *Database
	ProductRepo *ProductRepo
}

func NewMealRepository(db *Database, productRepo *ProductRepo) *MealRepo {
	return &MealRepo{data: db, ProductRepo: productRepo}
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
							SELECT $1, product.id, $2 FROM  product WHERE product.barcode = $3`

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

func (mealRepo *MealRepo) GetMeals(userEmail string) ([]response.Meal, error) {
	var database = mealRepo.data.DB
	var productRepo = mealRepo.ProductRepo
	var meals []response.Meal

	//retrieve meals
	getMealsQuery := "SELECT id , title, is_favourite FROM meal WHERE user_email = $1"
	rows, err := database.Query(getMealsQuery, userEmail)
	if err != nil {
		return []response.Meal{}, err
	}

	//retrieve meal tags
	getTagsByMealIDQuery := "SELECT  t.label, t.color FROM  meal_tag mt join tag t on mt.tag_id = t.id where mt.meal_id = $1"
	getTagsByMealIDStmt, err := database.Prepare(getTagsByMealIDQuery)
	if err != nil {
		return []response.Meal{}, err
	}

	//retrieve meal products
	getMealProductsQuery := "SELECT p.barcode, mp.quantity from product p join meal_product mp on p.id = mp.product_id where mp.meal_id = $1"
	getMealProductsStmt, err := database.Prepare(getMealProductsQuery)
	if err != nil {
		return []response.Meal{}, err
	}

	for rows.Next() {
		var meal response.Meal
		meal.Tags = []response.MealTag{}
		meal.Products = []response.MealProduct{}
		err := rows.Scan(&meal.ID, &meal.Title, &meal.IsFavourite)
		if err != nil {
			fmt.Println("HERE Z")
			return []response.Meal{}, err
		}

		//retrieve meal tags
		mealTagRows, err := getTagsByMealIDStmt.Query(meal.ID)
		if err != nil {
			return []response.Meal{}, err
		}

		for mealTagRows.Next() {
			var tag response.MealTag
			err := mealTagRows.Scan(&tag.Label, &tag.Color)
			if err != nil {
				return []response.Meal{}, err
			}

			meal.Tags = append(meal.Tags, tag)
		}

		//retrieve meal products
		mealProductRows, err := getMealProductsStmt.Query(meal.ID)
		for mealProductRows.Next() {
			var barcode string
			var quantity int
			err := mealProductRows.Scan(&barcode, &quantity)
			if err != nil {
				return []response.Meal{}, err
			}
			productInfo, err := productRepo.GetProductByBarCode(barcode)
			if err != nil {
				return []response.Meal{}, err
			}
			meal.Products = append(meal.Products, response.MealProduct{ProductInfo: productInfo, Quantity: quantity})
		}

		meals = append(meals, meal)
	}
	return meals, nil
}
