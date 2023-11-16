package consumedProduct

import (
	"database/sql"
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

	var userRepo = *getController.consumedProduct.AuthService.UserRepo
	user, dbError := userRepo.GetUserByEmail(email)

	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, getController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
	} else {
		var userId = user.Id

		var productRepo = *getController.consumedProduct.OpenFoodFactsService.ProductRepo
		consumedProducts, dbError := productRepo.GetConsumedProductsByUserId(userId)
		if dbError != nil && dbError != sql.ErrNoRows {
			context.JSON(http.StatusInternalServerError, getController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
		} else {
			context.JSON(http.StatusOK, getController.consumedProduct.ReturnAPIData.GetConsumedProductsSuccess(consumedProducts))

		}
	}

}
