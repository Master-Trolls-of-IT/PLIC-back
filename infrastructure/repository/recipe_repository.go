package repository

import (
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
)

type RecipeRepo struct {
	data *Database
}

func NewRecipeRepository(db *Database) *RecipeRepo {
	return &RecipeRepo{data: db}
}

func (recipeRepo *RecipeRepo) SaveRecipe(recipe request.Recipe) (*response.Recipe, error) {
	panic("implement me")
}

func (recipeRepo *RecipeRepo) GetAllRecipes(userEmail string) ([]response.Recipe, error) {
	panic("implement me")
}

func (recipeRepo *RecipeRepo) GetUserRecipes(userEmail string) ([]response.Recipe, error) {
	panic("implement me")
}

func (recipeRepo *RecipeRepo) DeleteRecipe(recipeID int) error {
	panic("implement me")
}
func (recipeRepo *RecipeRepo) UpdateRecipe(recipeID int, recipe request.Recipe) (*response.Recipe, error) {
	panic("implement me")
}
