package mapper

import (
	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

// CHARACTER
func ToCharacterEntity(d dto.CharacterDTO) entity.Character {
	return entity.Character{
		ID:         d.ID,
		Name:       d.Name,
		Status:     ToCharacterStatusEntity(d.Status),
		Attributes: ToAttributesEntity(d.Attributes),
		Alignment:  entity.Alignment(d.Alignment),
		Abilities:  ToAbilityEntityList(d.Abilities),
		Spells:     ToSpellEntityList(d.Spells),
		Tags:       d.Tags,
	}
}

func ToCharacterDTO(e entity.Character) dto.CharacterDTO {
	return dto.CharacterDTO{
		ID:         e.ID,
		Name:       e.Name,
		Status:     ToCharacterStatusDTO(e.Status),
		Attributes: ToAttributesDTO(e.Attributes),
		Alignment:  string(e.Alignment),
		Abilities:  ToAbilityDTOList(e.Abilities),
		Spells:     ToSpellDTOList(e.Spells),
		Tags:       e.Tags,
	}
}

// CHARACTER CLASS

func ToCharacterClassEntityList(list []dto.CharacterClassDTO) []entity.CharacterClass {
	result := make([]entity.CharacterClass, len(list))
	for i, v := range list {
		result[i] = ToCharacterClassEntity(v)
	}
	return result
}

func ToCharacterClassDTOList(list []entity.CharacterClass) []dto.CharacterClassDTO {
	result := make([]dto.CharacterClassDTO, len(list))
	for i, v := range list {
		result[i] = ToCharacterClassDTO(v)
	}
	return result
}
func ToCharacterClassEntity(d dto.CharacterClassDTO) entity.CharacterClass {
	return entity.CharacterClass{
		ID:          d.ID,
		Name:        d.Name,
		Level:       d.Level,
		HitDice:     d.HitDice,
		Abilities:   ToAbilityEntityList(d.Abilities),
		ExtraFields: d.ExtraFields,
	}
}

func ToCharacterClassDTO(e entity.CharacterClass) dto.CharacterClassDTO {
	return dto.CharacterClassDTO{
		ID:          e.ID,
		Name:        e.Name,
		Level:       e.Level,
		HitDice:     e.HitDice,
		Abilities:   ToAbilityDTOList(e.Abilities),
		ExtraFields: e.ExtraFields,
	}
}

// CHARACTER RACE

func ToCharacterRaceEntity(d dto.CharacterRaceDTO) entity.CharacterRace {
	return entity.CharacterRace{
		ID:          d.ID,
		Name:        d.Name,
		Subrace:     d.Subrace,
		Traits:      d.Traits,
		Abilities:   ToAbilityEntityList(d.Abilities),
		ExtraFields: d.ExtraFields,
	}
}

func ToCharacterRaceDTO(e entity.CharacterRace) dto.CharacterRaceDTO {
	return dto.CharacterRaceDTO{
		ID:          e.ID,
		Name:        e.Name,
		Subrace:     e.Subrace,
		Traits:      e.Traits,
		Abilities:   ToAbilityDTOList(e.Abilities),
		ExtraFields: e.ExtraFields,
	}
}

// BACKGROUND
func ToBackgroundEntity(d dto.BackgroundDTO) entity.Background {
	return entity.Background{
		ID:            d.ID,
		Name:          d.Name,
		Abilities:     ToAbilityEntityList(d.Abilities),
		Proficiencies: d.Proficiencies,
		ExtraFields:   d.ExtraFields,
	}
}

func ToBackgroundDTO(e entity.Background) dto.BackgroundDTO {
	return dto.BackgroundDTO{
		ID:            e.ID,
		Name:          e.Name,
		Abilities:     ToAbilityDTOList(e.Abilities),
		Proficiencies: e.Proficiencies,
		ExtraFields:   e.ExtraFields,
	}
}

// ABILITY

func ToAbilityEntity(d dto.AbilityDTO) entity.Ability {
	return entity.Ability{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		LevelGained: d.LevelGained,
	}
}

func ToAbilityDTO(e entity.Ability) dto.AbilityDTO {
	return dto.AbilityDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		LevelGained: e.LevelGained,
	}
}

func ToAbilityEntityList(list []dto.AbilityDTO) []entity.Ability {
	result := make([]entity.Ability, len(list))
	for i, v := range list {
		result[i] = ToAbilityEntity(v)
	}
	return result
}

func ToAbilityDTOList(list []entity.Ability) []dto.AbilityDTO {
	result := make([]dto.AbilityDTO, len(list))
	for i, v := range list {
		result[i] = ToAbilityDTO(v)
	}
	return result
}

// SPELL

func ToSpellEntity(d dto.SpellDTO) entity.Spell {
	return entity.Spell{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Level:       d.Level,
	}
}

func ToSpellDTO(e entity.Spell) dto.SpellDTO {
	return dto.SpellDTO{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Level:       e.Level,
	}
}

func ToSpellEntityList(list []dto.SpellDTO) []entity.Spell {
	result := make([]entity.Spell, len(list))
	for i, v := range list {
		result[i] = ToSpellEntity(v)
	}
	return result
}

func ToSpellDTOList(list []entity.Spell) []dto.SpellDTO {
	result := make([]dto.SpellDTO, len(list))
	for i, v := range list {
		result[i] = ToSpellDTO(v)
	}
	return result
}

