package recipe

import (
	"fmt"
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllController struct {
	recipe *Recipe
}

func NewGetAllController(recipe *Recipe) *GetAllController {
	getAllController := &GetAllController{recipe: recipe}
	getAllController.Start()
	return getAllController
}

func (getAllController *GetAllController) Start() {
	getAllController.recipe.GinEngine.GET("/recipe", getAllController.getAllRecipe)
}

func (getAllController *GetAllController) getAllRecipe(context *gin.Context) {
	var recipeRepo = *getAllController.recipe.RecipeService.RecipeRepo

	recipes, err := recipeRepo.GetAllRecipes()
	if err != nil {
		fmt.Print(err)
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	}
	returnAPI.Success(context, http.StatusOK, recipes)
}
