package port

import (
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
)

type RecipeInterface interface {
	AddRecipe(recipe request.Recipe) (*response.Recipe, error)
	GetAllRecipes() ([]response.Recipe, error)
	GetUserRecipes(userEmail string) ([]response.Recipe, error)
	DeleteRecipe(recipeID int) error
	UpdateRecipe(recipeID int, recipe request.Recipe) (*response.Recipe, error)
}
