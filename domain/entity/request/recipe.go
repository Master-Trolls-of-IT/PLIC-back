package request

type RecipeTag struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Recipe struct {
	Title       string      `json:"title"`
	Duration    int         `json:"duration"`
	UserEmail   string      `json:"email"`
	Ingredients []string    `json:"ingredients"`
	Steps       []string    `json:"steps"`
	Tags        []RecipeTag `json:"tags"`
	Difficulty  string      `json:"difficulty"`
}
