package consumedProduct

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type ConsumedProduct struct {
	GinEngine      *gin.Engine
	UserService    *service.UserService
	ProductService *service.ProductService
}

func NewConsumedProductController(ginEngine *gin.Engine, userService *service.UserService, productService *service.ProductService) *ConsumedProduct {
	consumedProduct := &ConsumedProduct{GinEngine: ginEngine, UserService: userService, ProductService: productService}
	consumedProduct.Start()
	return consumedProduct
}

func (consumedProduct *ConsumedProduct) Start() {
	NewGetController(consumedProduct)
	NewAddController(consumedProduct)
	NewUpdateController(consumedProduct)
	NewDeleteController(consumedProduct)
}
