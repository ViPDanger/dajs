package jsonDTO

type ItemDTO struct {
	Id                string    `json:"Id,omitempty"`
	Name              string    `json:"Name,omitempty"`
	OrigName          string    `json:"OrigName,omitempty"`
	Comment           string    `json:"Comment,omitempty"`
	Price             float64   `json:"Price,omitempty"`
	Money             int       `json:"Money,omitempty"`
	Weight            float64   `json:"Weight,omitempty"`
	HtmlText          string    `json:"HtmlText,omitempty"`
	Tags              string    `json:"Tags,omitempty"`
	ArmorInt          int       `json:"ArmorInt,omitempty"`
	StealthDis        bool      `json:"StealthDis,omitempty"`
	StrArmor          int       `json:"StrArmor,omitempty"`
	FullDexArmor      bool      `json:"FullDexArmor,omitempty"`
	ShortDexArmor     bool      `json:"ShortDexArmor,omitempty"`
	NoDexArmor        bool      `json:"NoDexArmor,omitempty"`
	PropertyArmor     string    `json:"PropertyArmor,omitempty"`
	IsCounting        bool      `json:"IsCounting,omitempty"`
	DefaultCount      int       `json:"DefaultCount,omitempty"`
	CustomTags        string    `json:"CustomTags,omitempty"`
	IsFencing         bool      `json:"IsFencing,omitempty"`
	WeaponFormula     string    `json:"WeaponFormula,omitempty"`
	WeaponAttackBonus int       `json:"WeaponAttackBonus,omitempty"`
	PropertyWeapon    string    `json:"PropertyWeapon,omitempty"`
	WeaponDamageType  int       `json:"WeaponDamageType,omitempty"`
	HasDescription    bool      `json:"HasDescription,omitempty"`
	List              []ItemDTO `json:"List,omitempty"`
}
