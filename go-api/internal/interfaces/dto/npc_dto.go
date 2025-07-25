package dto

type NPCdto struct {
	CharacterDTO
	Occupation string `json:"occupation"` // Например: "blacksmith", "priest"
	Notes      string `json:"notes"`
}
