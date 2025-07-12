package dto

type MonsterDTO struct {
	CharacterDTO
	Type             string      `json:"type"`      // Например: "dragon", "fiend", "beast"
	Challenge        float64     `json:"challenge"` // CR (Challenge Rating)
	Legendary        bool        `json:"legendary"` // Есть ли легендарные действия
	Actions          []ActionDTO `json:"actions"`   // Обычные действия
	LegendaryActions []ActionDTO `json:"legendary_actions,omitempty"`
}

type ActionDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Damage      string `json:"damage,omitempty"` // Например: "2d6+3"
}
