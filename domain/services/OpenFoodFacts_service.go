package services

import (
	ports "gaia-api/domain/interfaces"
)

type OpenFoodFactsService struct {
	ProductRepo *ports.ProductInterface
}

func NewOpenFoodFactsService(productRepo ports.ProductInterface) *OpenFoodFactsService {
	return &OpenFoodFactsService{ProductRepo: &productRepo}
}
