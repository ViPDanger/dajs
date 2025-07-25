package entity

type PlayerCharacter struct {
	Character
	CharacterRace `json:"race"`
	Background    `json:"background"`
	Inventory     `json:"inventory"`
	Classes       []CharacterClass  `json:"classes"`
	Level         int               `json:"level"`
	Experience    int               `json:"experience"`
	SpellSlots    map[int]SpellSlot `json:"spell_slots"`
	Skills        []Skill           `json:"skills"`
}
