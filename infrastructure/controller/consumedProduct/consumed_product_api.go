package consumedProduct

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type ConsumedProduct struct {
	GinEngine            *gin.Engine
	AuthService          *service.AuthService
	OpenFoodFactsService *service.OpenFoodFactsService
}

func NewConsumedProductController(ginEngine *gin.Engine, authService *service.AuthService, openFoodFactsService *service.OpenFoodFactsService) *ConsumedProduct {
	consumedProduct := &ConsumedProduct{GinEngine: ginEngine, AuthService: authService, OpenFoodFactsService: openFoodFactsService}
	consumedProduct.Start()
	return consumedProduct
}

func (consumedProduct *ConsumedProduct) Start() {
	NewGetController(consumedProduct)
	NewAddController(consumedProduct)
	NewUpdateController(consumedProduct)
	NewDeleteController(consumedProduct)
}
