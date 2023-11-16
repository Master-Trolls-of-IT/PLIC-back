package consumedProduct

import (
	"database/sql"
	"gaia-api/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddController struct {
	consumedProduct *ConsumedProduct
}

func NewAddController(consumedProduct *ConsumedProduct) *AddController {
	addController := &AddController{consumedProduct: consumedProduct}
	addController.Start()
	return addController
}

func (addController *AddController) Start() {
	addController.consumedProduct.GinEngine.POST("/product/consumed", addController.addConsumedProduct)
}

func (addController *AddController) addConsumedProduct(context *gin.Context) {
	type MyRequestBody struct {
		Barcode  string `json:"barcode"`
		Email    string `json:"email"`
		Quantity string `json:"quantity"`
	}

	var requestBody MyRequestBody
	// Bind the request body
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, addController.consumedProduct.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	// Retrieve values from the request body
	email := requestBody.Email
	barcode := requestBody.Barcode
	quantity := requestBody.Quantity

	var productRepo = *addController.consumedProduct.OpenFoodFactsService.ProductRepo
	var userRepo = *addController.consumedProduct.AuthService.UserRepo

	product, dbError := productRepo.GetProductByBarCode(barcode)
	user, dbError := userRepo.GetUserByEmail(email)
	var userId = user.Id

	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, addController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))

	} else if product == (entity.Product{}) {
		context.JSON(http.StatusNotFound, addController.consumedProduct.ReturnAPIData.Error(http.StatusNotFound, "Produit non existant dans la base de données"))

	} else {
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			context.JSON(http.StatusInternalServerError, addController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, "Erreur de la conversion de la quantité  (atoi)"))
		}
		productSaved, err := productRepo.SaveConsumedProduct(product, userId, quantityInt)
		if err == nil {
			context.JSON(http.StatusOK, addController.consumedProduct.ReturnAPIData.ProductAddedToConsumed(productSaved))
		} else {
			context.JSON(http.StatusInternalServerError, addController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
		}
	}
}
