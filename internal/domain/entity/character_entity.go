package entity

type CharacterInventory struct {
	Name           string
	IsInfinity     bool
	CapacityCount  int
	CapacityWeight float64
	IsBlocked      bool
	Items          []CharacterItem
}

type CharacterItem struct {
	Id           *string `json:"Id,omitempty"`
	Item         Item
	SimpleItem   string  `json:"SimpleItem,omitempty"`
	Weight       float64 `json:"Weight,omitempty"`
	Count        int     `json:"Count"`
	InventoryId  string  `json:"InventoryId,omitempty"`
	IsMagicGlow  bool    `json:"IsMagicGlow,omitempty"`
	IsMagicFocus bool    `json:"IsMagicFocus,omitempty"`
	IsRecognized bool    `json:"IsRecognized,omitempty"`
	OnEquip      bool    `json:"OnEquip,omitempty"`
	Focused      bool    `json:"Focused,omitempty"`
}

type Character struct {
	ID           string `json:"Id"`
	Name         string `json:"name"`
	UserRace     string `json:"UserRace"`
	Speed        int    `json:"Speed"`
	EyeEnabled   bool   `json:"EyeEnabled"`
	SeeInTheDark bool   `json:"SeeInTheDark"`
	Money
	HitPoints            int          `json:"HitPoints"`
	CurrentHitPoints     int          `json:"CurrentHitPoints"`
	TempHitPoints        int          `json:"TempHitPoints"`
	TempCurrentHitPoints int          `json:"TempCurrentHitPoints"`
	Alignment            int          `json:"Alignment"`
	HitDice              int          `json:"HitDice"`
	HitDiceCount         int          `json:"HitDiceCount"`
	SizeIndex            int          `json:"SizeIndex"`
	TagString            *string      `json:"TagString"`
	Languages            string       `json:"Languages"`
	Multiplier           float64      `json:"Multiplier"`
	Inspiration          int          `json:"Inspiration"`
	Armor                int          `json:"Armor"`
	IniBonus             int          `json:"IniBonus"`
	IsPlaying            bool         `json:"IsPlaying"`
	Note                 string       `json:"Note"`
	Spells               []SimpleItem `json:"Spells"` // Можно уточнить, если известен формат заклинаний
	Inventory            []CharacterInventory
	Parameters           []Parameter    `json:"Parameters"`
	Classes              []Class        `json:"Classes"`
	DamageResist         string         `json:"DamageResist"`
	DamageImmun          string         `json:"DamageImmun"`
	DamageVulner         string         `json:"DamageVulner"`
	CustomStatuses       []CustomStatus `json:"CustomStatuses"`
}

func (c Character) GetID() string {
	return c.ID
}

type Money struct {
	Gold   int `json:"Gold"`
	Silver int `json:"Silver"`
	Copper int `json:"Copper"`
}
type Parameter struct {
	Name          string    `json:"Name"`
	Value         int       `json:"Value"`
	UserSpasValue int       `json:"UserSpasValue"`
	Proficiency   bool      `json:"Proficiency"`
	Abilities     []Ability `json:"Abilities"`
}

type Ability struct {
	Name        string `json:"Name,omitempty"`
	UserValue   int    `json:"UserValue"`
	MinValue    int    `json:"MinValue"`
	Proficiency bool   `json:"Proficiency"`
}

type Class struct {
	ID         string      `json:"Id"`
	Level      int         `json:"Level"`
	SpellCells []SpellCell `json:"SpellCells"`
}

func (c Class) GetID() string {
	return c.ID
}

type SpellCell struct {
	Level int `json:"Level"`
	Left  int `json:"Left"`
	Max   int `json:"Max"`
}

type CustomStatus struct {
	ID                    string        `json:"Id"`
	Name                  string        `json:"Name"`
	IsHided               bool          `json:"IsHided,omitempty"`
	InnerStatusCollection []interface{} `json:"InnerStatusCollection"`
	SelectedModes         []StatusMode  `json:"SelectedModes"`
	Description           string        `json:"Description"`
	TokenPicPath          string        `json:"TokenPicPath"`
	IconPath              string        `json:"IconPath"`
}

func (c CustomStatus) GetID() string {
	return c.ID
}

type StatusMode struct {
	Name        string  `json:"Name"`
	Mode        int     `json:"Mode"`
	RunTimeType int     `json:"RunTimeType"`
	Value       float64 `json:"Value,omitempty"`
	Formula     string  `json:"Formula,omitempty"`
	Ability     string  `json:"Ability,omitempty"`
}
