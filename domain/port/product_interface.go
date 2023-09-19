package port

import "gaia-api/domain/entity"

type ProductInterface interface {
	GetProductByBarCode(barcode string) (entity.Product, error)
	SaveProduct(product entity.Product, barcode string) (bool, error)
	SaveConsumedProduct(product entity.Product, id int, quantity int) (bool, error)
	GetConsumedProductsByUserId(id int) ([]entity.ConsumedProduct, error)
	DeleteConsumedProduct(id int, userId int) (bool, error)
}
