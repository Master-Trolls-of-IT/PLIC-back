package repository

import (
	"gaia-api/infrastructure/model/requests/meal"
	response "gaia-api/infrastructure/model/responses/meal"
	"github.com/lib/pq"
	"strconv"
)

type MealRepo struct {
	data        *Database
	ProductRepo *ProductRepo
}

func NewMealRepository(db *Database, productRepo *ProductRepo) *MealRepo {
	return &MealRepo{data: db, ProductRepo: productRepo}
}

func (mealRepo *MealRepo) SaveMeal(myMeal request.Meal) (*response.Meal, error) {
	database := mealRepo.data.DB

	tagLabels := make([]string, len(myMeal.Tags))
	for i, tag := range myMeal.Tags {
		tagLabels[i] = tag.Label
	}

	// Insert the meal and retrieve the meal ID
	var mealID string
	mealInsertQuery := `INSERT INTO meal (title, user_email) VALUES ($1, $2) RETURNING id`
	if err := database.QueryRow(mealInsertQuery, myMeal.Title, myMeal.UserEmail).Scan(&mealID); err != nil {
		return nil, err
	}
	responseMealID, _ := strconv.Atoi(mealID)

	// Prepare the meal product insert statement
	mealProductInsertQuery := `INSERT INTO meal_product (meal_id, product_id, quantity)
							SELECT $1, product.id, $2 FROM product WHERE product.barcode = $3`
	mealProductInsertStmt, err := database.Prepare(mealProductInsertQuery)
	if err != nil {
		return nil, err
	}

	// Insert meal products
	for _, product := range myMeal.Products {
		if _, err := mealProductInsertStmt.Exec(mealID, product.Quantity, product.Barcode); err != nil {
			return nil, err
		}
	}

	// Associate meal ID with tags
	mealTagsInsert := `INSERT INTO meal_tag (meal_id, tag_id)
    						SELECT $1, tag.id FROM tag WHERE tag.label = ANY($2::TEXT[])`
	if _, err := database.Exec(mealTagsInsert, mealID, pq.Array(tagLabels)); err != nil {
		return nil, err
	}

	responseMeal := response.Meal{ID: responseMealID, Title: myMeal.Title, Tags: []response.MealTag{}, UserEmail: myMeal.UserEmail, Products: []response.MealProduct{}, IsFavourite: false}
	responseMeals := []response.Meal{responseMeal}
	err = mealRepo.retrieveMealProducts(responseMeals)
	if err != nil {
		return nil, err
	}
	err = mealRepo.retrieveMealTags(responseMeals)
	if err != nil {
		return nil, err
	}
	responseMeal = responseMeals[0]
	responseMeal.NbProducts = len(responseMeal.Products)
	return &responseMeal, nil
}

func (mealRepo *MealRepo) GetMeals(userEmail string) ([]response.Meal, error) {
	meals, err := mealRepo.retrieveMeals(userEmail)
	if err != nil {
		return []response.Meal{}, err
	}

	err = mealRepo.retrieveMealTags(meals)
	if err != nil {
		return []response.Meal{}, err
	}

	err = mealRepo.retrieveMealProducts(meals)
	if err != nil {
		return []response.Meal{}, err
	}

	return meals, nil
}

func (mealRepo *MealRepo) retrieveMeals(userEmail string) ([]response.Meal, error) {
	var meals []response.Meal
	database := mealRepo.data.DB

	getMealsQuery := "SELECT id, title, is_favourite FROM meal WHERE user_email = $1"
	rows, err := database.Query(getMealsQuery, userEmail)
	if err != nil {
		return []response.Meal{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var meal response.Meal
		meal.Tags = []response.MealTag{}
		meal.Products = []response.MealProduct{}
		err := rows.Scan(&meal.ID, &meal.Title, &meal.IsFavourite)
		if err != nil {
			return []response.Meal{}, err
		}
		meals = append(meals, meal)
	}

	err = rows.Err()
	if err != nil {
		return []response.Meal{}, err
	}

	return meals, nil
}

func (mealRepo *MealRepo) retrieveMealTags(meals []response.Meal) error {
	database := mealRepo.data.DB
	getTagsByMealIDQuery := "SELECT t.label, t.color FROM meal_tag mt JOIN tag t ON mt.tag_id = t.id WHERE mt.meal_id = $1"
	getTagsByMealIDStmt, err := database.Prepare(getTagsByMealIDQuery)
	if err != nil {
		return err
	}

	for i := range meals {
		meal := &meals[i]
		mealTagRows, err := getTagsByMealIDStmt.Query(meal.ID)
		if err != nil {
			return err
		}

		for mealTagRows.Next() {
			var tag response.MealTag
			err := mealTagRows.Scan(&tag.Label, &tag.Color)
			if err != nil {
				return err
			}
			meal.Tags = append(meal.Tags, tag)
		}
	}

	return nil
}

func (mealRepo *MealRepo) retrieveMealProducts(meals []response.Meal) error {
	database := mealRepo.data.DB
	getMealProductsQuery := "SELECT p.barcode, mp.quantity FROM product p JOIN meal_product mp ON p.id = mp.product_id WHERE mp.meal_id = $1"
	getMealProductsStmt, err := database.Prepare(getMealProductsQuery)
	if err != nil {
		return err
	}

	for i := range meals {
		meal := &meals[i]
		mealProductRows, err := getMealProductsStmt.Query(meal.ID)
		if err != nil {
			return err
		}

		for mealProductRows.Next() {
			var barcode string
			var quantity int
			err := mealProductRows.Scan(&barcode, &quantity)
			if err != nil {
				return err
			}
			productInfo, err := mealRepo.ProductRepo.GetProductByBarCode(barcode)
			if err != nil {
				return err
			}
			meal.Products = append(meal.Products, response.MealProduct{ProductInfo: productInfo, Quantity: quantity})
		}
		meal.NbProducts = len(meal.Products)
	}

	return nil
}
