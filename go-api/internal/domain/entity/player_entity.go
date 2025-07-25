package entity

type PlayerChar struct {
	Character     `bson:",inline"`
	CharacterRace `json:"race"`
	Background    `json:"background"`
	Inventory     `json:"inventory"`
	Classes       []CharacterClass  `json:"classes"`
	Level         int               `json:"level"`
	Experience    int               `json:"experience"`
	SpellSlots    map[int]SpellSlot `json:"spell_slots"`
	Skills        []Skill           `json:"skills"`
}
