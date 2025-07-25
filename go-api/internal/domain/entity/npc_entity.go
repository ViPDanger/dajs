package entity

type NPC struct {
	Character  `bson:",inline"`
	Occupation string `json:"occupation"` // Например: "blacksmith", "priest"
	Notes      string `json:"notes"`
}
