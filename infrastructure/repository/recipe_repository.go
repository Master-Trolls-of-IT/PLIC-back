package repository

import (
	"gaia-api/domain/entity/convert"
	"gaia-api/domain/entity/mapping"
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/response"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
)

type RecipeRepo struct {
	data *Database
}

func NewRecipeRepository(db *Database) *RecipeRepo {
	return &RecipeRepo{data: db}
}

func (recipeRepo *RecipeRepo) AddRecipe(recipe request.Recipe) (*response.Recipe, error) {
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
	recipeInsertQuery := `INSERT INTO recipes (title, author, duration, difficulty, rating, score, kcal, icon, number_of_ratings) VALUES ($1, $2, $3, $4, 0, 0, 0, 'test', 0) RETURNING id`
	if err := database.QueryRow(recipeInsertQuery, recipe.Title, recipe.UserEmail, recipe.Duration, recipe.Difficulty).Scan(&recipeID); err != nil {

		return nil, err
	}
	responseRecipeID, _ := strconv.Atoi(recipeID)

	// Insert recipe ingredients / steps / tags
	err := recipeRepo.InsertRecipeIngredients(responseRecipeID, recipe.Ingredients)
	if err != nil {
		return nil, err
	}

	err = recipeRepo.InsertRecipeTags(responseRecipeID, tagLabels)
	if err != nil {
		return nil, err
	}

	err = recipeRepo.InsertRecipeSteps(responseRecipeID, recipe.Steps)
	if err != nil {
		return nil, err
	}

	// Retrieve the recipe with a conversion to response.Recipe
	responseRecipe := convert.RequestRecipeToResponseRecipe(recipe, responseRecipeID)
	responseRecipes := []response.Recipe{responseRecipe}

	err = recipeRepo.retrieveRecipeTags(responseRecipes)
	if err != nil {
		return nil, err
	}
	responseRecipe = responseRecipes[0]
	return &responseRecipe, nil
}

func (recipeRepo *RecipeRepo) GetAllRecipes() ([]response.Recipe, error) {
	recipes, err := recipeRepo.retrieveRecipes(`SELECT * FROM recipes`)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeIngredients(recipes)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeSteps(recipes)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeTags(recipes)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) GetUserRecipes(userEmail string) ([]response.Recipe, error) {
	recipes, err := recipeRepo.retrieveRecipes(`SELECT * FROM recipes WHERE author = $1`, userEmail)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeIngredients(recipes)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeSteps(recipes)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.retrieveRecipeTags(recipes)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) DeleteRecipe(recipeID int) error {
	database := recipeRepo.data.DB
	err := recipeRepo.DeleteRecipeIngredient(recipeID)
	if err != nil {
		return err
	}
	err = recipeRepo.DeleteRecipeStep(recipeID)
	if err != nil {
		return err
	}
	err = recipeRepo.DeleteRecipeTag(recipeID)
	if err != nil {
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

	// Get all the lists
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
		return nil, err
	}
	// Delete all recipe ingredients / steps / tags
	err := recipeRepo.DeleteRecipeIngredient(recipeID)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.DeleteRecipeStep(recipeID)
	if err != nil {
		return nil, err
	}
	err = recipeRepo.DeleteRecipeTag(recipeID)
	if err != nil {
		return nil, err
	}

	// Insert recipe ingredients
	err = recipeRepo.InsertRecipeIngredients(recipeID, recipe.Ingredients)
	if err != nil {
		return nil, err
	}
	// Insert recipe tags
	err = recipeRepo.InsertRecipeTags(recipeID, tagLabels)
	if err != nil {
		return nil, err
	}
	// Insert recipe steps
	err = recipeRepo.InsertRecipeSteps(recipeID, recipe.Steps)
	if err != nil {
		return nil, err
	}

	// Retrieve the recipe
	responseRecipe := response.Recipe{ID: recipeID, Title: recipe.Title, Rating: 0, NumberOfRatings: 0, Duration: recipe.Duration, Difficulty: recipe.Difficulty, Score: 0, Ingredients: recipe.Ingredients, Author: recipe.UserEmail, Steps: recipe.Steps, Tags: []response.RecipeTag{}, Kcal: 0, Image: ""}
	responseRecipes := []response.Recipe{responseRecipe}

	err = recipeRepo.retrieveRecipeTags(responseRecipes)
	if err != nil {
		return nil, err
	}
	responseRecipe = responseRecipes[0]
	return &responseRecipe, nil

}

