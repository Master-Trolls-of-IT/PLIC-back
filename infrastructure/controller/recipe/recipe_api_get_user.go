package recipe

import (
	"fmt"
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserController struct {
	recipe *Recipe
}

func NewGetUserController(recipe *Recipe) *GetUserController {
	getUserController := &GetUserController{recipe: recipe}
	getUserController.Start()
	return getUserController
}

func (getUserController *GetUserController) Start() {
	getUserController.recipe.GinEngine.GET("/recipe/:email", getUserController.getUserRecipes)
}

func (getUserController *GetUserController) getUserRecipes(context *gin.Context) {
	var email = context.Param("email")
	var userRepo, recipeRepo = *getUserController.recipe.AuthService.UserRepo, *getUserController.recipe.OpenFoodFactsService.RecipeRepo

	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		fmt.Print(err)
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	}

	recipes, err := recipeRepo.GetUserRecipes(user.Email)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	} else if len(recipes) == 0 {
		returnAPI.Success(context, http.StatusOK, nil)
	}
	returnAPI.Success(context, http.StatusOK, recipes)
}
