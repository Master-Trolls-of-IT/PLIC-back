package product

import (
	"database/sql"
	"errors"
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/response"
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	ginEngine      *gin.Engine
	ProductService *service.ProductService
}

func NewProductController(ginEngine *gin.Engine, productService *service.ProductService) *Product {
	product := &Product{ginEngine: ginEngine, ProductService: productService}
	product.Start()
	return product
}

func (product *Product) Start() {
	product.ginEngine.GET("/product/:barcode", product.getProduct)
}

func (product *Product) getProduct(context *gin.Context) {
	var barcode = context.Param("barcode")
	var productRepo = *product.ProductService.ProductRepo
	productEntity, dbError := productRepo.GetProductByBarCode(barcode)

	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)

	} else if productEntity == (response.Product{}) {
		mappedProduct, err := product.RetrieveOpenFoodFactsProduct(barcode)
		if err != nil {
			returnAPI.Error(context, http.StatusNotFound)

		} else {
			productSaved, err := productRepo.SaveProduct(mappedProduct, barcode)

			if productSaved {
				returnAPI.Success(context, http.StatusOK, mappedProduct)

			} else if err != nil {
				returnAPI.Error(context, http.StatusInternalServerError)
			}
		}
	} else {
		returnAPI.Success(context, http.StatusOK, productEntity)
	}
}
