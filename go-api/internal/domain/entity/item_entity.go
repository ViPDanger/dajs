package entity

type Item interface {
	Identifiable
	GetSimpleItem() *SimpleItem
}

type SimpleItem struct {
	ID
	Name     string
	OrigName string
	Comment  string
	Price    float64
	Money    int
	Weight   float64
	HtmlText string
	Tags     []string
}

// Armor Entity
type Armor struct {
	SimpleItem
	ArmorInt      int
	StealthDis    bool
	StrArmor      int
	FullDexArmor  bool
	ShortDexArmor bool
	NoDexArmor    bool
	PropertyArmor string
}

func (item SimpleItem) GetSimpleItem() *SimpleItem {
	return &item
}

// Trinket Entity

type Container struct {
	SimpleItem
	List []Item
}

// Weapon Entity
type Weapon struct {
	SimpleItem
	IsFencing         bool
	WeaponFormula     string
	WeaponAttackBonus int
	PropertyWeapon    string
	WeaponDamageType  int
}

// Ammunition Entity
type Ammunition struct {
	SimpleItem
	IsCounting   bool
	DefaultCount int
	CustomTags   string
}
