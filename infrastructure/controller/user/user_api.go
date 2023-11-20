package user

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	GinEngine   *gin.Engine
	AuthService *service.AuthService
}

func NewUserController(ginEngine *gin.Engine, authService *service.AuthService) *User {
	user := &User{GinEngine: ginEngine, AuthService: authService}
	user.Start()
	return user
}

func (user *User) Start() {
	NewLoginController(user)
	NewCheckUserController(user)
	NewRegisterController(user)
	NewUpdateController(user)
	NewDeleteController(user)
}
