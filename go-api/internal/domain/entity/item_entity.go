package entity

import (
	"errors"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

type Item struct {
	IItem
}

func (i *Item) UnmarshalBSON(data []byte) error {
	var raw map[string]interface{}
	if err := bson.Unmarshal(data, &raw); err != nil {
		return err
	}
	Type, ok := raw["_type"].(string)
	if !ok {
		return errors.New("Cant reach _type")
	}
	delete(raw, "_type")
	typedData, err := bson.Marshal(raw)
	if err != nil {
		return err
	}
	switch Type {
	case "*entity.Armor":
		a := new(Armor)
		err = bson.Unmarshal(typedData, a)
		i.IItem = a
		return err
	case "*entity.SimpleItem":
		a := new(SimpleItem)
		err = bson.Unmarshal(typedData, a)
		i.IItem = a
		return err
	case "*entity.Ammunition":
		a := new(Ammunition)
		err = bson.Unmarshal(typedData, a)
		i.IItem = a
		return err
	case "*entity.Weapon":
		a := new(Weapon)
		err = bson.Unmarshal(typedData, a)
		i.IItem = a
		return err
	case "*entity.Container":
		a := new(Container)
		err = bson.Unmarshal(typedData, a)
		i.IItem = a
		return err
	default:
		return errors.New("Can't UnmarshalBSON data")
	}

}
func (i *Item) MarshalBSON() ([]byte, error) {
	if i.IItem == nil {
		return nil, fmt.Errorf("nil IItem")
	}

	data, err := bson.Marshal(i.IItem)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := bson.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	m["_type"] = reflect.TypeOf(i.IItem).String()
	return bson.Marshal(m)
}

type IItem interface {
	GetSimpleItem() *SimpleItem
}

type SimpleItem struct {
	ID       `bson:"_id"`
	Name     string
	OrigName string
	Comment  string
	Price    float64
	Money    int
	Weight   float64
	HtmlText string
	Tags     []string
}

func (s *SimpleItem) GetSimpleItem() *SimpleItem {
	return s
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
