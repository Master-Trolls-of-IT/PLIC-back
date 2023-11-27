package recipe

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	GinEngine     *gin.Engine
	UserService   *service.UserService
	RecipeService *service.RecipeService
}

func NewRecipeController(ginEngine *gin.Engine, UserService *service.UserService, recipeService *service.RecipeService) *Recipe {
	recipe := &Recipe{GinEngine: ginEngine, UserService: UserService, RecipeService: recipeService}
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
