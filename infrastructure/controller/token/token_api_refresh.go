package token

import (
	"gaia-api/application/returnAPI"
	"gaia-api/infrastructure/controller/token/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenController struct {
	token *Token
}

func NewRefreshTokenController(token *Token) *RefreshTokenController {
	refreshTokenController := &RefreshTokenController{token: token}
	refreshTokenController.Start()

	return refreshTokenController
}

func (refreshTokenController *RefreshTokenController) Start() {
	refreshTokenController.token.GinEngine.GET("/refresh_token/:password", refreshTokenController.getRefreshToken)
	refreshTokenController.token.GinEngine.GET("/refresh_token/check/:token", refreshTokenController.checkRefreshToken)
}

func (refreshTokenController *RefreshTokenController) getRefreshToken(context *gin.Context) {
	refreshToken, err := helper.GenerateRefreshToken([]byte(context.Param("password")))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	} else {
		returnAPI.Success(context, http.StatusOK, refreshToken)
	}
}

func (refreshTokenController *RefreshTokenController) checkRefreshToken(context *gin.Context) {
	isTokenValid, err := helper.CheckAccessToken(context.Param("token"))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
	} else if !isTokenValid {
		returnAPI.Error(context, http.StatusUnauthorized)
	} else {
		returnAPI.Success(context, http.StatusOK, isTokenValid)
	}
}
