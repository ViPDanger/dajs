package mapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
	"strings"
)

func ToItemEntity(dto dto.ItemDTO) entity.Item {
	tags := strings.Split(dto.Tags, "|")

	switch {
	case containsTag(tags, "Тяжелый доспех"), containsTag(tags, "Легкий доспех"), containsTag(tags, "Средний доспех"):
		return entity.Armor{
			SimpleItem:    ToSimpleItemEntity(dto),
			ArmorInt:      dto.ArmorInt,
			StealthDis:    dto.StealthDis,
			StrArmor:      dto.StrArmor,
			FullDexArmor:  dto.FullDexArmor,
			ShortDexArmor: dto.ShortDexArmor,
			NoDexArmor:    dto.NoDexArmor,
			PropertyArmor: dto.PropertyArmor,
		}
	case containsTag(tags, "Безделушка"):
		list := make([]entity.Item, len(dto.List))
		for i := range dto.List {
			list[i] = ToItemEntity(dto.List[i])
		}
		return entity.Container{
			SimpleItem: ToSimpleItemEntity(dto),
			List:       list,
		}

	case containsTag(tags, "Оружие"):
		return entity.Weapon{
			SimpleItem:        ToSimpleItemEntity(dto),
			IsFencing:         dto.IsFencing,
			WeaponFormula:     dto.WeaponFormula,
			WeaponAttackBonus: dto.WeaponAttackBonus,
			PropertyWeapon:    dto.PropertyWeapon,
			WeaponDamageType:  dto.WeaponDamageType,
		}

	case containsTag(tags, "Снаряжение"):
		return entity.Ammunition{
			SimpleItem:   ToSimpleItemEntity(dto),
			IsCounting:   dto.IsCounting,
			DefaultCount: dto.DefaultCount,
			CustomTags:   dto.CustomTags,
		}

	case containsTag(tags, "Инструмент"), containsTag(tags, "Весовые товары"):
		return ToSimpleItemEntity(dto)
	default:
		return ToSimpleItemEntity(dto)
	}

}

func ToItemDTO(e entity.Item) dto.ItemDTO {
	base := toSimpleItemDTO(*e.GetSimpleItem())
	switch v := e.(type) {
	case entity.Armor:
		base.ArmorInt = v.ArmorInt
		base.StealthDis = v.StealthDis
		base.StrArmor = v.StrArmor
		base.FullDexArmor = v.FullDexArmor
		base.ShortDexArmor = v.ShortDexArmor
		base.NoDexArmor = v.NoDexArmor
		base.PropertyArmor = v.PropertyArmor
	case entity.Weapon:
		base.IsFencing = v.IsFencing
		base.WeaponFormula = v.WeaponFormula
		base.WeaponAttackBonus = v.WeaponAttackBonus
		base.PropertyWeapon = v.PropertyWeapon
		base.WeaponDamageType = v.WeaponDamageType
	case entity.Ammunition:
		base.IsCounting = v.IsCounting
		base.DefaultCount = v.DefaultCount
		base.CustomTags = v.CustomTags
	case entity.Container:
		subDTOs := make([]dto.ItemDTO, len(v.List))
		for i, item := range v.List {
			subDTOs[i] = ToItemDTO(item)
		}
		base.List = subDTOs
	}

	return base
}

func ToSimpleItemEntity(dto dto.ItemDTO) entity.SimpleItem {
	return entity.SimpleItem{
		Id:       dto.Id,
		Name:     dto.Name,
		OrigName: dto.OrigName,
		Comment:  dto.Comment,
		Price:    dto.Price,
		Money:    dto.Money,
		Weight:   dto.Weight,
		HtmlText: dto.HtmlText,
		Tags:     strings.Split(dto.Tags, "|"),
	}
}

func toSimpleItemDTO(i entity.SimpleItem) dto.ItemDTO {
	return dto.ItemDTO{
		Id:       i.Id,
		Name:     i.Name,
		OrigName: i.OrigName,
		Comment:  i.Comment,
		Price:    i.Price,
		Money:    i.Money,
		Weight:   i.Weight,
		HtmlText: i.HtmlText,
		Tags:     strings.Join(i.Tags, "|"),
	}
}

func containsTag(tags []string, target string) bool {
	for _, t := range tags {
		if strings.TrimSpace(t) == target {
			return true
		}
	}
	return false
}
