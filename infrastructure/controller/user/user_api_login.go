package user

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	user *User
}

func NewLoginController(user *User) *LoginController {
	loginController := &LoginController{user: user}
	loginController.Start()

	return loginController
}

func (loginController *LoginController) Start() {
	loginController.user.GinEngine.POST("/login", loginController.login)
}

func (loginController *LoginController) login(context *gin.Context) {
	var login = request.Login{}
	//binds Json Body to Entities.Login Class
	if err := context.ShouldBindJSON(&login); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	}
	var userRepo = *loginController.user.UserService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		returnAPI.Error(context, http.StatusUnauthorized)
	} else if loggedIn {
		var user shared.User
		if login.Email == "" {
			user, _ = userRepo.GetUserByUsername(login.Username)
		} else {
			user, _ = userRepo.GetUserByEmail(login.Email)
		}

		// A function that generates a token using JWT
		returnAPI.Success(context, http.StatusOK, user)
	} else {
		returnAPI.Error(context, http.StatusInternalServerError)
	}
}
