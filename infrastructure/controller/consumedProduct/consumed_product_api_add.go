package consumedProduct

import (
	"database/sql"
	"gaia-api/application/returnAPI"
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
		returnAPI.Error(context, http.StatusBadRequest)
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
		returnAPI.Error(context, http.StatusInternalServerError)

	} else if product == (entity.Product{}) {
		returnAPI.Error(context, http.StatusNotFound)

	} else {
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			returnAPI.Error(context, http.StatusInternalServerError)
		}
		productSaved, err := productRepo.SaveConsumedProduct(product, userId, quantityInt)
		if err != nil {
			returnAPI.Error(context, http.StatusInternalServerError)
		} else {
			returnAPI.Success(context, http.StatusOK, productSaved)
		}
	}
}
