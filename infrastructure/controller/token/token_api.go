package token

import (
	interfaces "gaia-api/application/interface"
	"github.com/gin-gonic/gin"
)

type Token struct {
	GinEngine     *gin.Engine
	ReturnAPIData *interfaces.ReturnAPIData
}

func NewTokenController(ginEngine *gin.Engine, returnAPIData *interfaces.ReturnAPIData) *Token {
	token := &Token{GinEngine: ginEngine, ReturnAPIData: returnAPIData}
	token.Start()
	return token
}

func (token *Token) Start() {
	NewRefreshTokenController(token)
	NewAccessTokenController(token)
}
