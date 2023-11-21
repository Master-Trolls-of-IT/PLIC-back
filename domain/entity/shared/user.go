package shared

type User struct {
	Id              int     `db:"id" json:"id"`
	Rights          int     `db:"rights" json:"rights"`
	Email           string  `db:"email" json:"email"`
	Username        string  `db:"username" json:"username"`
	Birthdate       string  `db:"birthdate" json:"birthdate"`
	Weight          float32 `db:"weight" json:"weight"`
	Height          int     `db:"height" json:"height"`
	Gender          int     `db:"gender" json:"gender"`
	Sportiveness    int     `db:"sportiveness" json:"sportiveness"`
	BasalMetabolism int     `db:"basalmetabolism" json:"basalMetabolism"`
	Password        string  `db:"password" json:"password"`
	Pseudo          string  `db:"pseudo" json:"pseudo"`
	AvatarId        *int    `db:"avatar_id" json:"avatarId"`
}
