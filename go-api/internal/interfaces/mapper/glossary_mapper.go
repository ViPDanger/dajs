package mapper

import (
	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

func ToGlossaryEntity(d dto.GlossaryDTO) entity.Glossary {
	return entity.Glossary{
		ID:   entity.ID(d.ID),
		Text: d.Text,
	}
}

func ToGlossaryDTO(e entity.Glossary) dto.GlossaryDTO {
	return dto.GlossaryDTO{
		ID:   e.ID.String(),
		Text: e.Text,
	}
}
