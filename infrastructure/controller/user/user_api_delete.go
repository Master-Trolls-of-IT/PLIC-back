package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteController struct {
	user *User
}

func NewDeleteController(user *User) *DeleteController {
	deleteController := &DeleteController{user: user}
	deleteController.Start()

	return deleteController
}

func (deleteController *DeleteController) Start() {
	deleteController.user.GinEngine.DELETE("/users/:id", deleteController.deleteUser)
}

func (deleteController *DeleteController) deleteUser(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))
	var userRepo = *deleteController.user.AuthService.UserRepo
	userDeleted, dbError := userRepo.DeleteUser(userId)
	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, deleteController.user.ReturnAPIData.Error(http.StatusInternalServerError, dbError.Error()))
	} else if userDeleted {
		context.JSON(http.StatusOK, deleteController.user.ReturnAPIData.DeletedUser(http.StatusOK, "Utilisateur Supprimé"))
	} else {
		context.JSON(http.StatusNotFound, deleteController.user.ReturnAPIData.Error(http.StatusNotFound, "Utilisateur non existant dans la base de données"))
	}
}
