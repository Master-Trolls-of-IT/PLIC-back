package mapping

type RecipeItem struct {
	ID              int    `db:"id"`
	Title           string `db:"title"`
	Author          string `db:"author"`
	Duration        int    `db:"duration"`
	Difficulty      string `db:"difficulty"`
	Rating          int    `db:"rating"`
	NumberOfRatings int    `db:"number_of_ratings"`
	Score           int    `db:"score"`
	Kcal            int    `db:"kcal"`
	Image           string `db:"icon"`
}
