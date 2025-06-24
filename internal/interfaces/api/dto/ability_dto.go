package dto

type AbilityDTO struct {
	Name        string `json:"Name,omitempty"`
	UserValue   int    `json:"userValue"`
	MinValue    int    `json:"minValue"`
	Proficiency bool   `json:"proficiency"`
}
