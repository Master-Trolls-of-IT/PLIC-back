package repository

import (
	"database/sql"
	"gaia-api/domain/entity"
	"gaia-api/domain/entity/requests/mealRequest"
	"gaia-api/domain/entity/responses/meal"
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
	"strconv"
	"time"
)

type MealRepo struct {
	data        *Database
	ProductRepo *ProductRepo
}

func NewMealRepository(db *Database, productRepo *ProductRepo) *MealRepo {
	return &MealRepo{data: db, ProductRepo: productRepo}
}

func (mealRepo *MealRepo) SaveMeal(myMeal mealRequest.Meal) (*response.Meal, error) {
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

	responseMeal := response.Meal{ID: responseMealID, Title: myMeal.Title, Tags: []response.MealTag{}, UserEmail: myMeal.UserEmail, Products: []entity.Product{}, IsFavourite: false}
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
	responseMeal.UserEmail = myMeal.UserEmail
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
		meal.Products = []entity.Product{}
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
			productInfo.Quantity = strconv.Itoa(quantity)
			meal.Products = append(meal.Products, productInfo)
		}
		meal.NbProducts = len(meal.Products)
	}

	return nil
}

func (mealRepo *MealRepo) DeleteMeal(mealID int) error {
	database := mealRepo.data.DB

	deleteMealProductQuery := "DELETE FROM meal_product WHERE meal_id = $1"
	_, err := database.Exec(deleteMealProductQuery, mealID)
	if err != nil {
		return err
	}

	deleteMealTagQuery := "DELETE FROM meal_tag where meal_id = $1"
	_, err = database.Exec(deleteMealTagQuery, mealID)
	if err != nil {
		return err
	}

	deleteMealQuery := "DELETE FROM meal where id = $1"
	_, err = database.Exec(deleteMealQuery, mealID)
	if err != nil {
		return err
	}

	return nil
}

func (mealRepo *MealRepo) ConsumeMeal(meal response.Meal) ([]entity.ConsumedProduct, error) {
	database := mealRepo.data.DB
	var consumedProducts []entity.ConsumedProduct
	var userID int
	err := database.QueryRow("SELECT id FROM users where email= $1", meal.UserEmail).Scan(&userID)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO consumed_products (product_id, user_id, quantity, consumed_date) VALUES ($1, $2, $3, $4)"
	statement, err := database.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	transaction, err := database.Begin()
	if err != nil {
		return nil, err
	}
	defer func(transaction *sql.Tx) {
		txError := transaction.Rollback()
		if txError != nil {
			err = txError
		}
	}(transaction)

	for _, product := range meal.Products {
		quantity, err := strconv.Atoi(product.Quantity)
		date := time.Now().UTC()
		if err != nil {
			return nil, err
		}
		consumedProducts = append(consumedProducts, entity.ConsumedProduct{Product: product, Quantity: quantity, Consumed_Date: pgtype.Date{Time: date, Status: pgtype.Present}})
		_, err = statement.Exec(product.ID, userID, product.Quantity, date.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
	}

	if err = transaction.Commit(); err != nil {
		return nil, err
	}

	return consumedProducts, nil
}
