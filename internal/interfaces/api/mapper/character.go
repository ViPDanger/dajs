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
		InventoryItems:   ToInventoryItemEntityList(dto.InventoryItems),
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
		InventoryItems:   ToInventoryItemDTOList(entity.InventoryItems),
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

// Item -> ItemDTO
func ToItemDTO(entity entity.Item) dto.ItemDTO {
	return dto.ItemDTO{
		SimpleItem: entity.SimpleItem,
		Weight:     entity.Weight,
		Count:      entity.Count,
	}
}

func ToItemEntity(dto dto.ItemDTO) entity.Item {
	return entity.Item{
		SimpleItem: dto.SimpleItem,
		Weight:     dto.Weight,
		Count:      dto.Count,
	}
}

// InventoryItem -> InventoryItemDTO
func ToInventoryItemDTO(entitys entity.InventoryItem) dto.InventoryItemDTO {
	itemsDTO := make([]dto.ItemDTO, len(entitys.Items))
	for i, item := range entitys.Items {
		itemsDTO[i] = ToItemDTO(item)
	}
	return dto.InventoryItemDTO{
		Name:           entitys.Name,
		IsInfinity:     entitys.IsInfinity,
		CapacityCount:  entitys.CapacityCount,
		CapacityWeight: entitys.CapacityWeight,
		IsBlocked:      entitys.IsBlocked,
		Items:          itemsDTO,
	}
}

func ToInventoryItemEntity(dtos dto.InventoryItemDTO) entity.InventoryItem {
	items := make([]entity.Item, len(dtos.Items))
	for i, itemDTO := range dtos.Items {
		items[i] = ToItemEntity(itemDTO)
	}
	return entity.InventoryItem{
		Name:           dtos.Name,
		IsInfinity:     dtos.IsInfinity,
		CapacityCount:  dtos.CapacityCount,
		CapacityWeight: dtos.CapacityWeight,
		IsBlocked:      dtos.IsBlocked,
		Items:          items,
	}
}

// Ability -> AbilityDTO
func ToAbilityDTO(ability entity.Ability) dto.AbilityDTO {
	return dto.AbilityDTO{
		UserValue:   ability.UserValue,
		MinValue:    ability.MinValue,
		Proficiency: ability.Proficiency,
	}
}

func ToAbilityEntity(abilityDTO dto.AbilityDTO) entity.Ability {
	return entity.Ability{
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
func ToInventoryItemDTOList(entitys []entity.InventoryItem) []dto.InventoryItemDTO {
	result := make([]dto.InventoryItemDTO, len(entitys))
	for i, item := range entitys {
		result[i] = ToInventoryItemDTO(item)
	}
	return result
}

func ToInventoryItemEntityList(dtos []dto.InventoryItemDTO) []entity.InventoryItem {
	result := make([]entity.InventoryItem, len(dtos))
	for i, itemDTO := range dtos {
		result[i] = ToInventoryItemEntity(itemDTO)
	}
	return result
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
