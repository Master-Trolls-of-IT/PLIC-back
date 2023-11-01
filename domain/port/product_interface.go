package port

import "gaia-api/domain/entity"

type ProductInterface interface {
	GetProductByBarCode(barcode string) (entity.Product, error)
	SaveProduct(product entity.Product, barcode string) (bool, error)
	SaveConsumedProduct(product entity.Product, id int, quantity int) (entity.ConsumedProduct, error)
	GetConsumedProductsByUserId(id int) ([]entity.ConsumedProduct, error)
	UpdateConsumedProductQuantity(quantity int, barcode string, userID int) error
	DeleteConsumedProduct(id int, userId int) (bool, error)
}
