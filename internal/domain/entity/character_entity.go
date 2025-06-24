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
	Id           *string
	Item         Item
	SimpleItem   string
	Weight       float64
	Count        int
	InventoryId  string
	IsMagicGlow  bool
	IsMagicFocus bool
	IsRecognized bool
	OnEquip      bool
	Focused      bool
}

type Character struct {
	ID           string
	Name         string
	UserRace     string
	Speed        int
	EyeEnabled   bool
	SeeInTheDark bool
	Money
	HitPoints            int
	CurrentHitPoints     int
	TempHitPoints        int
	TempCurrentHitPoints int
	Alignment            int
	HitDice              int
	HitDiceCount         int
	SizeIndex            int
	TagString            *string
	Languages            string
	Multiplier           float64
	Inspiration          int
	Armor                int
	IniBonus             int
	IsPlaying            bool
	Note                 string
	Spells               []Spell
	Inventory            []CharacterInventory
	Parameters           []Parameter
	Classes              []Class
	DamageResist         string
	DamageImmun          string
	DamageVulner         string
	CustomStatuses       []CustomStatus
}

func (c Character) GetID() string {
	return c.ID
}

type Money struct {
	Gold   int
	Silver int
	Copper int
}

type Parameter struct {
	Name          string
	Value         int
	UserSpasValue int
	Proficiency   bool
	Abilities     []Ability
}

type Ability struct {
	Name        string
	UserValue   int
	MinValue    int
	Proficiency bool
}

type Class struct {
	ID         string
	Level      int
	SpellCells []SpellCell
}

func (c Class) GetID() string {
	return c.ID
}

type SpellCell struct {
	Level int
	Left  int
	Max   int
}

type CustomStatus struct {
	ID                    string
	Name                  string
	IsHided               bool
	InnerStatusCollection []interface{}
	SelectedModes         []StatusMode
	Description           string
	TokenPicPath          string
	IconPath              string
}

func (c CustomStatus) GetID() string {
	return c.ID
}

type StatusMode struct {
	Name        string
	Mode        int
	RunTimeType int
	Value       float64
	Formula     string
	Ability     string
}

type Spell struct {
	Id      string
	Name    string
	Ability string
}