// ATTRIBUTES

func ToAttributesEntity(d dto.AttributesDTO) entity.Attributes {
	return entity.Attributes{
		Strength:     d.Strength,
		Dexterity:    d.Dexterity,
		Constitution: d.Constitution,
		Intelligence: d.Intelligence,
		Wisdom:       d.Wisdom,
		Charisma:     d.Charisma,
		Skills:       ToSkillsEntity(d.SkillsDTO),
	}
}

func ToAttributesDTO(e entity.Attributes) dto.AttributesDTO {
	return dto.AttributesDTO{
		Strength:     e.Strength,
		Dexterity:    e.Dexterity,
		Constitution: e.Constitution,
		Intelligence: e.Intelligence,
		Wisdom:       e.Wisdom,
		Charisma:     e.Charisma,
		SkillsDTO:    ToSkillsDTO(e.Skills),
	}
}

// MAPPER

func ToCharacterStatusEntity(d dto.CharacterStatusDTO) entity.CharacterStatus {
	return entity.CharacterStatus(d)
}

func ToCharacterStatusDTO(e entity.CharacterStatus) dto.CharacterStatusDTO {
	return dto.CharacterStatusDTO(e)
}

// Skills
func ToSkillsEntity(d dto.SkillsDTO) entity.Skills {
	return entity.Skills{
		Acrobatics:     ToSkillEntity(d.Acrobatics),
		AnimalHandling: ToSkillEntity(d.AnimalHandling),
		Arcana:         ToSkillEntity(d.Arcana),
		Athletics:      ToSkillEntity(d.Athletics),
		Deception:      ToSkillEntity(d.Deception),
		History:        ToSkillEntity(d.History),
		Insight:        ToSkillEntity(d.Insight),
		Intimidation:   ToSkillEntity(d.Intimidation),
		Investigation:  ToSkillEntity(d.Investigation),
		Medicine:       ToSkillEntity(d.Medicine),
		Nature:         ToSkillEntity(d.Nature),
		Perception:     ToSkillEntity(d.Perception),
		Performance:    ToSkillEntity(d.Performance),
		Persuasion:     ToSkillEntity(d.Persuasion),
		Religion:       ToSkillEntity(d.Religion),
		SleightOfHand:  ToSkillEntity(d.SleightOfHand),
		Stealth:        ToSkillEntity(d.Stealth),
		Survival:       ToSkillEntity(d.Survival),
	}
}

func ToSkillsDTO(e entity.Skills) dto.SkillsDTO {
	return dto.SkillsDTO{
		Acrobatics:     ToSkillDTO(e.Acrobatics),
		AnimalHandling: ToSkillDTO(e.AnimalHandling),
		Arcana:         ToSkillDTO(e.Arcana),
		Athletics:      ToSkillDTO(e.Athletics),
		Deception:      ToSkillDTO(e.Deception),
		History:        ToSkillDTO(e.History),
		Insight:        ToSkillDTO(e.Insight),
		Intimidation:   ToSkillDTO(e.Intimidation),
		Investigation:  ToSkillDTO(e.Investigation),
		Medicine:       ToSkillDTO(e.Medicine),
		Nature:         ToSkillDTO(e.Nature),
		Perception:     ToSkillDTO(e.Perception),
		Performance:    ToSkillDTO(e.Performance),
		Persuasion:     ToSkillDTO(e.Persuasion),
		Religion:       ToSkillDTO(e.Religion),
		SleightOfHand:  ToSkillDTO(e.SleightOfHand),
		Stealth:        ToSkillDTO(e.Stealth),
		Survival:       ToSkillDTO(e.Survival),
	}
}

// SKILL
func ToSkillEntity(d dto.SkillDTO) entity.Skill {
	return entity.Skill(d)
}

func ToSkillDTO(e entity.Skill) dto.SkillDTO {
	return dto.SkillDTO(e)
}

// SKILL

func ToSkillEntityList(list []dto.SkillDTO) []entity.Skill {
	result := make([]entity.Skill, len(list))
	for i, v := range list {
		result[i] = ToSkillEntity(v)
	}
	return result
}

func ToSkillDTOList(list []entity.Skill) []dto.SkillDTO {
	result := make([]dto.SkillDTO, len(list))
	for i, v := range list {
		result[i] = ToSkillDTO(v)
	}
	return result
}

// INVENTORY

func ToInventoryEntity(d dto.InventoryDTO) entity.Inventory {
	return entity.Inventory{
		Items:       ToItemEntityList(d.Items),
		TotalWeight: d.TotalWeight,
		Currency:    ToCurrencyEntity(d.Currency),
	}
}

func ToInventoryDTO(e entity.Inventory) dto.InventoryDTO {
	return dto.InventoryDTO{
		Items:       ToItemDTOList(e.Items),
		TotalWeight: e.TotalWeight,
		Currency:    ToCurrencyDTO(e.Currency),
	}
}

// CURRENCY

func ToCurrencyEntity(d dto.CurrencyDTO) entity.Currency {
	return entity.Currency(d)
}

func ToCurrencyDTO(e entity.Currency) dto.CurrencyDTO {
	return dto.CurrencyDTO(e)
}
