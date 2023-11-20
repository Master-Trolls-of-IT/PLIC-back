package user

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct {
	user *User
}

func NewRegisterController(user *User) *RegisterController {
	registerController := &RegisterController{user: user}
	registerController.Start()

	return registerController
}

func (registerController *RegisterController) Start() {
	registerController.user.GinEngine.POST("/register", registerController.register)
}

func (registerController *RegisterController) register(context *gin.Context) {
	var user = shared.User{}
	//binds Json Body to Entities.User Class
	if err := context.ShouldBindJSON(&user); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	var userRepo = *registerController.user.AuthService.UserRepo
	registered, _ := userRepo.Register(&user)

	if registered {
		returnAPI.Success(context, http.StatusCreated, user)
	} else {
		returnAPI.Error(context, http.StatusInternalServerError)
	}
}
