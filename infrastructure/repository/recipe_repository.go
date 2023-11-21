package repository

import (
	"fmt"
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
	"github.com/lib/pq"
	"strconv"
)

type RecipeRepo struct {
	data *Database
}

func NewRecipeRepository(db *Database) *RecipeRepo {
	return &RecipeRepo{data: db}
}

func (recipeRepo *RecipeRepo) SaveRecipe(recipe request.Recipe) (*response.Recipe, error) {
	database := recipeRepo.data.DB

	tagLabels := make([]string, len(recipe.Tags))
	for i, tag := range recipe.Tags {
		tagLabels[i] = tag.Label
	}
	ingredientLabels := make([]string, len(recipe.Ingredients))
	for i, ingredient := range recipe.Ingredients {
		ingredientLabels[i] = ingredient
	}
	stepLabels := make([]string, len(recipe.Steps))
	for i, step := range recipe.Steps {
		stepLabels[i] = step
	}

	// Insert the recipe and retrieve the recipe ID
	var recipeID string
	recipeInsertQuery := `INSERT INTO recipes (title, author, duration, difficulty) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := database.QueryRow(recipeInsertQuery, recipe.Title, recipe.UserEmail, recipe.Duration, recipe.Difficulty).Scan(&recipeID); err != nil {
		return nil, err
	}
	responseRecipeID, _ := strconv.Atoi(recipeID)
	// Insert recipe ingredients
	recipeIngredientInsert := `INSERT INTO recipe_ingredient (recipe_id, label) VALUES ($1, $2)`
	for _, ingredient := range recipe.Ingredients {
		if _, err := database.Exec(recipeIngredientInsert, responseRecipeID, ingredient); err != nil {
			fmt.Print("Error inserting recipe ingredients")
			fmt.Print(err)
			return nil, err
		}
	}
	// Associate recipe ID with tags
	recipeTagsInsert := `INSERT INTO recipe_tag (recipe_id, tag_id)
    						SELECT $1, tag.id FROM tag WHERE tag.label = ANY($2::TEXT[])`
	fmt.Print(tagLabels)
	if _, err := database.Exec(recipeTagsInsert, responseRecipeID, pq.Array(tagLabels)); err != nil {
		fmt.Print("Error inserting recipe tags")
		return nil, err
	}
	// Associate recipe ID with steps
	recipeStepsInsert := `INSERT INTO recipe_step (recipe_id, step, label) VALUES ($1, $2, $3)`
	for i, step := range recipe.Steps {
		if _, err := database.Exec(recipeStepsInsert, responseRecipeID, i+1, step); err != nil {
			fmt.Print("Error inserting recipe steps")
			fmt.Print(err)
			return nil, err
		}
	}
	// Retrieve the recipe
	responseRecipe := response.Recipe{RecipeItem: response.RecipeItem{ID: responseRecipeID, Title: recipe.Title, Rating: 0, NumberOfRatings: 0, Duration: recipe.Duration, Difficulty: recipe.Difficulty, Score: 0, Ingredients: recipe.Ingredients, Author: recipe.UserEmail, Steps: recipe.Steps, Tags: []response.RecipeTag{}, Kcal: 0, Image: ""}}
	responseRecipes := []response.Recipe{responseRecipe}

	err := recipeRepo.retrieveRecipeTags(responseRecipes)
	if err != nil {
		fmt.Print("Error retrieving recipe tags")
		return nil, err
	}
	responseRecipe = responseRecipes[0]
	return &responseRecipe, nil
}

func (recipeRepo *RecipeRepo) GetAllRecipes() ([]response.Recipe, error) {
	recipes, err := recipeRepo.retrieveRecipes()
	if err != nil {
		fmt.Print("Error retrieving recipes")
		fmt.Print(err)
		return nil, err
	}
	err = recipeRepo.retrieveRecipeIngredients(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe ingredients")
		fmt.Print(err)
		return nil, err
	}
	err = recipeRepo.retrieveRecipeSteps(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe steps")
		fmt.Print(err)
		return nil, err
	}
	err = recipeRepo.retrieveRecipeTags(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe tags")
		fmt.Print(err)
		return nil, err
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) GetUserRecipes(userEmail string) ([]response.Recipe, error) {
	recipes, err := recipeRepo.retrieveUserRecipes(userEmail)
	if err != nil {
		fmt.Print("Error retrieving recipes")
		return nil, err
	}
	err = recipeRepo.retrieveRecipeIngredients(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe ingredients")
		return nil, err
	}
	err = recipeRepo.retrieveRecipeSteps(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe steps")
		return nil, err
	}
	err = recipeRepo.retrieveRecipeTags(recipes)
	if err != nil {
		fmt.Print("Error retrieving recipe tags")
		return nil, err
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) DeleteRecipe(recipeID int) error {
	database := recipeRepo.data.DB

	deleteRecipeIngredientsQuery := `DELETE FROM recipe_ingredient WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeIngredientsQuery, recipeID); err != nil {
		return err
	}
	deleteRecipeTagsQuery := `DELETE FROM recipe_tag WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeTagsQuery, recipeID); err != nil {
		return err
	}
	deleteRecipeStepsQuery := `DELETE FROM recipe_step WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeStepsQuery, recipeID); err != nil {
		return err
	}
	deleteRecipeQuery := `DELETE FROM recipes WHERE id = $1`
	if _, err := database.Exec(deleteRecipeQuery, recipeID); err != nil {
		return err
	}
	return nil
}

