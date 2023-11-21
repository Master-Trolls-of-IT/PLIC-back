package user

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateController struct {
	user *User
}

func NewUpdateController(user *User) *UpdateController {
	updateController := &UpdateController{user: user}
	updateController.Start()

	return updateController
}

func (updateController *UpdateController) Start() {
	updateController.user.GinEngine.PATCH("/users/:id", updateController.updateUser)
}

func (updateController *UpdateController) updateUser(context *gin.Context) {
	var user = shared.User{}
	var userId, err = strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&user); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	var userRepo = *updateController.user.UserService.UserRepo
	newUser, err := userRepo.UpdateUserById(userId, &user)

	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusOK, newUser)
	}
}
