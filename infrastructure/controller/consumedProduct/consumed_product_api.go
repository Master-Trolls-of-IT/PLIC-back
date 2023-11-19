package consumedProduct

import (
	interfaces "gaia-api/application/interface"
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type ConsumedProduct struct {
	GinEngine            *gin.Engine
	AuthService          *service.AuthService
	ReturnAPIData        *interfaces.ReturnAPIData
	OpenFoodFactsService *service.OpenFoodFactsService
}

func NewConsumedProductController(ginEngine *gin.Engine, authService *service.AuthService, returnAPIData *interfaces.ReturnAPIData, openFoodFactsService *service.OpenFoodFactsService) *ConsumedProduct {
	consumedProduct := &ConsumedProduct{GinEngine: ginEngine, AuthService: authService, ReturnAPIData: returnAPIData, OpenFoodFactsService: openFoodFactsService}
	consumedProduct.Start()
	return consumedProduct
}

func (consumedProduct *ConsumedProduct) Start() {
	NewGetController(consumedProduct)
	NewAddController(consumedProduct)
	NewUpdateController(consumedProduct)
	NewDeleteController(consumedProduct)
}
