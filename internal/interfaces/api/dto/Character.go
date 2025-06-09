package dto

type CharacterDTO struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Race             string             `json:"race"`
	Classes          []ClassDTO         `json:"classes"`
	Parameters       []ParameterDTO     `json:"parameters"`
	InventoryItems   []InventoryItemDTO `json:"inventoryitems"`
	HitPoints        int                `json:"hitPoints"`
	CurrentHitPoints int                `json:"currentHitPoints"`
	Armor            int                `json:"armor"`
	CustomStatuses   []CustomStatusDTO  `json:"customStatuses"`
}

type ParameterDTO struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ClassDTO struct {
	Id         string         `json:"id"`
	Level      int            `json:"level"`
	SpellCells []SpellCellDTO `json:"spellCells"`
}

type SpellCellDTO struct {
	Level int `json:"level"`
	Left  int `json:"left"`
	Max   int `json:"max"`
}

type ItemDTO struct {
	Id           string  `json:"Id,omitempty"`
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

type InventoryItemDTO struct {
	Name           string    `json:"name"`
	IsInfinity     bool      `json:"isInfinity"`
	CapacityCount  int       `json:"capacityCount"`
	CapacityWeight float64   `json:"capacityWeight"`
	IsBlocked      bool      `json:"isBlocked"`
	Items          []ItemDTO `json:"items"`
}

type AbilityDTO struct {
	UserValue   int  `json:"userValue"`
	MinValue    int  `json:"minValue"`
	Proficiency bool `json:"proficiency"`
}

type CustomStatusDTO struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Name        string `json:"mame"`
}
