package consumedProduct

import (
	"database/sql"
	"gaia-api/domain/entity/requests/consumedProductRequest"
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
	var consumedProductUpdateQuantity consumedProductRequest.ConsumedProductUpdateQuantity
	if err := context.ShouldBindJSON(&consumedProductUpdateQuantity); err != nil {
		context.JSON(http.StatusBadRequest, updateController.consumedProduct.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	var userRepo = *updateController.consumedProduct.AuthService.UserRepo
	user, dbError := userRepo.GetUserByEmail(consumedProductUpdateQuantity.UserEmail)
	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, updateController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
	}

	err := productRepo.UpdateConsumedProductQuantity(consumedProductUpdateQuantity.Quantity, consumedProductUpdateQuantity.Barcode, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, updateController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else {
		context.JSON(http.StatusOK, updateController.consumedProduct.ReturnAPIData.UpdateProduct(http.StatusOK, "La quantité consommée du produit a été mise à jour avec succès."))
	}

}
