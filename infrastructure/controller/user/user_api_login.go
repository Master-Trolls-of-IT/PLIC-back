package user

import (
	"gaia-api/domain/entity"
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
	var login = entity.Login_info{}
	//binds Json Body to Entities.Login_info Class
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, loginController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	var userRepo = *loginController.user.AuthService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		context.JSON(http.StatusUnauthorized, loginController.user.ReturnAPIData.Error(http.StatusUnauthorized, "Informations de connexion non valides"))
	} else if loggedIn {
		var user entity.User
		if login.Email == "" {
			user, _ = userRepo.GetUserByUsername(login.Username)
		} else {
			user, _ = userRepo.GetUserByEmail(login.Email)
		}

		//a function that generates a token using JWT
		context.JSON(http.StatusAccepted, loginController.user.ReturnAPIData.LoginSuccess(user))
	} else {
		context.JSON(http.StatusInternalServerError, loginController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
}
