from dto.player import PlayerCharacterDTO
from entity.player_char import PlayerChar


class PlayerCharacterMapper:

    @staticmethod
    def to_dto(entity: PlayerChar) -> PlayerCharacterDTO:
        return PlayerCharacterDTO(
            id=entity.id,
            name=entity.name,
            status=entity.status,
            attributes=entity.attributes,
            alignment=entity.alignment,
            abilities=entity.abilities,
            spells=entity.spells,
            tags=entity.tags,
            race=entity.race,
            background=entity.background,
            inventory=entity.inventory,
            classes=entity.classes,
            level=entity.level,
            experience=entity.experience,
            spell_slots=entity.spell_slots,
            skills=entity.skills,
        )

    @staticmethod
    def from_dto(dto: PlayerCharacterDTO) -> PlayerChar:
        return PlayerChar(
            id=dto.id,
            name=dto.name,
            status=dto.status,
            attributes=dto.attributes,
            alignment=dto.alignment,
            abilities=dto.abilities,
            spells=dto.spells,
            tags=dto.tags,
            race=dto.race,
            background=dto.background,
            inventory=dto.inventory,
            classes=dto.classes,
            level=dto.level,
            experience=dto.experience,
            spell_slots=dto.spell_slots,
            skills=dto.skills,
        )
