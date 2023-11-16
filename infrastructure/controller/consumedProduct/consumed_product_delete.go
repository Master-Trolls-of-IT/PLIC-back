package consumedProduct

import (
	"database/sql"
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
	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, deleteController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
	}
	var userId = user.Id
	var id, _ = strconv.Atoi(context.Param("id"))
	var productRepo = *deleteController.consumedProduct.OpenFoodFactsService.ProductRepo

	productDeleted, dbError := productRepo.DeleteConsumedProduct(id, userId)
	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, deleteController.consumedProduct.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
	} else if productDeleted {
		var productRepo = *deleteController.consumedProduct.OpenFoodFactsService.ProductRepo
		consumedProducts, dbError := productRepo.GetConsumedProductsByUserId(userId)
		if dbError != nil && dbError != sql.ErrNoRows {
			context.JSON(http.StatusInternalServerError, deleteController.consumedProduct.ReturnAPIData.DeletedProduct(http.StatusInternalServerError, "Could not retrieve the list of consumed products after deleting product"))
		}
		context.JSON(http.StatusOK, deleteController.consumedProduct.ReturnAPIData.ProductDeletedFromConsumed(consumedProducts))
	} else {
		context.JSON(http.StatusNotFound, deleteController.consumedProduct.ReturnAPIData.Error(http.StatusNotFound, "Produit non existant dans la base de données"))
	}
}
