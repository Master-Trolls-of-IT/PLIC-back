package consumedProduct

import (
	"database/sql"
	"errors"
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetController struct {
	consumedProduct *ConsumedProduct
}

func NewGetController(consumedProduct *ConsumedProduct) *GetController {
	getController := &GetController{consumedProduct: consumedProduct}
	getController.Start()
	return getController
}

func (getController *GetController) Start() {
	getController.consumedProduct.GinEngine.GET("/product/consumed/user/:email", getController.getConsumedProducts)
}

func (getController *GetController) getConsumedProducts(context *gin.Context) {
	email := context.Param("email")

	var userRepo = *getController.consumedProduct.UserService.UserRepo
	user, dbError := userRepo.GetUserByEmail(email)

	if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		var userId = user.Id

		var productRepo = *getController.consumedProduct.ProductService.ProductRepo
		consumedProducts, dbError := productRepo.GetConsumedProductsByUserId(userId)
		if dbError != nil && !errors.Is(sql.ErrNoRows, dbError) {
			returnAPI.Error(context, http.StatusInternalServerError)
		} else if len(consumedProducts) == 0 {
			returnAPI.Success(context, http.StatusOK, nil)
		} else {
			returnAPI.Success(context, http.StatusOK, consumedProducts)
		}
	}

}
