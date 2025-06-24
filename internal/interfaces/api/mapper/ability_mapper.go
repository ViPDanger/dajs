package mapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
)

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
