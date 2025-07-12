package dto

type ItemDTO struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	OrigName string  `json:"OrigName"`
	Comment  string  `json:"Comment"`
	Price    float64 `json:"Price"`
	Money    int     `json:"Money"`
	Weight   float64 `json:"Weight"`
	HtmlText string  `json:"HtmlText"`
	Tags     string  `json:"Tags"`

	ArmorInt      int    `json:"ArmorInt"`
	StealthDis    bool   `json:"StealthDis"`
	StrArmor      int    `json:"StrArmor"`
	FullDexArmor  bool   `json:"FullDexArmor"`
	ShortDexArmor bool   `json:"ShortDexArmor"`
	NoDexArmor    bool   `json:"NoDexArmor"`
	PropertyArmor string `json:"PropertyArmor"`

	IsCounting   bool   `json:"IsCounting"`
	DefaultCount int    `json:"DefaultCount"`
	CustomTags   string `json:"CustomTags"`

	IsFencing         bool   `json:"IsFencing"`
	WeaponFormula     string `json:"WeaponFormula"`
	WeaponAttackBonus int    `json:"WeaponAttackBonus"`
	PropertyWeapon    string `json:"PropertyWeapon"`
	WeaponDamageType  int    `json:"WeaponDamageType"`

	HasDescription bool      `json:"HasDescription"`
	List           []ItemDTO `json:"List"`
}
