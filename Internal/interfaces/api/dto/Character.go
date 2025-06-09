package dto

type CharacterDTO struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Race             string             `json:"race"`
	Classes          []ClassDTO         `json:"classes"`
	Parameters       []ParameterDTO     `json:"parameters"`
	InventoryItem    []InventoryItemDTO `json:"inventoryitems"`
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
	SimpleItem string  `json:"simpleItem"`
	Weight     float64 `json:"weight"`
	Count      int     `json:"count"`
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
