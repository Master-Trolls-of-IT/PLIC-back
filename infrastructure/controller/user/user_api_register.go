package user

import (
	"gaia-api/domain/entity"
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
	var user = entity.User{}
	//binds Json Body to Entities.User Class
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, registerController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	var userRepo = *registerController.user.AuthService.UserRepo
	registered, err := userRepo.Register(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, registerController.user.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else if registered {
		context.JSON(http.StatusOK, registerController.user.ReturnAPIData.RegisterSuccess(user))
	} else {
		context.JSON(http.StatusConflict, registerController.user.ReturnAPIData.Error(http.StatusConflict, "Nom d'utilisateur ou email déjà utilisée"))
	}
}
