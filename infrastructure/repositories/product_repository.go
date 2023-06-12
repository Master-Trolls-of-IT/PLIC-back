package repositories

import "gaia-api/domain/entities"

type ProductRepo struct {
	data *Database
}

func NewProductRepository(db *Database) *ProductRepo {
	return &ProductRepo{data: db}
}

func (productRepo *ProductRepo) getProduct(query string, args ...interface{}) (entities.Nutrient, error) {
	stmt, err := productRepo.data.DB.Prepare(query)
	if err != nil {
		return entities.Nutrient{}, err
	}
	var nutrient entities.Nutrient
	err = stmt.QueryRow(args...).Scan(&nutrient.Id, &nutrient.Name)
	if err != nil {
		return entities.Nutrient{}, err
	}
	return nutrient, nil
}

func (productRepo *ProductRepo) GetProductByBarCode(barcode string) (entities.Nutrient, error) {
	return productRepo.getProduct("SELECT * FROM NUTRIENT WHERE BARCODE LIKE $1", barcode)
}
