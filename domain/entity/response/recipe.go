package response

type RecipeTag struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Recipe struct {
	RecipeItem RecipeItem `json:"recipeItem"`
}

type RecipeItem struct {
	ID              int         `json:"id"`
	Title           string      `json:"title"`
	Rating          int         `json:"rating"`
	NumberOfRatings int         `json:"numberOfRatings"`
	Duration        int         `json:"duration"`
	Difficulty      int         `json:"difficulty"`
	Score           int         `json:"score"`
	Ingredients     []string    `json:"ingredients"`
	Author          string      `json:"author"`
	Steps           []string    `json:"steps"`
	Tags            []RecipeTag `json:"tags"`
	IsFavorite      bool        `json:"isFavorite"`
	Kcal            int         `json:"kcal"`
	Image           string      `json:"image"`
}
