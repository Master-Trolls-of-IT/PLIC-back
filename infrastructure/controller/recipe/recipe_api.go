package recipe

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	GinEngine   *gin.Engine
	UserService *service.UserService
}

func NewRecipeController(ginEngine *gin.Engine, UserService *service.UserService) *Recipe {
	recipe := &Recipe{GinEngine: ginEngine, UserService: UserService}
	recipe.Start()
	return recipe
}

func (recipe *Recipe) Start() {
}
