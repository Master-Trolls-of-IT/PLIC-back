package ports

import "gaia-api/domain/entities"

type ProductInterface interface {
	GetProductByBarCode(barcode string) (entities.Nutrient, error)
}
