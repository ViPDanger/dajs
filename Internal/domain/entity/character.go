package entity

import "time"

type Character struct {
	ID                   string    `json:"Id"`
	Name                 string    `json:"name"`
	MyRaceId             *string   `json:"MyRaceId"`
	UserRace             string    `json:"UserRace"`
	Hold                 bool      `json:"Hold"`
	HoldTime             time.Time `json:"HoldTime"`
	TokenColor           string    `json:"TokenColor"`
	Speed                int       `json:"Speed"`
	IHaveLight           bool      `json:"IHaveLight"`
	TorchValue           int       `json:"TorchValue"`
	TorchValueSecond     int       `json:"TorchValueSecond"`
	CellEyeValue         int       `json:"CellEyeValue"`
	EyeEnabled           bool      `json:"EyeEnabled"`
	SeeInTheDark         bool      `json:"SeeInTheDark"`
	Gold                 int       `json:"Gold"`
	Silver               int       `json:"Silver"`
	Copper               int       `json:"Copper"`
	HitPoints            int       `json:"HitPoints"`
	CurrentHitPoints     int       `json:"CurrentHitPoints"`
	TempHitPoints        int       `json:"TempHitPoints"`
	TempCurrentHitPoints int       `json:"TempCurrentHitPoints"`
	HasInspiration       bool      `json:"HasInspiration"`
	Alignment            int       `json:"Alignment"`
	HitDice              int       `json:"HitDice"`
	HitDiceCount         int       `json:"HitDiceCount"`
	IsArmorTakeOf        bool      `json:"IsArmorTakeOf"`
	TwoHanded            bool      `json:"TwoHanded"`
	SelectedSaveThrowKey int       `json:"SelectedSaveThrowKey"`
	SizeIndex            int       `json:"SizeIndex"`
	TagString            *string   `json:"TagString"`
	Languages            string    `json:"Languages"`
	Multiplier           float64   `json:"Multiplier"`
	Inspiration          int       `json:"Inspiration"`
	Armor                int       `json:"Armor"`
	Bditelnost           int       `json:"Bditelnost"`
	IniBonus             int       `json:"IniBonus"`
	IsPlaying            bool      `json:"IsPlaying"`
	Note                 string    `json:"Note"`
	FirstSpellText       string    `json:"FirstSpellText"`
	SecondSpellText      *string   `json:"SecondSpellText"`
	Spells               []Item    `json:"Spells"` // Можно уточнить, если известен формат заклинаний
	HandsCapacity        int       `json:"HandsCapacity"`
	HandsItems           []Item    `json:"HandsItems"`
	MainHandsItems       []Item    `json:"MainHandsItems"`
	ArrowItems           []Item    `json:"ArrowItems"`

	InventoryItems []InventoryItem `json:"InventoryItems"`
	Parameters     []Parameter     `json:"Parameters"`
	Classes        []Class         `json:"Classes"`

	DamageResist      string         `json:"DamageResist"`
	DamageImmun       string         `json:"DamageImmun"`
	DamageVulner      string         `json:"DamageVulner"`
	HasAura           bool           `json:"HasAura"`
	AuraSize          int            `json:"AuraSize"`
	AuraColor         string         `json:"AuraColor"`
	AuraAngle         float64        `json:"AuraAngle"`
	AuraAngleSize     float64        `json:"AuraAngleSize"`
	AuraOpacity       float64        `json:"AuraOpacity"`
	AuraType          string         `json:"AuraType"`
	AuraColorEnable   bool           `json:"AuraColorEnable"`
	ShowAuraCells     bool           `json:"ShowAuraCells"`
	IsRotationEnable  bool           `json:"IsRotationEnable"`
	IsWallBlock       bool           `json:"IsWallBlock"`
	ShowAuraToPlayers bool           `json:"ShowAuraToPlayers"`
	CustomAuraImage   *string        `json:"CustomAuraImage"`
	CustomStatuses    []CustomStatus `json:"CustomStatuses"`
}

type InventoryItem struct {
	Name           string  `json:"Name"`
	IsInfinity     bool    `json:"IsInfinity"`
	CapacityCount  int     `json:"CapacityCount"`
	CapacityWeight float64 `json:"CapacityWeight"`
	IsBlocked      bool    `json:"IsBlocked"`
	Items          []Item  `json:"Items"`
}

type Item struct {
	SimpleItem string  `json:"SimpleItem"`
	Weight     float64 `json:"Weight"`
	Count      int     `json:"Count"`
}

type Parameter struct {
	Name          string    `json:"Name"`
	Value         int       `json:"Value"`
	UserSpasValue int       `json:"UserSpasValue"`
	Proficiency   bool      `json:"Proficiency"`
	Abilities     []Ability `json:"Abilities"`
}

type Ability struct {
	UserValue   int  `json:"UserValue"`
	MinValue    int  `json:"MinValue"`
	Proficiency bool `json:"Proficiency"`
}

type Class struct {
	Id         string      `json:"Id"`
	Level      int         `json:"Level"`
	SpellCells []SpellCell `json:"SpellCells"`
}

type SpellCell struct {
	Level int `json:"Level"`
	Left  int `json:"Left"`
	Max   int `json:"Max"`
}

type CustomStatus struct {
	IsHided               bool          `json:"IsHided,omitempty"`
	InnerStatusCollection []interface{} `json:"InnerStatusCollection"`
	SelectedModes         []StatusMode  `json:"SelectedModes"`
	Description           string        `json:"Description"`
	TokenPicPath          string        `json:"TokenPicPath"`
	IconPath              string        `json:"IconPath"`
	Id                    string        `json:"Id"`
	Name                  string        `json:"Name"`
}

type StatusMode struct {
	Name        string  `json:"Name"`
	Mode        int     `json:"Mode"`
	RunTimeType int     `json:"RunTimeType"`
	Value       float64 `json:"Value,omitempty"`
	Formula     string  `json:"Formula,omitempty"`
	Ability     string  `json:"Ability,omitempty"`
}
