package entity

type Item interface {
	Identifiable
	GetSimpleItem() SimpleItem
}

type SimpleItem struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	OrigName string  `json:"OrigName"`
	Comment  string  `json:"Comment"`
	Price    float64 `json:"Price"`
	Money    int     `json:"Money"`
	Weight   float64 `json:"Weight"`
	HtmlText string  `json:"HtmlText"`
	Tags     string  `json:"Tags"`
}

// Armor Entity
type Armor struct {
	SimpleItem
	ArmorInt      int    `json:"ArmorInt"`
	StealthDis    bool   `json:"StealthDis"`
	StrArmor      int    `json:"StrArmor"`
	FullDexArmor  bool   `json:"FullDexArmor"`
	ShortDexArmor bool   `json:"ShortDexArmor"`
	NoDexArmor    bool   `json:"NoDexArmor"`
	PropertyArmor string `json:"PropertyArmor"`
}

func (item SimpleItem) GetSimpleItem() SimpleItem { return item }
func (item SimpleItem) GetID() string {
	return item.Id
}

// Trinket Entity

type Container struct {
	SimpleItem
	List []Item `json:"List"`
}

func (t Container) GetID() string {
	return t.SimpleItem.Id
}

// Weapon Entity
type Weapon struct {
	SimpleItem
	IsFencing         bool   `json:"IsFencing"`
	WeaponFormula     string `json:"WeaponFormula"`
	WeaponAttackBonus int    `json:"WeaponAttackBonus"`
	PropertyWeapon    string `json:"PropertyWeapon"`
	WeaponDamageType  int    `json:"WeaponDamageType"`
}

func (w Weapon) GetID() string {
	return w.SimpleItem.Id
}

// Ammunition Entity
type Ammunition struct {
	SimpleItem
	IsCounting   bool   `json:"IsCounting"`
	DefaultCount int    `json:"DefaultCount"`
	CustomTags   string `json:"CustomTags"`
}
