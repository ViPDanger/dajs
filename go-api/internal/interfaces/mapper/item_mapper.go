package mapper

import (
	"strings"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

func ToItemEntity(dto dto.ItemDTO) entity.Item {
	switch {
	case containsTag(dto.Tags, "Тяжелый доспех"), containsTag(dto.Tags, "Легкий доспех"), containsTag(dto.Tags, "Средний доспех"):
		return entity.Item{
			IItem: &entity.Armor{
				SimpleItem:    *ToSimpleItemEntity(dto),
				ArmorInt:      dto.ArmorInt,
				StealthDis:    dto.StealthDis,
				StrArmor:      dto.StrArmor,
				FullDexArmor:  dto.FullDexArmor,
				ShortDexArmor: dto.ShortDexArmor,
				NoDexArmor:    dto.NoDexArmor,
				PropertyArmor: dto.PropertyArmor,
			}}

	case containsTag(dto.Tags, "Безделушка"):
		list := make([]entity.Item, len(dto.List))
		for i := range dto.List {
			list[i] = ToItemEntity(dto.List[i])
		}
		return entity.Item{
			IItem: &entity.Container{
				SimpleItem: *ToSimpleItemEntity(dto),
				List:       list,
			}}

	case containsTag(dto.Tags, "Оружие"):
		return entity.Item{
			IItem: &entity.Weapon{
				SimpleItem:        *ToSimpleItemEntity(dto),
				IsFencing:         dto.IsFencing,
				WeaponFormula:     dto.WeaponFormula,
				WeaponAttackBonus: dto.WeaponAttackBonus,
				PropertyWeapon:    dto.PropertyWeapon,
				WeaponDamageType:  dto.WeaponDamageType,
			}}

	case containsTag(dto.Tags, "Снаряжение"):
		return entity.Item{
			IItem: &entity.Ammunition{
				SimpleItem:   *ToSimpleItemEntity(dto),
				IsCounting:   dto.IsCounting,
				DefaultCount: dto.DefaultCount,
				CustomTags:   dto.CustomTags,
			}}

	case containsTag(dto.Tags, "Инструмент"), containsTag(dto.Tags, "Весовые товары"):
		return entity.Item{
			IItem: ToSimpleItemEntity(dto)}
	default:
		return entity.Item{
			IItem: ToSimpleItemEntity(dto)}
	}

}

func ToItemDTO(e entity.Item) dto.ItemDTO {
	base := toSimpleItemDTO(e.GetSimpleItem())
	switch v := e.IItem.(type) {
	case *entity.Armor:
		base.ArmorInt = v.ArmorInt
		base.StealthDis = v.StealthDis
		base.StrArmor = v.StrArmor
		base.FullDexArmor = v.FullDexArmor
		base.ShortDexArmor = v.ShortDexArmor
		base.NoDexArmor = v.NoDexArmor
		base.PropertyArmor = v.PropertyArmor
	case *entity.Weapon:
		base.IsFencing = v.IsFencing
		base.WeaponFormula = v.WeaponFormula
		base.WeaponAttackBonus = v.WeaponAttackBonus
		base.PropertyWeapon = v.PropertyWeapon
		base.WeaponDamageType = v.WeaponDamageType
	case *entity.Ammunition:
		base.IsCounting = v.IsCounting
		base.DefaultCount = v.DefaultCount
		base.CustomTags = v.CustomTags
	case *entity.Container:
		subDTOs := make([]dto.ItemDTO, len(v.List))
		for i, item := range v.List {
			subDTOs[i] = ToItemDTO(item)
		}
		base.List = subDTOs
	}

	return base
}

func ToItemEntityList(dtos []dto.ItemDTO) []entity.Item {
	r := make([]entity.Item, len(dtos))
	for i := range r {
		r[i] = ToItemEntity(dtos[i])
	}
	return r
}

func ToItemDTOList(dtos []entity.Item) []dto.ItemDTO {
	r := make([]dto.ItemDTO, len(dtos))
	for i := range r {
		r[i] = ToItemDTO(dtos[i])
	}
	return r
}

func ToSimpleItemEntity(dto dto.ItemDTO) *entity.SimpleItem {
	return &entity.SimpleItem{
		ID:       entity.ID(dto.Id),
		Name:     dto.Name,
		OrigName: dto.OrigName,
		Comment:  dto.Comment,
		Price:    dto.Price,
		Money:    dto.Money,
		Weight:   dto.Weight,
		HtmlText: dto.HtmlText,
		Tags:     dto.Tags,
	}
}

func toSimpleItemDTO(i *entity.SimpleItem) dto.ItemDTO {
	return dto.ItemDTO{
		Id:       string(i.ID),
		Name:     i.Name,
		OrigName: i.OrigName,
		Comment:  i.Comment,
		Price:    i.Price,
		Money:    i.Money,
		Weight:   i.Weight,
		HtmlText: i.HtmlText,
		Tags:     i.Tags,
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
