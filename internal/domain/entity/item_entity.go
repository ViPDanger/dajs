package entity

type Item interface {
	Identifiable
	GetSimpleItem() *SimpleItem
}

type SimpleItem struct {
	ID       string
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

func (item SimpleItem) GetID() string {
	return item.ID
}

// Trinket Entity

type Container struct {
	SimpleItem
	List []Item
}

func (t Container) GetID() string {
	return t.SimpleItem.ID
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

func (w Weapon) GetID() string {
	return w.SimpleItem.ID
}

// Ammunition Entity
type Ammunition struct {
	SimpleItem
	IsCounting   bool
	DefaultCount int
	CustomTags   string
}
