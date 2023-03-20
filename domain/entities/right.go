package entities

type Right struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"access_level"`
}
