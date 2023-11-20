package repository

import (
	"database/sql"
	"gaia-api/domain/entity/mapping"
	"github.com/jackc/pgtype"
	"time"
)

type ProductRepo struct {
	data *Database
}

func NewProductRepository(db *Database) *ProductRepo {
	return &ProductRepo{data: db}
}

func (productRepo *ProductRepo) getProduct(query string, args ...interface{}) (mapping.Product, error) {
	stmt, err := productRepo.data.DB.Prepare(query)
	if err != nil {
		return mapping.Product{}, err
	}
	var product mapping.Product
	err = stmt.QueryRow(args...).Scan(&product.ID, &product.Brand, &product.Name, &product.Nutrients.EnergyKj, &product.Nutrients.EnergyKcal,
		&product.Nutrients.Fat, &product.Nutrients.SaturatedFat, &product.Nutrients.Sugar, &product.Nutrients.Fiber,
		&product.Nutrients.Proteins, &product.Nutrients.Salt, &product.ImageURL, &product.NutriScore.Score,
		&product.NutriScore.Grade, &product.EcoScore, &product.IsWater, &product.Quantity, &product.ServingQuantity, &product.ServingSize)
	if err != nil {
		return mapping.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) getConsumedProducts(query string, args ...interface{}) ([]mapping.ConsumedProduct, error) {
	stmt, err := productRepo.data.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var consumedProducts []mapping.ConsumedProduct
	for rows.Next() {
		var consumedProduct mapping.ConsumedProduct
		err := rows.Scan(&consumedProduct.Product.ID,
			&consumedProduct.Product.Name,
			&consumedProduct.Product.Brand,
			&consumedProduct.Product.Nutrients.EnergyKj,
			&consumedProduct.Product.Nutrients.EnergyKcal,
			&consumedProduct.Product.Nutrients.Fat,
			&consumedProduct.Product.Nutrients.SaturatedFat,
			&consumedProduct.Product.Nutrients.Sugar,
			&consumedProduct.Product.Nutrients.Fiber,
			&consumedProduct.Product.Nutrients.Proteins,
			&consumedProduct.Product.Nutrients.Salt,
			&consumedProduct.Product.ImageURL,
			&consumedProduct.Product.NutriScore.Score,
			&consumedProduct.Product.NutriScore.Grade,
			&consumedProduct.Product.EcoScore,
			&consumedProduct.Product.Barcode,
			&consumedProduct.Product.IsWater,
			&consumedProduct.Product.ServingQuantity,
			&consumedProduct.Quantity,
			&consumedProduct.ConsumedDate)
		if err != nil {
			return nil, err
		}
		consumedProducts = append(consumedProducts, consumedProduct)
	}

	return consumedProducts, nil
}

func (productRepo *ProductRepo) GetProductByBarCode(barcode string) (mapping.Product, error) {
	product, err := productRepo.getProduct("SELECT id, brand, name, energy_kj, energy_kcal, fat, saturated_fat, sugar, fiber, "+
		"proteins, salt, image_url, nutriscore_score, nutriscore_grade, ecoscore, isWater, quantity, serving_quantity, serving_size FROM product WHERE barcode = $1", barcode)

	if err != nil {
		return mapping.Product{}, err
	}

	nutrients100g, err := productRepo.getNutrientsDataByProductID(product.ID, "nutrients_100g")
	if err != nil {
		return mapping.Product{}, err
	}

	nutrientsValue, err := productRepo.getNutrientsDataByProductID(product.ID, "nutrients_value")
	if err != nil {
		return mapping.Product{}, err
	}

	nutrientsServing, err := productRepo.getNutrientsDataByProductID(product.ID, "nutrients_serving")
	if err != nil {
		return mapping.Product{}, err
	}

	product.Nutrients100g, product.NutrientsValue, product.NutrientsServing = mapping.Nutrients100g{nutrients100g}, mapping.NutrientsValue{nutrientsValue}, mapping.NutrientsServing{nutrientsServing}

	return product, nil
}

func (productRepo *ProductRepo) GetConsumedProductsByUserId(userID int) ([]mapping.ConsumedProduct, error) {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	query := "SELECT p.id, p.name ,p.brand, p.energy_kj, p.energy_kcal, p.fat, p.saturated_fat, p.sugar, p.fiber," +
		" p.proteins, p.salt, p.image_url, p.nutriscore_score, p.nutriscore_grade, p.ecoscore, p.barcode, p.iswater, p.serving_quantity, cp.quantity, cp.consumed_date" +
		" FROM consumed_products cp INNER JOIN product p ON cp.product_id = p.id WHERE cp.user_id = $1 AND cp.consumed_date = $2"
	return productRepo.getConsumedProducts(query, userID, today)
}

func (productRepo *ProductRepo) getNutrientsDataByProductID(productID int, table string) (mapping.Nutrients, error) {
	var db = productRepo.data.DB
	var nutrients mapping.Nutrients
	query := "SELECT  energy_kj, energy_kcal, fat, saturated_fat, sugar, fiber, proteins, salt from " + table + " WHERE product_id = $1"
	stmt, err := db.Prepare(query)
	if err != nil {
		return mapping.Nutrients{}, err
	}
	err = stmt.QueryRow(productID).Scan(&nutrients.EnergyKj, &nutrients.EnergyKcal, &nutrients.Fat, &nutrients.SaturatedFat, &nutrients.Sugar, &nutrients.Fiber, &nutrients.Proteins, &nutrients.Salt)
	if err != nil {
		return mapping.Nutrients{}, err
	}

	return nutrients, nil
}

func (productRepo *ProductRepo) SaveProduct(product mapping.Product, barcode string) (bool, error) {
	var database = productRepo.data.DB
	var productID int
	err := database.QueryRow("INSERT INTO product (brand, name, energy_kj, energy_kcal, fat, saturated_fat, sugar,"+
		" fiber, proteins, salt, image_url,  nutriscore_score, nutriscore_grade, ecoscore, barcode, isWater, quantity, serving_quantity, serving_size) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) RETURNING id",
		product.Brand, product.Name, product.Nutrients.EnergyKj, product.Nutrients.EnergyKcal, product.Nutrients.Fat,
		product.Nutrients.SaturatedFat, product.Nutrients.Sugar, product.Nutrients.Fiber, product.Nutrients.Proteins,
		product.Nutrients.Salt, product.ImageURL, product.NutriScore.Score, product.NutriScore.Grade, product.EcoScore,
		barcode, product.IsWater, product.Quantity, product.ServingQuantity, product.ServingSize).Scan(&productID)

	if err != nil {
		return false, err
	}
	err = productRepo.insertNutrientData(productID, "nutrients_100g", product.Nutrients100g.Nutrients)
	if err != nil {
		return false, err
	}

	err = productRepo.insertNutrientData(productID, "nutrients_Value", product.NutrientsValue.Nutrients)
	if err != nil {
		return false, err
	}

	err = productRepo.insertNutrientData(productID, "nutrients_Serving", product.NutrientsServing.Nutrients)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (productRepo *ProductRepo) SaveConsumedProduct(product mapping.Product, userID int, quantity int) (mapping.ConsumedProduct, error) {
	var database = productRepo.data.DB

	var date = time.Now().UTC()
	_, err := database.Exec("INSERT INTO consumed_products (product_id, user_id, quantity, consumed_date)\n    VALUES ($1, $2, $3, $4)", product.ID, userID, quantity, date.Format("2006-01-02"))

	if err != nil {
		return mapping.ConsumedProduct{}, err
	}
	consumedProduct := mapping.ConsumedProduct{Product: product, Quantity: quantity, ConsumedDate: pgtype.Date{Time: date, Status: pgtype.Present}}

	return consumedProduct, nil
}

func (productRepo *ProductRepo) insertNutrientData(productID int, table string, nutrients mapping.Nutrients) error {
	var db = productRepo.data.DB

	_, err := db.Exec(
		"INSERT INTO "+table+" (energy_kj, energy_kcal, fat, saturated_fat, sugar, fiber, proteins, salt, product_id) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		nutrients.EnergyKj, nutrients.EnergyKcal, nutrients.Fat,
		nutrients.SaturatedFat, nutrients.Sugar, nutrients.Fiber, nutrients.Proteins,
		nutrients.Salt, productID)

	if err != nil {
		return err
	}

	return nil
}

func (productRepo *ProductRepo) DeleteConsumedProduct(id int, userID int) (bool, error) {
	var database = productRepo.data.DB
	_, err := database.Exec("DELETE FROM consumed_products WHERE product_id = $1 AND user_id = $2", id, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (productRepo *ProductRepo) UpdateConsumedProductQuantity(quantity int, barcode string, userID int) error {
	query := "UPDATE consumed_products SET quantity=$1 WHERE product_id = (SELECT id FROM product WHERE barcode = $2) AND user_id = $3"

	var database = productRepo.data.DB
	_, err := database.Exec(query, quantity, barcode, userID)
	if err != nil {
		return err
	}
	return nil
}
