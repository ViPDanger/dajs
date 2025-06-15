package mapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
)

// Character -> CharacterDTO
func ToCharacterEntity(dto dto.CharacterDTO) entity.Character {
	return entity.Character{
		ID:               dto.ID,
		Name:             dto.Name,
		UserRace:         dto.Race,
		Classes:          ToClassEntity(dto.Classes),
		Parameters:       ToParametersEntity(dto.Parameters),
		Inventory:        ToInventoryEntityList(dto.InventoryItems),
		HitPoints:        dto.HitPoints,
		CurrentHitPoints: dto.CurrentHitPoints,
		Armor:            dto.Armor,
		CustomStatuses:   ToCustomStatusEntityList(dto.CustomStatuses),
	}
}

func ToCharacterDTO(entity entity.Character) dto.CharacterDTO {
	return dto.CharacterDTO{
		ID:               entity.ID,
		Name:             entity.Name,
		Race:             entity.UserRace,
		Classes:          ToClassDTO(entity.Classes),
		Parameters:       ToParametersDTO(entity.Parameters),
		InventoryItems:   ToInventoryDTOList(entity.Inventory),
		HitPoints:        entity.HitPoints,
		CurrentHitPoints: entity.CurrentHitPoints,
		Armor:            entity.Armor,
		CustomStatuses:   ToCustomStatusDTOList(entity.CustomStatuses),
	}
}

// Parameter -> ParameterDTO
func ToParametersEntity(dtos []dto.ParameterDTO) (r []entity.Parameter) {
	for i := range dtos {
		r = append(r,
			entity.Parameter{
				Name:  dtos[i].Name,
				Value: dtos[i].Value,
			})
	}
	return
}

func ToParametersDTO(entitys []entity.Parameter) (r []dto.ParameterDTO) {
	for i := range entitys {
		r = append(r,
			dto.ParameterDTO{
				Name:  entitys[i].Name,
				Value: entitys[i].Value,
			})
	}
	return
}

// Class -> ClassDTO
func ToClassEntity(dtos []dto.ClassDTO) []entity.Class {
	result := make([]entity.Class, len(dtos))
	for i := range dtos {
		result[i] = entity.Class{
			ID:    dtos[i].Id,
			Level: dtos[i].Level,
		}
	}
	return result
}

func ToClassDTO(entitys []entity.Class) []dto.ClassDTO {
	result := make([]dto.ClassDTO, len(entitys))
	for i := range entitys {
		result[i] = dto.ClassDTO{
			Id:         entitys[i].ID,
			Level:      entitys[i].Level,
			SpellCells: ToSpellCellDTO(entitys[i].SpellCells),
		}
	}
	return result
}

// SpellCell -> SpellCellDTO
func ToSpellCellEntity(dtos []dto.SpellCellDTO) []entity.SpellCell {
	result := make([]entity.SpellCell, len(dtos))
	for i := range dtos {
		result[i] = entity.SpellCell{
			Level: dtos[i].Level,
			Left:  dtos[i].Left,
			Max:   dtos[i].Max,
		}
	}
	return result
}

func ToSpellCellDTO(entitys []entity.SpellCell) []dto.SpellCellDTO {
	result := make([]dto.SpellCellDTO, len(entitys))
	for i := range entitys {
		result[i] = dto.SpellCellDTO{
			Level: entitys[i].Level,
			Left:  entitys[i].Left,
			Max:   entitys[i].Max,
		}
	}
	return result
}

//

// InventoryItem -> InventoryItemDTO
func ToInventoryItemDTO(entitys entity.CharacterInventory) dto.InventoryItemDTO {
	itemsDTO := make([]dto.CharacterItemDTO, len(entitys.Items))
	for i, item := range entitys.Items {
		itemsDTO[i] = ToCharacterItemDTO(item)
	}
	return dto.InventoryItemDTO{
		Name:           entitys.Name,
		IsInfinity:     entitys.IsInfinity,
		IsBlocked:      entitys.IsBlocked,
		CapacityCount:  entitys.CapacityCount,
		CapacityWeight: entitys.CapacityWeight,
		Items:          itemsDTO,
	}
}

