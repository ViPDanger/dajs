package jsonDTO

type CharacterDTO struct {
	ID               string             `json:"Id"`
	Name             string             `json:"Name"`
	Race             string             `json:"Race"`
	Classes          []ClassDTO         `json:"Classes"`
	Parameters       []ParameterDTO     `json:"Parameters"`
	InventoryItems   []InventoryItemDTO `json:"Inventoryitems"`
	HitPoints        int                `json:"HitPoints"`
	CurrentHitPoints int                `json:"CurrentHitPoints"`
	Armor            int                `json:"Armor"`
	CustomStatuses   []CustomStatusDTO  `json:"CustomStatuses"`
}

type ParameterDTO struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
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

type AbilityDTO struct {
	Name        string `json:"Name,omitempty"`
	UserValue   int    `json:"userValue"`
	MinValue    int    `json:"minValue"`
	Proficiency bool   `json:"proficiency"`
}

type CustomStatusDTO struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Name        string `json:"mame"`
}
