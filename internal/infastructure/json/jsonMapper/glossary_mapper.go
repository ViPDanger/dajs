package jsonMapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/infastructure/json/jsonDTO"
)

func ToGlossaryEntity(dto jsonDTO.GlossaryDTO) entity.Glossary {
	return entity.Glossary{
		ID:   dto.ID,
		Text: dto.Text,
	}
}

func ToGlossaryDTO(entity entity.Glossary) jsonDTO.GlossaryDTO {
	return jsonDTO.GlossaryDTO{
		ID:   entity.ID,
		Text: entity.Text,
	}
}
