package recipe

import (
	interfaces "gaia-api/application/interface"
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	GinEngine     *gin.Engine
	AuthService   *service.AuthService
	ReturnAPIData *interfaces.ReturnAPIData
}

func NewRecipeController(ginEngine *gin.Engine, authService *service.AuthService, returnAPIData *interfaces.ReturnAPIData) *Recipe {
	recipe := &Recipe{GinEngine: ginEngine, AuthService: authService, ReturnAPIData: returnAPIData}
	recipe.Start()
	return recipe
}

func (recipe *Recipe) Start() {
}
