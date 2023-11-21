package user

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	GinEngine   *gin.Engine
	UserService *service.UserService
}

func NewUserController(ginEngine *gin.Engine, UserService *service.UserService) *User {
	user := &User{GinEngine: ginEngine, UserService: UserService}
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
