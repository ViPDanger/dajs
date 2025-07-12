package mapper

import (
	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

func ToPlayerEntity(d dto.PlayerDTO) entity.Player {
	return entity.Player{
		Character:     ToCharacterEntity(d.CharacterDTO),
		CharacterRace: ToCharacterRaceEntity(d.CharacterRaceDTO),
		Background:    ToBackgroundEntity(d.BackgroundDTO),
		Inventory:     ToInventoryEntity(d.InventoryDTO),
		Classes:       ToCharacterClassEntityList(d.Classes),
		Level:         d.Level,
		Experience:    d.Experience,
		SpellSlots:    ToSpellSlotEntityMap(d.SpellSlots),
		Skills:        ToSkillEntityList(d.Skills),
	}
}

func ToPlayerDTO(e entity.Player) dto.PlayerDTO {
	return dto.PlayerDTO{
		CharacterDTO:     ToCharacterDTO(e.Character),
		CharacterRaceDTO: ToCharacterRaceDTO(e.CharacterRace),
		BackgroundDTO:    ToBackgroundDTO(e.Background),
		InventoryDTO:     ToInventoryDTO(e.Inventory),
		Classes:          ToCharacterClassDTOList(e.Classes),
		Level:            e.Level,
		Experience:       e.Experience,
		SpellSlots:       ToSpellSlotDTOMap(e.SpellSlots),
		Skills:           ToSkillDTOList(e.Skills),
	}
}

func ToSpellSlotEntity(d dto.SpellSlotDTO) entity.SpellSlot {
	return entity.SpellSlot{
		Max:     d.Max,
		Current: d.Current,
	}
}

func ToSpellSlotDTO(e entity.SpellSlot) dto.SpellSlotDTO {
	return dto.SpellSlotDTO{
		Max:     e.Max,
		Current: e.Current,
	}
}

func ToSpellSlotEntityMap(m map[int]dto.SpellSlotDTO) map[int]entity.SpellSlot {
	result := make(map[int]entity.SpellSlot, len(m))
	for k, v := range m {
		result[k] = ToSpellSlotEntity(v)
	}
	return result
}

func ToSpellSlotDTOMap(m map[int]entity.SpellSlot) map[int]dto.SpellSlotDTO {
	result := make(map[int]dto.SpellSlotDTO, len(m))
	for k, v := range m {
		result[k] = ToSpellSlotDTO(v)
	}
	return result
}
