package token

import (
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
		context.JSON(http.StatusBadRequest, refreshTokenController.token.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, refreshTokenController.token.ReturnAPIData.GetToken(refreshToken))
}

func (refreshTokenController *RefreshTokenController) checkRefreshToken(context *gin.Context) {
	isTokenValid, err := helper.CheckAccessToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, refreshTokenController.token.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, refreshTokenController.token.ReturnAPIData.CheckToken(isTokenValid))
	}
}
