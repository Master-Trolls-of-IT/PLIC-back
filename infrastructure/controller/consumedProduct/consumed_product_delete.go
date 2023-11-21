package consumedProduct

import (
	"database/sql"
	"errors"
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteController struct {
	consumedProduct *ConsumedProduct
}

func NewDeleteController(consumedProduct *ConsumedProduct) *DeleteController {
	deleteController := &DeleteController{consumedProduct: consumedProduct}
	deleteController.Start()
	return deleteController
}

func (deleteController *DeleteController) Start() {
	deleteController.consumedProduct.GinEngine.DELETE("/product/consumed/:id/user/:email", deleteController.deleteConsumedProduct)
}

func (deleteController *DeleteController) deleteConsumedProduct(context *gin.Context) {
	email := context.Param("email")
	var userRepo = *deleteController.consumedProduct.AuthService.UserRepo
	user, dbError := userRepo.GetUserByEmail(email)
	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)
	}
	var userId = user.Id
	var id, _ = strconv.Atoi(context.Param("id"))
	var productRepo = *deleteController.consumedProduct.OpenFoodFactsService.ProductRepo

	productDeleted, dbError := productRepo.DeleteConsumedProduct(id, userId)
	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else if productDeleted {
		var productRepo = *deleteController.consumedProduct.OpenFoodFactsService.ProductRepo
		consumedProducts, dbError := productRepo.GetConsumedProductsByUserId(userId)
		if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
			returnAPI.Error(context, http.StatusInternalServerError)
		} else if len(consumedProducts) == 0 {
			returnAPI.Success(context, http.StatusOK, nil)
		} else {
			returnAPI.Success(context, http.StatusOK, consumedProducts)
		}
	} else {
		returnAPI.Error(context, http.StatusNotFound)
	}
}
