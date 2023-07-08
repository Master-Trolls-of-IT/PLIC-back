package repository

import (
	"database/sql"
	"gaia-api/domain/entity"
	"time"
)

type ProductRepo struct {
	data *Database
}

func NewProductRepository(db *Database) *ProductRepo {
	return &ProductRepo{data: db}
}

func (productRepo *ProductRepo) getProduct(query string, args ...interface{}) (entity.Product, error) {
	stmt, err := productRepo.data.DB.Prepare(query)
	if err != nil {
		return entity.Product{}, err
	}
	var product entity.Product
	err = stmt.QueryRow(args...).Scan(&product.ID, &product.Name, &product.Nutrients.EnergyKj, &product.Nutrients.EnergyKcal,
		&product.Nutrients.Fat, &product.Nutrients.SaturatedFat, &product.Nutrients.Sugar, &product.Nutrients.Fiber,
		&product.Nutrients.Proteins, &product.Nutrients.Salt, &product.ImageURL, &product.NutriScore.Score,
		&product.NutriScore.Grade)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) GetProductByBarCode(barcode string) (entity.Product, error) {
	return productRepo.getProduct("SELECT id, name, energy_kj, energy_kcal, fat, saturated_fat, sugar, fiber, "+
		"proteins, salt, image_url,  nutriscore_score, nutriscore_grade FROM product WHERE barcode = $1", barcode)
}

func (productRepo *ProductRepo) SaveProduct(product entity.Product, barcode string) (bool, error) {
	var database = productRepo.data.DB

	_, err := database.Exec("INSERT INTO product (name, energy_kj, energy_kcal, fat, saturated_fat, sugar,"+
		" fiber, proteins, salt, image_url,  nutriscore_score, nutriscore_grade, barcode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) ",
		product.Name, product.Nutrients.EnergyKj, product.Nutrients.EnergyKcal, product.Nutrients.Fat,
		product.Nutrients.SaturatedFat, product.Nutrients.Sugar, product.Nutrients.Fiber, product.Nutrients.Proteins,
		product.Nutrients.Salt, product.ImageURL, product.NutriScore.Score, product.NutriScore.Grade, barcode)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (productRepo *ProductRepo) SaveConsumedProduct(product entity.Product, userID int) (bool, error) {
	// Prepare the insert statement
	var database = productRepo.data.DB
	// Execute the insert statement
	_, err := database.Exec("INSERT INTO consumed_products (product_id, user_id, consumed_date)\n    VALUES ($1, $2, $3)", product.ID, userID, time.Now().UTC().Format("2006-01-02"))

	if err != nil {
		// Return an error if the insert operation fails
		return false, err
	}

	return true, nil
}

func (productRepo *ProductRepo) GetConsumedProductsByUserId(userID int) ([]entity.Product, error) {
	var database = productRepo.data.DB
	var products []entity.Product
	rows, err := database.Query("SELECT p.id, p.name, p.energy_kj, p.energy_kcal, p.fat, p.saturated_fat, p.sugar, p.fiber, p.proteins, p.salt, p.image_url, p.nutriscore_score, p.nutriscore_grade FROM consumed_products cp INNER JOIN product p ON cp.product_id = p.id WHERE cp.user_id = $1", userID)
	if err != nil {
		return products, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Nutrients.EnergyKj, &product.Nutrients.EnergyKcal,
			&product.Nutrients.Fat, &product.Nutrients.SaturatedFat, &product.Nutrients.Sugar, &product.Nutrients.Fiber,
			&product.Nutrients.Proteins, &product.Nutrients.Salt, &product.ImageURL, &product.NutriScore.Score,
			&product.NutriScore.Grade)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}
