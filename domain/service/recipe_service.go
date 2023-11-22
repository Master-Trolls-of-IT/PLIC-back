package service

import ports "gaia-api/domain/port"

type RecipeService struct {
	RecipeRepo *ports.RecipeInterface
}

func NewRecipeService(recipeInterface ports.RecipeInterface) *RecipeService {
	return &RecipeService{
		RecipeRepo: &recipeInterface,
	}
}
