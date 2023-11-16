package user

import (
	interfaces "gaia-api/application/interface"
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	GinEngine     *gin.Engine
	AuthService   *service.AuthService
	ReturnAPIData *interfaces.ReturnAPIData
}

func NewUserController(ginEngine *gin.Engine, authService *service.AuthService, returnAPIData *interfaces.ReturnAPIData) *User {
	user := &User{GinEngine: ginEngine, AuthService: authService, ReturnAPIData: returnAPIData}
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
