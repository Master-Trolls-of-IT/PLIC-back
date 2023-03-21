package entities

type User struct {
	Id              int          `json:"id"`
	Email           string       `json:"email"`
	Username        string       `json:"username"`
	Password        string       `json:"password"`
	Birthdate       string       `json:"birthdate"`
	Weight          float32      `json:"weight"`
	Height          int          `json:"height"`
	Gender          Gender       `json:"gender"`
	Sportiveness    Sportiveness `json:"sportiveness"`
	BasalMetabolism int          `json:"basal_metabolism"`
}

type Gender string

const (
	male   Gender = "Homme"
	female Gender = "Femme"
	other  Gender = "Autre"
)

type Sportiveness int

const (
	Sedentaire Sportiveness = iota
	ActifLeger
	ActifModere
	ActifIntense
	Athlete
)
