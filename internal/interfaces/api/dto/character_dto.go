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
	CustomStatuses   []StatusDTO        `json:"customStatuses"`
}

type ParameterDTO struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ClassDTO struct {
	Id         string         `json:"Id"`
	Level      int            `json:"Level"`
	SpellCells []SpellCellDTO `json:"SpellCells"`
}

type SpellCellDTO struct {
	Level int `json:"Level"`
	Left  int `json:"Left"`
	Max   int `json:"Max"`
}

type CharacterItemDTO struct {
	Id           *string `json:"Id,omitempty"`
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
	Name           string             `json:"name"`
	IsInfinity     bool               `json:"isInfinity"`
	CapacityCount  int                `json:"capacityCount"`
	CapacityWeight float64            `json:"capacityWeight"`
	IsBlocked      bool               `json:"isBlocked"`
	Items          []CharacterItemDTO `json:"items"`
}
