package port

import "gaia-api/domain/entity"

type ProductInterface interface {
	GetProductByBarCode(barcode string) (entity.Product, error)
	SaveProduct(product entity.Product) (bool, error)
}
