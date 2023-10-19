package entity

type Meal struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserEmail   string `json:"user_email"`
	IsFavourite bool   `json:"is_favourite"`
}
