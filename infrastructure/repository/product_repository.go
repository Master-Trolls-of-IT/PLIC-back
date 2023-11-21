package repository

import (
	"database/sql"
	"errors"
	"gaia-api/domain/entity/convert"
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/response"
	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
	"time"
)

type ProductRepo struct {
	data *Database
}

func NewProductRepository(db *Database) *ProductRepo {
	return &ProductRepo{data: db}
}

func (productRepo *ProductRepo) GetProductByBarCode(barcode string) (response.Product, error) {
	query := `SELECT * FROM product WHERE barcode = $1`
	productMapping, err := productRepo.getProduct(query, barcode)
	if err != nil {
		return response.Product{}, err
	}

	var product = convert.ProductMappingToProduct(productMapping)
	product.Nutrients, err = productRepo.getNutrientsDataByNutrientID(productMapping.NutrientsID)
	if err != nil {
		return response.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) GetProductByID(productID int) (response.Product, error) {
	query := `SELECT * FROM product WHERE id = $1`
	productMapping, err := productRepo.getProduct(query, productID)
	if err != nil {
		return response.Product{}, err
	}

	var product = convert.ProductMappingToProduct(productMapping)
	product.Nutrients, err = productRepo.getNutrientsDataByNutrientID(productMapping.NutrientsID)
	if err != nil {
		return response.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) getProduct(query string, args ...interface{}) (mapping.Product, error) {
	stmt, err := productRepo.data.DB.Preparex(query)
	if err != nil {
		return mapping.Product{}, err
	}

	var product mapping.Product
	err = stmt.Get(&product, args...)
	if errors.Is(sql.ErrNoRows, err) {
		return mapping.Product{}, nil
	}

	if err != nil {
		return mapping.Product{}, err
	}
	return product, nil
}

func (productRepo *ProductRepo) getNutrientsDataByNutrientID(nutrientID int) (response.Nutrients, error) {
	var db = productRepo.data.DB

	query := "SELECT  energy_kj, energy_kcal, fat, saturated_fat, carbohydrates, sugar, fiber, proteins, salt FROM nutrients WHERE id=$1"
	stmt, err := db.Preparex(query)
	if err != nil {
		return response.Nutrients{}, err
	}

	var nutrientsMapping mapping.Nutrients
	err = stmt.Get(&nutrientsMapping, nutrientID)
	if errors.Is(sql.ErrNoRows, err) {
		return response.Nutrients{}, nil
	}
	if err != nil {
		return response.Nutrients{}, err
	}

	return convert.NutrientsMappingToNutrients(nutrientsMapping), nil
}

func (productRepo *ProductRepo) SaveProduct(product response.Product, barcode string) (bool, error) {
	var database = productRepo.data.DB

	nutrientsID, err := productRepo.insertNutrientData(product.Nutrients)
	if err != nil {
		return false, err
	}

	var insertProductQuery = `
		INSERT INTO product (brand, name, image_url, nutriscore_score, nutriscore_grade, ecoscore, barcode, isWater,
		quantity, serving_quantity, serving_size, nutrients_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
		`
	var productID int
	err = database.QueryRow(insertProductQuery, product.Brand, product.Name, product.ImageURL, product.NutriScore.Score,
		product.NutriScore.Grade, product.EcoScore, barcode, product.IsWater, product.Quantity, product.ServingQuantity,
		product.ServingSize, nutrientsID).Scan(&productID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (productRepo *ProductRepo) getConsumedProducts(query string, args ...interface{}) ([]response.ConsumedProduct, error) {
	stmt, err := productRepo.data.DB.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Queryx(args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var consumedProducts []response.ConsumedProduct
	for rows.Next() {
		var consumedProductMapping mapping.ConsumedProduct
		err := rows.StructScan(&consumedProductMapping)
		if err != nil {
			return nil, err
		}

		var consumedProduct response.ConsumedProduct
		consumedProduct.Product, err = productRepo.GetProductByID(consumedProductMapping.ProductID)
		if err != nil {
			return nil, err
		}
		consumedProduct.Quantity = consumedProductMapping.Quantity
		consumedProduct.ConsumedDate = consumedProductMapping.ConsumedDate

		consumedProducts = append(consumedProducts, consumedProduct)
	}
	return consumedProducts, nil
}

func (productRepo *ProductRepo) GetConsumedProductsByUserId(userID int) ([]response.ConsumedProduct, error) {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")
	query := `SELECT * FROM consumed_products WHERE user_id=$1 AND consumed_date=$2`
	return productRepo.getConsumedProducts(query, userID, today)
}

func (productRepo *ProductRepo) SaveConsumedProduct(product response.Product, userID int, quantity int) (response.ConsumedProduct, error) {
	var database = productRepo.data.DB

	var date = time.Now().UTC()
	_, err := database.Exec("INSERT INTO consumed_products (product_id, user_id, quantity, consumed_date)\n    VALUES ($1, $2, $3, $4)", product.ID, userID, quantity, date.Format("2006-01-02"))

	if err != nil {
		return response.ConsumedProduct{}, err
	}
	consumedProduct := response.ConsumedProduct{Product: product, Quantity: quantity, ConsumedDate: pgtype.Date{Time: date, Status: pgtype.Present}}

	return consumedProduct, nil
}

func (productRepo *ProductRepo) insertNutrientData(nutrients response.Nutrients) (*int, error) {
	var db = productRepo.data.DB

	query := `
		INSERT INTO nutrients (energy_kj, energy_kcal, fat, saturated_fat, carbohydrates, sugar, fiber, proteins, salt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
		`
	var nutrientsID int
	err := db.QueryRow(query, nutrients.EnergyKj, nutrients.EnergyKcal, nutrients.Fat, nutrients.SaturatedFat,
		nutrients.Carbohydrates, nutrients.Sugar, nutrients.Fiber, nutrients.Proteins, nutrients.Salt).Scan(&nutrientsID)
	if err != nil {
		return nil, err
	}

	return &nutrientsID, nil
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
