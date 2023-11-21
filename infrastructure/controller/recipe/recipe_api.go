package recipe

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	GinEngine            *gin.Engine
	AuthService          *service.AuthService
	OpenFoodFactsService *service.OpenFoodFactsService
}

func NewRecipeController(ginEngine *gin.Engine, authService *service.AuthService, openFoodFactsService *service.OpenFoodFactsService) *Recipe {
	recipe := &Recipe{GinEngine: ginEngine, AuthService: authService, OpenFoodFactsService: openFoodFactsService}
	recipe.Start()
	return recipe
}

func (recipe *Recipe) Start() {
	NewGetAllController(recipe)
	NewGetUserController(recipe)
	NewAddController(recipe)
	NewDeleteController(recipe)
	NewUpdateController(recipe)
}
