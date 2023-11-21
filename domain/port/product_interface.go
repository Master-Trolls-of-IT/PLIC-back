package port

import (
	"gaia-api/domain/entity/response"
)

type ProductInterface interface {
	GetProductByBarCode(barcode string) (response.Product, error)
	SaveProduct(product *response.Product, barcode string) (bool, error)
	SaveConsumedProduct(product response.Product, id int, quantity int) (response.ConsumedProduct, error)
	GetConsumedProductsByUserId(id int) ([]response.ConsumedProduct, error)
	UpdateConsumedProductQuantity(quantity int, barcode string, userID int) error
	DeleteConsumedProduct(id int, userId int) (bool, error)
}
