package service

import (
	ports "gaia-api/domain/port"
)

type ProductService struct {
	ProductRepo *ports.ProductInterface
}

func NewProductService(productInterface ports.ProductInterface) *ProductService {
	return &ProductService{
		ProductRepo: &productInterface,
	}
}
