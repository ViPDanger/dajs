package dto

type CharacterDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     CharacterStatusDTO
	Attributes AttributesDTO
	Alignment  string       `json:"alignment"`
	Abilities  []AbilityDTO `json:"abilities"`
	Spells     []SpellDTO   `json:"spells"`
	Tags       []string     `json:"tags"`
}

type CharacterClassDTO struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Level       int                    `json:"level"`
	HitDice     string                 `json:"hit_dice"`
	Abilities   []AbilityDTO           `json:"abilities"`
	ExtraFields map[string]interface{} `json:"extra_fields,omitempty"`
}

type CharacterRaceDTO struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Subrace     string                 `json:"subrace"`
	Traits      []string               `json:"traits"`
	Abilities   []AbilityDTO           `json:"abilities"`
	ExtraFields map[string]interface{} `json:"extra_fields,omitempty"`
}

type BackgroundDTO struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Abilities     []AbilityDTO           `json:"abilities"`
	Proficiencies []string               `json:"proficiencies"`
	ExtraFields   map[string]interface{} `json:"extra_fields,omitempty"`
}

type AbilityDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LevelGained int    `json:"level_gained"`
}

type SpellDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type SpellSlotDTO struct {
	Max     int `json:"max"`
	Current int `json:"current"`
}

type AttributesDTO struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
	SkillsDTO
}

type SkillsDTO struct {
	Acrobatics     SkillDTO `json:"acrobatics"`      // Ловкость
	AnimalHandling SkillDTO `json:"animal_handling"` // Мудрость
	Arcana         SkillDTO `json:"arcana"`          // Интеллект
	Athletics      SkillDTO `json:"athletics"`       // Сила
	Deception      SkillDTO `json:"deception"`       // Харизма
	History        SkillDTO `json:"history"`         // Интеллект
	Insight        SkillDTO `json:"insight"`         // Мудрость
	Intimidation   SkillDTO `json:"intimidation"`    // Харизма
	Investigation  SkillDTO `json:"investigation"`   // Интеллект
	Medicine       SkillDTO `json:"medicine"`        // Мудрость
	Nature         SkillDTO `json:"nature"`          // Интеллект
	Perception     SkillDTO `json:"perception"`      // Мудрость
	Performance    SkillDTO `json:"performance"`     // Харизма
	Persuasion     SkillDTO `json:"persuasion"`      // Харизма
	Religion       SkillDTO `json:"religion"`        // Интеллект
	SleightOfHand  SkillDTO `json:"sleight_of_hand"` // Ловкость
	Stealth        SkillDTO `json:"stealth"`         // Ловкость
	Survival       SkillDTO `json:"survival"`        // Мудрость
}

type SkillDTO struct {
	Proficient bool `json:"proficient"` // Владеет ли персонаж
	Modifier   int  `json:"modifier"`   // Модификатор: атрибут + proficiency (если есть)
}

// Инвентарь со всеми предметами, весом и валютой
type InventoryDTO struct {
	Items       []ItemDTO   `json:"items"`
	TotalWeight float64     `json:"total_weight"`
	Currency    CurrencyDTO `json:"currency"`
}

type CurrencyDTO struct {
	Copper int `json:"copper"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

type CharacterStatusDTO struct {
	HP          int `json:"hp"`
	MaxHP       int `json:"max_hp"`
	TemporaryHP int `json:"temporary_hp"`
	ArmorClass  int `json:"armor_class"`
	Speed       int `json:"speed"`
	Initiative  int `json:"initiative"`
}
