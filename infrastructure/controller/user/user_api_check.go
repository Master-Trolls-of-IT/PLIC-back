package user

import (
	"gaia-api/domain/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckUserController struct {
	user *User
}

func NewCheckUserController(user *User) *CheckUserController {
	checkUserController := &CheckUserController{user: user}
	checkUserController.Start()

	return checkUserController
}

func (checkUserController *CheckUserController) Start() {
	checkUserController.user.GinEngine.POST("/check_user", checkUserController.checkUser)
}

func (checkUserController *CheckUserController) checkUser(context *gin.Context) {
	var login = entity.Login_info{}
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, checkUserController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	var userRepo = *checkUserController.user.AuthService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		context.JSON(http.StatusUnauthorized, checkUserController.user.ReturnAPIData.Error(http.StatusUnauthorized, "Informations de connexion non valides"))
	} else if loggedIn {

		//a function that generates a token using JWT
		context.JSON(http.StatusAccepted, checkUserController.user.ReturnAPIData.ValidPassword("Mot de passe Valide"))
	} else {
		context.JSON(http.StatusInternalServerError, checkUserController.user.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
}
