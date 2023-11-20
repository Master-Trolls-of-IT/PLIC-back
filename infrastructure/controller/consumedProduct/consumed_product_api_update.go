package consumedProduct

import (
	"database/sql"
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateController struct {
	consumedProduct *ConsumedProduct
}

func NewUpdateController(consumedProduct *ConsumedProduct) *UpdateController {
	updateController := &UpdateController{consumedProduct: consumedProduct}
	updateController.Start()
	return updateController
}

func (updateController *UpdateController) Start() {
	updateController.consumedProduct.GinEngine.PATCH("/product/consumed", updateController.updateConsumedProduct)
}

func (updateController *UpdateController) updateConsumedProduct(context *gin.Context) {
	var productRepo = *updateController.consumedProduct.OpenFoodFactsService.ProductRepo
	var consumedProductUpdateQuantity request.ConsumedProductUpdateQuantity
	if err := context.ShouldBindJSON(&consumedProductUpdateQuantity); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	var userRepo = *updateController.consumedProduct.AuthService.UserRepo
	user, dbError := userRepo.GetUserByEmail(consumedProductUpdateQuantity.UserEmail)
	if dbError != nil && dbError != sql.ErrNoRows {
		returnAPI.Error(context, http.StatusInternalServerError)
	}

	err := productRepo.UpdateConsumedProductQuantity(consumedProductUpdateQuantity.Quantity, consumedProductUpdateQuantity.Barcode, user.Id)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusOK, gin.H{})
	}

}
