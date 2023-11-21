package user

import (
	"database/sql"
	"gaia-api/application/returnAPI"
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
	var userRepo = *deleteController.user.UserService.UserRepo
	userDeleted, dbError := userRepo.DeleteUser(userId)
	if dbError != nil && dbError != sql.ErrNoRows {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else if userDeleted {
		returnAPI.Success(context, http.StatusOK, nil)
	} else {
		returnAPI.Error(context, http.StatusNotFound)
	}
}
