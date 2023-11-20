package token

import (
	"gaia-api/application/returnAPI"
	"gaia-api/infrastructure/controller/token/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenController struct {
	token *Token
}

func NewAccessTokenController(token *Token) *AccessTokenController {
	accessTokenController := &AccessTokenController{token: token}
	accessTokenController.Start()
	return accessTokenController
}

func (accessTokenController *AccessTokenController) Start() {
	accessTokenController.token.GinEngine.GET("/access_token/:refresh_token/:password", accessTokenController.getAccessToken)
	accessTokenController.token.GinEngine.GET("/access_token/check/:token", accessTokenController.checkAccessToken)
}

func (accessTokenController *AccessTokenController) checkAccessToken(context *gin.Context) {
	isTokenValid, err := helper.CheckAccessToken(context.Param("token"))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	} else if !isTokenValid {
		returnAPI.Error(context, http.StatusUnauthorized)
	} else {
		returnAPI.Success(context, http.StatusOK, isTokenValid)
	}
}

func (accessTokenController *AccessTokenController) getAccessToken(context *gin.Context) {
	accessToken, err := helper.GenerateAccessToken(context.Param("password"), []byte(context.Param("refresh_token")))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	} else {
		returnAPI.Success(context, http.StatusOK, accessToken)
	}
}
