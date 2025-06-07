package mapper

import (
	"DAJ/Internal/domain/entity"
	"DAJ/Internal/interfaces/api/dto"
)

func ToGlossaryEntity(dto dto.GlossaryDTO) entity.Glossary {
	return entity.Glossary{
		ID:   dto.ID,
		Text: dto.Text,
	}
}

func ToGlossaryDTO(entity entity.Glossary) dto.GlossaryDTO {
	return dto.GlossaryDTO{
		ID:   entity.ID,
		Text: entity.Text,
	}
}