func (recipeRepo *RecipeRepo) retrieveRecipes(query string, args ...interface{}) ([]response.Recipe, error) {
	stmt, err := recipeRepo.data.DB.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sqlx.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Queryx(args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var recipes []response.Recipe
	for rows.Next() {
		var recipeMapping mapping.Recipe
		err := rows.StructScan(&recipeMapping)
		if err != nil {

			return nil, err
		}
		var recipe response.Recipe
		recipe = convert.RecipeMappingToRecipe(recipeMapping)
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeIngredients(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeIngredientsQuery := `SELECT label FROM recipe_ingredient WHERE recipe_id = $1`
		rows, err := database.Query(recipeIngredientsQuery, recipe.ID)
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
		recipes[i].Ingredients = ingredients
	}
	return nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeSteps(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeStepsQuery := `SELECT label FROM recipe_step WHERE recipe_id = $1`
		rows, err := database.Query(recipeStepsQuery, recipe.ID)
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
		recipes[i].Steps = steps
	}
	return nil
}

func (recipeRepo *RecipeRepo) retrieveRecipeTags(recipes []response.Recipe) error {
	database := recipeRepo.data.DB
	for i, recipe := range recipes {
		recipeTagsQuery := `SELECT tag.label, tag.color FROM recipe_tag INNER JOIN tag ON recipe_tag.tag_id = tag.id WHERE recipe_tag.recipe_id = $1`
		rows, err := database.Query(recipeTagsQuery, recipe.ID)
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
		recipes[i].Tags = tags
	}
	return nil
}

func (recipeRepo *RecipeRepo) DeleteRecipeIngredient(recipeId int) error {
	database := recipeRepo.data.DB

	deleteRecipeIngredientsQuery := `DELETE FROM recipe_ingredient WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeIngredientsQuery, recipeId); err != nil {
		return err
	}
	return nil
}

func (recipeRepo *RecipeRepo) DeleteRecipeTag(recipeId int) error {
	database := recipeRepo.data.DB

	deleteRecipeTagsQuery := `DELETE FROM recipe_tag WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeTagsQuery, recipeId); err != nil {
		return err
	}
	return nil
}

func (recipeRepo *RecipeRepo) DeleteRecipeStep(recipeId int) error {
	database := recipeRepo.data.DB

	deleteRecipeStepsQuery := `DELETE FROM recipe_step WHERE recipe_id = $1`
	if _, err := database.Exec(deleteRecipeStepsQuery, recipeId); err != nil {
		return err
	}
	return nil
}

func (recipeRepo *RecipeRepo) InsertRecipeIngredients(recipeId int, ingredients []string) error {
	database := recipeRepo.data.DB

	recipeIngredientInsert := `INSERT INTO recipe_ingredient (recipe_id, label) VALUES ($1, $2)`
	for _, ingredient := range ingredients {
		if _, err := database.Exec(recipeIngredientInsert, recipeId, ingredient); err != nil {
			return err
		}
	}
	return nil
}

func (recipeRepo *RecipeRepo) InsertRecipeTags(recipeId int, tags []string) error {
	database := recipeRepo.data.DB

	recipeTagsInsert := `INSERT INTO recipe_tag (recipe_id, tag_id)
							SELECT $1, tag.id FROM tag WHERE tag.label = ANY($2::TEXT[])`
	if _, err := database.Exec(recipeTagsInsert, recipeId, pq.Array(tags)); err != nil {
		return err
	}
	return nil
}

func (recipeRepo *RecipeRepo) InsertRecipeSteps(recipeId int, steps []string) error {
	database := recipeRepo.data.DB

	recipeStepsInsert := `INSERT INTO recipe_step (recipe_id, step, label) VALUES ($1, $2, $3)`
	for i, step := range steps {
		if _, err := database.Exec(recipeStepsInsert, recipeId, i+1, step); err != nil {
			return err
		}
	}
	return nil
}
