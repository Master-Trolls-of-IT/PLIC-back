package service

import (
	ports "gaia-api/domain/port"
)

type OpenFoodFactsService struct {
	ProductRepo *ports.ProductInterface
}

func NewOpenFoodFactsService(productRepo ports.ProductInterface) *OpenFoodFactsService {
	return &OpenFoodFactsService{ProductRepo: &productRepo}
}
