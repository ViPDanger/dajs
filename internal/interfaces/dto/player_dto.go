package dto

type PlayerDTO struct {
	CharacterDTO
	CharacterRaceDTO `json:"race"`
	BackgroundDTO    `json:"background"`
	InventoryDTO     `json:"inventory"`
	Classes          []CharacterClassDTO  `json:"classes"`
	Level            int                  `json:"level"`
	Experience       int                  `json:"experience"`
	SpellSlots       map[int]SpellSlotDTO `json:"spell_slots"`
	Skills           []SkillDTO           `json:"skills"`
}
