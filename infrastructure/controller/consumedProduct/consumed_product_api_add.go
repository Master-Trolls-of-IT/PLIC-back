package consumedProduct

import (
	"database/sql"
	"errors"
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
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
	var consumedProductAdd request.ConsumedProductAdd
	if err := context.ShouldBindJSON(&consumedProductAdd); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	email := consumedProductAdd.Email
	barcode := consumedProductAdd.Barcode
	quantity := consumedProductAdd.Quantity

	var productRepo = *addController.consumedProduct.OpenFoodFactsService.ProductRepo
	var userRepo = *addController.consumedProduct.AuthService.UserRepo

	product, dbError := productRepo.GetProductByBarCode(barcode)
	user, dbError := userRepo.GetUserByEmail(email)
	var userId = user.Id

	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)

	} else if product == (response.Product{}) {
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
			returnAPI.Success(context, http.StatusCreated, productSaved)
		}
	}
}
