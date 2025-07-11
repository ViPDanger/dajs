package mapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/interfaces/dto"
)

func ToMonsterEntity(d dto.MonsterDTO) entity.Monster {
	return entity.Monster{
		Character:        ToCharacterEntity(d.CharacterDTO),
		Type:             d.Type,
		Challenge:        d.Challenge,
		Legendary:        d.Legendary,
		Actions:          ToActionEntityList(d.Actions),
		LegendaryActions: ToActionEntityList(d.LegendaryActions),
	}
}

func ToMonsterDTO(e entity.Monster) dto.MonsterDTO {
	return dto.MonsterDTO{
		CharacterDTO:     ToCharacterDTO(e.Character),
		Type:             e.Type,
		Challenge:        e.Challenge,
		Legendary:        e.Legendary,
		Actions:          ToActionDTOList(e.Actions),
		LegendaryActions: ToActionDTOList(e.LegendaryActions),
	}
}

func ToActionEntity(d dto.ActionDTO) entity.Action {
	return entity.Action{
		Name:        d.Name,
		Description: d.Description,
		Damage:      d.Damage,
	}
}

func ToActionDTO(e entity.Action) dto.ActionDTO {
	return dto.ActionDTO{
		Name:        e.Name,
		Description: e.Description,
		Damage:      e.Damage,
	}
}

func ToActionEntityList(list []dto.ActionDTO) []entity.Action {
	result := make([]entity.Action, len(list))
	for i, v := range list {
		result[i] = ToActionEntity(v)
	}
	return result
}

func ToActionDTOList(list []entity.Action) []dto.ActionDTO {
	result := make([]dto.ActionDTO, len(list))
	for i, v := range list {
		result[i] = ToActionDTO(v)
	}
	return result
}
