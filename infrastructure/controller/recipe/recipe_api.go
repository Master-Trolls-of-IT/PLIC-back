package recipe

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	GinEngine   *gin.Engine
	AuthService *service.AuthService
}

func NewRecipeController(ginEngine *gin.Engine, authService *service.AuthService) *Recipe {
	recipe := &Recipe{GinEngine: ginEngine, AuthService: authService}
	recipe.Start()
	return recipe
}

func (recipe *Recipe) Start() {
}
