package entity

type Alignment string

const (
	LawfullGood    Alignment = "Lawfull Good"
	NeutralGood    Alignment = "Neutral Good"
	ChaoticGood    Alignment = "Chaotic Good"
	LawfullNeutral Alignment = "Lawfull Neutral"
	TrueNeutral    Alignment = "True Neutral"
	ChaoticNeutral Alignment = "Chaotic Neutral"
	LawfullEvil    Alignment = "Lawfull Evil"
	NeutralEvil    Alignment = "Neutral Evil"
	ChaoticEvil    Alignment = "Chaotic Evil"
)

type Character struct {
	ID         ID              `json:"_id" bson:"_id"`
	CreatorID  string          `json:"creator_id" bson:"creator_id"`
	Name       string          `json:"name"`
	Status     CharacterStatus `json:"status"`
	Attributes Attributes      `json:"attributes"`
	Alignment  Alignment       `json:"alignment"`
	Abilities  []Ability       `json:"abilities"`
	Spells     []Spell         `json:"spells"`
	Tags       []string        `json:"tags"`
}

type CharacterClass struct {
	ID          ID                     `json:"_id"`
	Name        string                 `json:"name"`
	Level       int                    `json:"level"`
	HitDice     string                 `json:"hit_dice"`
	Abilities   []Ability              `json:"abilities"`
	ExtraFields map[string]interface{} `json:"extra_fields,omitempty"`
}

type CharacterRace struct {
	ID          ID                     `json:"_id"`
	Name        string                 `json:"name"`
	Subrace     string                 `json:"subrace"`
	Traits      []string               `json:"traits"`
	Abilities   []Ability              `json:"abilities"`
	ExtraFields map[string]interface{} `json:"extra_fields,omitempty"`
}

type Background struct {
	ID            ID                     `json:"_id" `
	Name          string                 `json:"name"`
	Abilities     []Ability              `json:"abilities"`
	Proficiencies []string               `json:"proficiencies"`
	ExtraFields   map[string]interface{} `json:"extra_fields,omitempty"`
}

type Ability struct {
	ID          ID     `json:"_id" `
	Name        string `json:"name"`
	Description string `json:"description"`
	LevelGained int    `json:"level_gained"`
}

type Spell struct {
	ID          ID     `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type SpellSlot struct {
	Max     int `json:"max"`
	Current int `json:"current"`
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
	Skills
}

type Skills struct {
	Acrobatics     Skill `json:"acrobatics"`      // Ловкость
	AnimalHandling Skill `json:"animal_handling"` // Мудрость
	Arcana         Skill `json:"arcana"`          // Интеллект
	Athletics      Skill `json:"athletics"`       // Сила
	Deception      Skill `json:"deception"`       // Харизма
	History        Skill `json:"history"`         // Интеллект
	Insight        Skill `json:"insight"`         // Мудрость
	Intimidation   Skill `json:"intimidation"`    // Харизма
	Investigation  Skill `json:"investigation"`   // Интеллект
	Medicine       Skill `json:"medicine"`        // Мудрость
	Nature         Skill `json:"nature"`          // Интеллект
	Perception     Skill `json:"perception"`      // Мудрость
	Performance    Skill `json:"performance"`     // Харизма
	Persuasion     Skill `json:"persuasion"`      // Харизма
	Religion       Skill `json:"religion"`        // Интеллект
	SleightOfHand  Skill `json:"sleight_of_hand"` // Ловкость
	Stealth        Skill `json:"stealth"`         // Ловкость
	Survival       Skill `json:"survival"`        // Мудрость
}

type Skill struct {
	Proficient bool `json:"proficient"` // Владеет ли персонаж
	Modifier   int  `json:"modifier"`   // Модификатор: атрибут + proficiency (если есть)
}

// Инвентарь со всеми предметами, весом и валютой
type Inventory struct {
	Items       []Item   `json:"items"`
	TotalWeight float64  `json:"total_weight"`
	Currency    Currency `json:"currency"`
}

type Currency struct {
	Copper int `json:"copper"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

type CharacterStatus struct {
	HP          int `json:"hp"`
	MaxHP       int `json:"max_hp"`
	TemporaryHP int `json:"temporary_hp"`
	ArmorClass  int `json:"armor_class"`
	Speed       int `json:"speed"`
	Initiative  int `json:"initiative"`
}
