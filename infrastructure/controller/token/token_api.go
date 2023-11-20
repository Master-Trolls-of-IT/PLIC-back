package token

import (
	"github.com/gin-gonic/gin"
)

type Token struct {
	GinEngine *gin.Engine
}

func NewTokenController(ginEngine *gin.Engine) *Token {
	token := &Token{GinEngine: ginEngine}
	token.Start()
	return token
}

func (token *Token) Start() {
	NewRefreshTokenController(token)
	NewAccessTokenController(token)
}
