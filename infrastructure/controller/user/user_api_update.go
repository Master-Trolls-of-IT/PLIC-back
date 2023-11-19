package user

import (
	"gaia-api/domain/entity"
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
	var user = entity.User{}
	var userId, err = strconv.Atoi(context.Param("id"))

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, updateController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	var userRepo = *updateController.user.AuthService.UserRepo
	newUser, err := userRepo.UpdateUserById(userId, &user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, updateController.user.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else {
		context.JSON(http.StatusOK, updateController.user.ReturnAPIData.UserUpdateSuccess(newUser))
	}
}
