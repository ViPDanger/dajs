package entity

type Monster struct {
	Character        `bson:",inline"`
	Type             string   `json:"type"`      // Например: "dragon", "fiend", "beast"
	Challenge        float64  `json:"challenge"` // CR (Challenge Rating)
	Legendary        bool     `json:"legendary"` // Есть ли легендарные действия
	Actions          []Action `json:"actions"`   // Обычные действия
	LegendaryActions []Action `json:"legendary_actions,omitempty"`
}

type Action struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Damage      string `json:"damage,omitempty"` // Например: "2d6+3"
}
