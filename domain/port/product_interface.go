package port

import (
	"gaia-api/domain/entity/mapping"
)

type ProductInterface interface {
	GetProductByBarCode(barcode string) (mapping.Product, error)
	SaveProduct(product mapping.Product, barcode string) (bool, error)
	SaveConsumedProduct(product mapping.Product, id int, quantity int) (mapping.ConsumedProduct, error)
	GetConsumedProductsByUserId(id int) ([]mapping.ConsumedProduct, error)
	UpdateConsumedProductQuantity(quantity int, barcode string, userID int) error
	DeleteConsumedProduct(id int, userId int) (bool, error)
}
