package jsonMapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
)

// Ability -> AbilityDTO
func ToAbilityDTO(ability entity.Ability) jsonDTO.AbilityDTO {
	return jsonDTO.AbilityDTO{
		Name:        ability.Name,
		UserValue:   ability.UserValue,
		MinValue:    ability.MinValue,
		Proficiency: ability.Proficiency,
	}
}

func ToAbilityEntity(abilityDTO jsonDTO.AbilityDTO) entity.Ability {
	return entity.Ability{
		Name:        abilityDTO.Name,
		UserValue:   abilityDTO.UserValue,
		MinValue:    abilityDTO.MinValue,
		Proficiency: abilityDTO.Proficiency,
	}
}
