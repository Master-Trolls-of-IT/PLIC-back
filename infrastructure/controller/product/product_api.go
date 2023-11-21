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
	ginEngine            *gin.Engine
	openFoodFactsService *service.OpenFoodFactsService
	openFoodFactsAPI     *OpenFoodFactsAPI
}

func NewProductController(ginEngine *gin.Engine, openFoodFactsService *service.OpenFoodFactsService, openFoodFactsAPI *OpenFoodFactsAPI) *Product {
	product := &Product{ginEngine: ginEngine, openFoodFactsService: openFoodFactsService, openFoodFactsAPI: openFoodFactsAPI}
	product.Start()
	return product
}

func (product *Product) Start() {
	product.ginEngine.GET("/product/:barcode", product.getProduct)
}

func (product *Product) getProduct(context *gin.Context) {
	var barcode = context.Param("barcode")
	var productRepo = *product.openFoodFactsService.ProductRepo
	productEntity, dbError := productRepo.GetProductByBarCode(barcode)

	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)

	} else if productEntity == (response.Product{}) {
		openFoodFactAPI := product.openFoodFactsAPI
		mappedProduct, err := openFoodFactAPI.retrieveAndMapProduct(barcode)

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