func ToInventoryEntity(dto dto.InventoryItemDTO) entity.CharacterInventory {
	items := make([]entity.CharacterItem, len(dto.Items))
	for i, itemDTO := range dto.Items {
		items[i] = ToCharacterItemEntity(itemDTO)
	}
	return entity.CharacterInventory{
		Name:           dto.Name,
		IsInfinity:     dto.IsInfinity,
		IsBlocked:      dto.IsBlocked,
		CapacityCount:  dto.CapacityCount,
		CapacityWeight: dto.CapacityWeight,
		Items:          items,
	}
}

// Ability -> AbilityDTO
func ToAbilityDTO(ability entity.Ability) dto.AbilityDTO {
	return dto.AbilityDTO{
		Name:        ability.Name,
		UserValue:   ability.UserValue,
		MinValue:    ability.MinValue,
		Proficiency: ability.Proficiency,
	}
}

func ToAbilityEntity(abilityDTO dto.AbilityDTO) entity.Ability {
	return entity.Ability{
		Name:        abilityDTO.Name,
		UserValue:   abilityDTO.UserValue,
		MinValue:    abilityDTO.MinValue,
		Proficiency: abilityDTO.Proficiency,
	}
}

// CustomStatus -> CustomStatusDTO
func ToCustomStatusDTO(entity entity.CustomStatus) dto.CustomStatusDTO {
	return dto.CustomStatusDTO{
		Description: entity.Description,
		Id:          entity.ID,
		Name:        entity.Name, // в dto поле называется "mame" — вероятно, опечатка
	}
}

func ToCustomStatusEntity(dto dto.CustomStatusDTO) entity.CustomStatus {
	return entity.CustomStatus{
		Description: dto.Description,
		ID:          dto.Id,
		Name:        dto.Name,
	}
}

// []InventoryItem -> []InventoryItemDTO
func ToInventoryDTOList(entitys []entity.CharacterInventory) []dto.InventoryItemDTO {
	result := make([]dto.InventoryItemDTO, len(entitys))
	for i, item := range entitys {
		result[i] = ToInventoryItemDTO(item)
	}
	return result
}

func ToInventoryEntityList(dtos []dto.InventoryItemDTO) []entity.CharacterInventory {
	result := make([]entity.CharacterInventory, len(dtos))
	for i, itemDTO := range dtos {
		result[i] = ToInventoryEntity(itemDTO)
	}
	return result
}

func ToCharacterItemEntity(dto dto.CharacterItemDTO) entity.CharacterItem {
	return entity.CharacterItem{
		Id:           dto.Id,
		SimpleItem:   dto.SimpleItem,
		Weight:       dto.Weight,
		Count:        dto.Count,
		InventoryId:  dto.InventoryId,
		IsMagicGlow:  dto.IsMagicGlow,
		IsMagicFocus: dto.IsMagicFocus,
		IsRecognized: dto.IsRecognized,
		OnEquip:      dto.OnEquip,
		Focused:      dto.Focused,
	}
}

func ToCharacterItemDTO(item entity.CharacterItem) dto.CharacterItemDTO {
	return dto.CharacterItemDTO{
		Id:           item.Id,
		SimpleItem:   item.SimpleItem,
		Weight:       item.Weight,
		Count:        item.Count,
		InventoryId:  item.InventoryId,
		IsMagicGlow:  item.IsMagicGlow,
		IsMagicFocus: item.IsMagicFocus,
		IsRecognized: item.IsRecognized,
		OnEquip:      item.OnEquip,
		Focused:      item.Focused,
	}
}

// []CustomStatus -> []CustomStatusDTO
func ToCustomStatusDTOList(entitys []entity.CustomStatus) []dto.CustomStatusDTO {
	result := make([]dto.CustomStatusDTO, len(entitys))
	for i, status := range entitys {
		result[i] = ToCustomStatusDTO(status)
	}
	return result
}

func ToCustomStatusEntityList(dtos []dto.CustomStatusDTO) []entity.CustomStatus {
	result := make([]entity.CustomStatus, len(dtos))
	for i, dto := range dtos {
		result[i] = ToCustomStatusEntity(dto)
	}
	return result
}
