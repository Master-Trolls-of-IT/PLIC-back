package recipe

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddController struct {
	recipe *Recipe
}

func NewAddController(recipe *Recipe) *AddController {
	addController := &AddController{recipe: recipe}
	addController.Start()
	return addController
}

func (addController *AddController) Start() {
	addController.recipe.GinEngine.POST("/recipe", addController.addRecipe)
}

func (addController *AddController) addRecipe(context *gin.Context) {
	var recipe request.Recipe
	if err := context.ShouldBindJSON(&recipe); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	var recipeRepo = *addController.recipe.RecipeService.RecipeRepo

	responseRecipe, err := recipeRepo.SaveRecipe(recipe)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusCreated, *responseRecipe)
	}
}
