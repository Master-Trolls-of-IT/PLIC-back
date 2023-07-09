package repository

import (
	"gaia-api/domain/entity"
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
	err = stmt.QueryRow(args...).Scan(&product.Name, &product.Nutrients.EnergyKj, &product.Nutrients.EnergyKcal,
		&product.Nutrients.Fat, &product.Nutrients.SaturatedFat, &product.Nutrients.Sugar, &product.Nutrients.Fiber,
		&product.Nutrients.Proteins, &product.Nutrients.Salt, &product.ImageURL, &product.NutriScore.Score,
		&product.NutriScore.Grade)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) GetProductByBarCode(barcode string) (entity.Product, error) {
	return productRepo.getProduct("SELECT name, energy_kj, energy_kcal, fat, saturated_fat, sugar, fiber, "+
		"proteins, salt, image_url,  nutriscore_score, nutriscore_grade FROM product WHERE barcode = $1", barcode)
}

func (productRepo *ProductRepo) SaveProduct(product entity.Product, barcode string) (bool, error) {
	var database = productRepo.data.DB

	_, err := database.Exec("INSERT INTO product (name, energy_kj, energy_kcal, fat, saturated_fat, sugar,"+
		" fiber, proteins, salt, image_url,  nutriscore_score, nutriscore_grade, barcode, isWater) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) ",
		product.Name, product.Nutrients.EnergyKj, product.Nutrients.EnergyKcal, product.Nutrients.Fat,
		product.Nutrients.SaturatedFat, product.Nutrients.Sugar, product.Nutrients.Fiber, product.Nutrients.Proteins,
		product.Nutrients.Salt, product.ImageURL, product.NutriScore.Score, product.NutriScore.Grade, barcode, product.IsWater)
	if err != nil {
		return false, err
	}
	return true, nil
}
