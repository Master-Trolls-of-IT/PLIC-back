package token

import (
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
		context.JSON(http.StatusBadRequest, accessTokenController.token.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, accessTokenController.token.ReturnAPIData.CheckToken(isTokenValid))
	}
}

func (accessTokenController *AccessTokenController) getAccessToken(context *gin.Context) {
	// Use GenerateAccessToken function to generate a new access token
	accessToken, err := helper.GenerateAccessToken(context.Param("password"), []byte(context.Param("refresh_token")))
	if err != nil {
		context.JSON(http.StatusBadRequest, accessTokenController.token.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, accessTokenController.token.ReturnAPIData.GetToken(accessToken))
}
