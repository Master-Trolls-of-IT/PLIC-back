package entity

type User struct {
	Id              int          `json:"id"`
	Rights          int          `json:"rights"`
	Email           string       `json:"email"`
	Username        string       `json:"username"`
	Password        string       `json:"password"`
	Pseudo          string       `json:"pseudo"`
	Birthdate       string       `json:"birthdate"`
	Weight          float32      `json:"weight"`
	Height          int          `json:"height"`
	Gender          Gender       `json:"gender"`
	Sportiveness    Sportiveness `json:"sportiveness"`
	BasalMetabolism int          `json:"basalmetabolism"`
}

type Gender int

const (
	male Gender = iota
	female
	other
)

type Sportiveness int

const (
	Sedentaire Sportiveness = iota
	ActifLeger
	ActifModere
	ActifIntense
	Athlete
)
