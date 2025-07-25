package mapper

import (
	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

func ToNPCEntity(d dto.NPCdto) entity.NPC {
	return entity.NPC{
		Character:  ToCharacterEntity(d.CharacterDTO),
		Occupation: d.Occupation,
		Notes:      d.Notes,
	}
}

func ToNPCdto(e entity.NPC) dto.NPCdto {
	return dto.NPCdto{
		CharacterDTO: ToCharacterDTO(e.Character),
		Occupation:   e.Occupation,
		Notes:        e.Notes,
	}
}
