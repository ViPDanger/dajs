package entity

type NPC struct {
	Character
	Occupation string `json:"occupation"` // Например: "blacksmith", "priest"
	Notes      string `json:"notes"`
}
