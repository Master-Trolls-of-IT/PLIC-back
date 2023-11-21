package user

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
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
	var login = request.Login{}
	if err := context.ShouldBindJSON(&login); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	}
	var userRepo = *checkUserController.user.UserService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		returnAPI.Error(context, http.StatusUnauthorized)
	} else if loggedIn {
		// A function that generates a token using JWT
		returnAPI.Success(context, http.StatusOK, nil)
	} else {
		returnAPI.Error(context, http.StatusInternalServerError)
	}
}
