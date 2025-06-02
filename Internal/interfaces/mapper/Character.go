package mapper

import (
	"DAJ/Internal/entity"
	"DAJ/Internal/interfaces/api/dto"
	"fmt"
)

func CharacterDTOtoCharacter(dto dto.CharacterDTO) entity.Character {
	fmt.Println("dto to char")
	return entity.Character{
		ID:       dto.ID,
		Name:     dto.Name,
		UserRace: dto.Race,
		Parameters: []entity.Parameter{
			{
				Name:  "Strenght",
				Value: dto.Characteristics[0],
			},
			{
				Name:  "Dexterity",
				Value: dto.Characteristics[1],
			},
		},
	}
}

func CharacterToCharacterDTO(entity entity.Character) dto.CharacterDTO {
	fmt.Println("char to dto")
	return dto.CharacterDTO{
		ID:              entity.ID,
		Name:            entity.Name,
		Race:            entity.UserRace,
		Characteristics: [6]int{10, 10, 10, 10, 10, 10},
	}
}
