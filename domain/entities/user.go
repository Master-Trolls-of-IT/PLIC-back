package entities

type User struct {
	Id               int          `json:"id"`
	Email            string       `json:"email"`
	Username         string       `json:"username"`
	Password         string       `json:"password"`
	Birthdate        string       `json:"birthdate"`
	Weight           float32      `json:"weight"`
	Height           int          `json:"height"`
	Gender           Gender       `json:"gender"`
	Sportiveness     Sportiveness `json:"sportiveness"`
	Basal_metabolism int          `json:"basal metabolism"`
}

type Gender string

const (
	male   Gender = "Homme"
	female Gender = "Femme"
	other  Gender = "Autre"
)

type Sportiveness int

const (
	sédentaire Sportiveness = iota
	actif_léger
	actif_modéré
	actif_intense
	athlete
)