func (recipeRepo *RecipeRepo) UpdateRecipe(recipeID int, recipe request.Recipe) (*response.Recipe, error) {
	database := recipeRepo.data.DB

	tagLabels := make([]string, len(recipe.Tags))
	for i, tag := range recipe.Tags {
		tagLabels[i] = tag.Label
	}
	ingredientLabels := make([]string, len(recipe.Ingredients))
	for i, ingredient := range recipe.Ingredients {
		ingredientLabels[i] = ingredient
	}
	stepLabels := make([]string, len(recipe.Steps))
	for i, step := range recipe.Steps {
		stepLabels[i] = step
	}

	// Update the recipe
	recipeUpdateQuery := `UPDATE recipes SET title = $1, author = $2, duration = $3, difficulty = $4 WHERE id = $5`
	if _, err := database.Exec(recipeUpdateQuery, recipe.Title, recipe.UserEmail, recipe.Duration, recipe.Difficulty, recipeID); err != nil {
		fmt.Print("Error updating recipe")
		fmt.Print(err)
		return nil, err
	}
	// Delete all recipe ingredients
	deleteRecipeIngredientsQuery := `DELETE FROM recipe_ingredient WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeIngredientsQuery, recipeID); err != nil {
		fmt.Print("Error deleting recipe ingredients")
		fmt.Print(err)
		return nil, err
	}
	// Insert recipe ingredients
	recipeIngredientInsert := `INSERT INTO recipe_ingredient (recipe_id, label) VALUES ($1, $2)`
	for _, ingredient := range recipe.Ingredients {
		if _, err := database.Exec(recipeIngredientInsert, recipeID, ingredient); err != nil {
			fmt.Print("Error inserting recipe ingredients")
			fmt.Print(err)
			return nil, err
		}
	}
	// Delete all recipe tags
	deleteRecipeTagsQuery := `DELETE FROM recipe_tag WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeTagsQuery, recipeID); err != nil {
		fmt.Print("Error deleting recipe tags")
		fmt.Print(err)
		return nil, err
	}
	// Insert recipe tags
	recipeTagsInsert := `INSERT INTO recipe_tag (recipe_id, tag_id)
    						SELECT $1, tag.id FROM tag WHERE tag.label = ANY($2::TEXT[])`
	if _, err := database.Exec(recipeTagsInsert, recipeID, pq.Array(tagLabels)); err != nil {
		fmt.Print("Error inserting recipe tags")
		fmt.Print(err)
		return nil, err
	}
	// Delete all recipe steps
	deleteRecipeStepsQuery := `DELETE FROM recipe_step WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeStepsQuery, recipeID); err != nil {
		fmt.Print("Error deleting recipe steps")
		fmt.Print(err)
		return nil, err
	}
	// Insert recipe steps
	recipeStepsInsert := `INSERT INTO recipe_step (recipe_id, step, label) VALUES ($1, $2, $3)`
	for i, step := range recipe.Steps {
		if _, err := database.Exec(recipeStepsInsert, recipeID, i+1, step); err != nil {
			fmt.Print("Error inserting recipe steps")
			fmt.Print(err)
			return nil, err
		}
	}
	// Retrieve the recipe
	responseRecipe := response.Recipe{RecipeItem: response.RecipeItem{ID: recipeID, Title: recipe.Title, Rating: 0, NumberOfRatings: 0, Duration: recipe.Duration, Difficulty: recipe.Difficulty, Score: 0, Ingredients: recipe.Ingredients, Author: recipe.UserEmail, Steps: recipe.Steps, Tags: []response.RecipeTag{}, Kcal: 0, Image: ""}}
	responseRecipes := []response.Recipe{responseRecipe}

	err := recipeRepo.retrieveRecipeTags(responseRecipes)
	if err != nil {
		return nil, err
	}
	responseRecipe = responseRecipes[0]
	return &responseRecipe, nil

}

func (recipeRepo *RecipeRepo) retrieveRecipes() ([]response.Recipe, error) {
	var recipes []response.Recipe
	database := recipeRepo.data.DB

	recipesQuery := `SELECT id, title, author, duration, difficulty, rating, number_of_ratings, score, kcal, icon FROM recipes`
	rows, err := database.Query(recipesQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var recipe response.Recipe
		recipe.RecipeItem.Tags = []response.RecipeTag{}
		recipe.RecipeItem.Steps = []string{}
		recipe.RecipeItem.Ingredients = []string{}
		if err := rows.Scan(&recipe.RecipeItem.ID, &recipe.RecipeItem.Title, &recipe.RecipeItem.Author, &recipe.RecipeItem.Duration, &recipe.RecipeItem.Difficulty, &recipe.RecipeItem.Rating, &recipe.RecipeItem.NumberOfRatings, &recipe.RecipeItem.Score, &recipe.RecipeItem.Kcal, &recipe.RecipeItem.Image); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) retrieveUserRecipes(userEmail string) ([]response.Recipe, error) {
	var recipes []response.Recipe
	database := recipeRepo.data.DB

	recipesQuery := `SELECT id, title, author, duration, difficulty, rating, number_of_ratings, score, kcal, icon FROM recipes WHERE author = $1`
	rows, err := database.Query(recipesQuery, userEmail)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var recipe response.Recipe
		recipe.RecipeItem.Tags = []response.RecipeTag{}
		recipe.RecipeItem.Steps = []string{}
		recipe.RecipeItem.Ingredients = []string{}
		if err := rows.Scan(&recipe.RecipeItem.ID, &recipe.RecipeItem.Title, &recipe.RecipeItem.Author, &recipe.RecipeItem.Duration, &recipe.RecipeItem.Difficulty, &recipe.RecipeItem.Rating, &recipe.RecipeItem.NumberOfRatings, &recipe.RecipeItem.Score, &recipe.RecipeItem.Kcal, &recipe.RecipeItem.Image); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeIngredients(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeIngredientsQuery := `SELECT label FROM recipe_ingredient WHERE recipe_id = $1`
		rows, err := database.Query(recipeIngredientsQuery, recipe.RecipeItem.ID)
		if err != nil {
			return err
		}
		var ingredients []string
		for rows.Next() {
			var ingredient string
			if err := rows.Scan(&ingredient); err != nil {
				return err
			}
			ingredients = append(ingredients, ingredient)
		}
		recipes[i].RecipeItem.Ingredients = ingredients
	}
	return nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeSteps(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeStepsQuery := `SELECT label FROM recipe_step WHERE recipe_id = $1`
		rows, err := database.Query(recipeStepsQuery, recipe.RecipeItem.ID)
		if err != nil {
			return err
		}
		var steps []string
		for rows.Next() {
			var step string
			if err := rows.Scan(&step); err != nil {
				return err
			}
			steps = append(steps, step)
		}
		recipes[i].RecipeItem.Steps = steps
	}
	return nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeTags(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeTagsQuery := `SELECT tag.label, tag.color FROM recipe_tag INNER JOIN tag ON recipe_tag.tag_id = tag.id WHERE recipe_tag.recipe_id = $1`
		rows, err := database.Query(recipeTagsQuery, recipe.RecipeItem.ID)
		if err != nil {
			return err
		}
		var tags []response.RecipeTag
		for rows.Next() {
			var tag response.RecipeTag
			if err := rows.Scan(&tag.Label, &tag.Color); err != nil {
				return err
			}
			tags = append(tags, tag)
		}
		recipes[i].RecipeItem.Tags = tags
	}
	return nil
}
