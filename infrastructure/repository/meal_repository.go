package repository

type MealRepo struct {
	data *Database
}

func NewMealRepository(db *Database) *MealRepo {
	return &MealRepo{data: db}
}

func (mealRepo *MealRepo) SaveMeal(productsBarcodes []string, title string, email string, isFavourite bool) error {
	var database = mealRepo.data.DB
	var mealID int
	mealInsertQuery := "INSERT INTO meal (title, user_email, is_favourite) VALUES ($1, $2, $3) RETURNING id"

	if err := database.QueryRow(mealInsertQuery, title, email, isFavourite).Scan(&mealID); err != nil {
		return err
	}

	mealProductInsertQuery := "INSERT INTO meal_product (meal_id, product_id) VALUES ($1, $2)"
	mealProductInsertStatement, err := database.Prepare(mealProductInsertQuery)

	productSelectQuery := "SELECT id FROM product WHERE barcode = $1"
	productSelectStatement, err := database.Prepare(productSelectQuery)
	if err != nil {
		return err
	}

	for _, barcode := range productsBarcodes {
		var productID int

		err = productSelectStatement.QueryRow(barcode).Scan(&productID)
		if err != nil {
			return err
		}

		_, err = mealProductInsertStatement.Exec(mealID, productID)
		if err != nil {
			return err
		}

	}

	return nil
}
