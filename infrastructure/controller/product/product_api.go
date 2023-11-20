package product

import (
	"database/sql"
	interfaces "gaia-api/application/interface"
	"gaia-api/domain/entity"
	"gaia-api/domain/service"
	"gaia-api/infrastructure/error/openFoodFactsAPIError"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	ginEngine            *gin.Engine
	returnAPIData        *interfaces.ReturnAPIData
	openFoodFactsService *service.OpenFoodFactsService
	openFoodFactsAPI     *OpenFoodFactsAPI
}

func NewProductController(ginEngine *gin.Engine, returnAPIData *interfaces.ReturnAPIData, openFoodFactsService *service.OpenFoodFactsService, openFoodFactsAPI *OpenFoodFactsAPI) *Product {
	product := &Product{ginEngine: ginEngine, returnAPIData: returnAPIData, openFoodFactsService: openFoodFactsService, openFoodFactsAPI: openFoodFactsAPI}
	product.Start()
	return product
}

func (product *Product) Start() {
	product.ginEngine.GET("/product/:barcode", product.GetProduct)
}

func (product *Product) GetProduct(context *gin.Context) {
	var barcode = context.Param("barcode")
	var productRepo = *product.openFoodFactsService.ProductRepo
	productEntity, dbError := productRepo.GetProductByBarCode(barcode)

	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, product.returnAPIData.Error(http.StatusInternalServerError, dbError.Error()))

	} else if productEntity == (entity.Product{}) {
		openFoodFactAPI := product.openFoodFactsAPI
		mappedProduct, err := openFoodFactAPI.retrieveAndMapProduct(barcode)

		if _, productNotFound := err.(openFoodFactsAPIError.ProductNotFoundError); productNotFound {
			context.JSON(http.StatusInternalServerError, product.returnAPIData.ProductNotAvailable(barcode))

		} else {
			productSaved, err := productRepo.SaveProduct(mappedProduct, barcode)

			if productSaved {
				context.JSON(http.StatusOK, product.returnAPIData.ProductFound(mappedProduct))

			} else {
				context.JSON(http.StatusInternalServerError, product.returnAPIData.Error(http.StatusInternalServerError, err.Error()))
			}
		}
	} else {
		context.JSON(http.StatusOK, product.returnAPIData.ProductFound(productEntity))
	}
}
