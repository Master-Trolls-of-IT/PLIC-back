package request

type RecipeTags struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Recipe struct {
	Title       string       `json:"title"`
	Duration    string       `json:"duration"`
	UserEmail   string       `json:"email"`
	Ingredients []string     `json:"ingredients"`
	Steps       []string     `json:"steps"`
	Tags        []RecipeTags `json:"tags"`
	Difficulty  string       `json:"difficulty"`
}
