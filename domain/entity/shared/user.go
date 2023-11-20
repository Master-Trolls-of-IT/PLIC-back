package shared

type User struct {
	Id              int     `json:"id"`
	Rights          int     `json:"rights"`
	Email           string  `json:"email"`
	Username        string  `json:"username"`
	Password        string  `json:"password"`
	Pseudo          string  `json:"pseudo"`
	Birthdate       string  `json:"birthdate"`
	Weight          float32 `json:"weight"`
	Height          int     `json:"height"`
	Gender          int     `json:"gender"`
	Sportiveness    int     `json:"sportiveness"`
	BasalMetabolism int     `json:"basalMetabolism"`
	AvatarId        int     `json:"avatarId"`
}
