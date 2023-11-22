package recipe

import (
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteController struct {
	recipe *Recipe
}

func NewDeleteController(recipe *Recipe) *DeleteController {
	deleteController := &DeleteController{recipe: recipe}
	deleteController.Start()
	return deleteController
}

func (deleteController *DeleteController) Start() {
	deleteController.recipe.GinEngine.DELETE("/recipe/:id", deleteController.deleteRecipe)
}

func (deleteController *DeleteController) deleteRecipe(context *gin.Context) {
	recipeID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}
	var recipeRepo = *deleteController.recipe.RecipeService.RecipeRepo

	err = recipeRepo.DeleteRecipe(recipeID)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	}
	returnAPI.Success(context, http.StatusOK, nil)
}
