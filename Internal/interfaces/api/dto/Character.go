package dto

type CharacterDTO struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Race            string `json:"race"`
	Class           string `json:"class"`
	Characteristics [6]int `json:"characteristics"`
}
