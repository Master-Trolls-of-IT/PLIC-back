package response

type MealTag struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type Meal struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	UserEmail   string    `json:"email"`
	Products    []Product `json:"products"`
	Tags        []MealTag `json:"tags"`
	IsFavourite bool      `json:"isFavourite"`
	NbProducts  int       `json:"numberOfProducts"`
}
