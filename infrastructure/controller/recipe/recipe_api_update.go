package recipe

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateController struct {
	recipe *Recipe
}

func NewUpdateController(recipe *Recipe) *UpdateController {
	updateController := &UpdateController{recipe: recipe}
	updateController.Start()
	return updateController
}

func (updateController *UpdateController) Start() {
	updateController.recipe.GinEngine.PUT("/recipe/:id", updateController.updateRecipe)
}

func (updateController *UpdateController) updateRecipe(context *gin.Context) {
	recipeId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}
	var recipe request.Recipe
	if err := context.ShouldBindJSON(&recipe); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}
	var recipeRepo = *updateController.recipe.OpenFoodFactsService.RecipeRepo

	responseRecipe, err := recipeRepo.UpdateRecipe(recipeId, recipe)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusCreated, *responseRecipe)
	}

}
