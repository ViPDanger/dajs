package jsonMapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
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
