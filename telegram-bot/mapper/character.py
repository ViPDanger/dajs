from character import Character as CharacterDTO
from entity.character import Character as CharacterEntity


class CharacterMapper:

    @staticmethod
    def to_dto(entity: CharacterEntity) -> CharacterDTO:
        return CharacterDTO(
            id=entity.id,
            name=entity.name,
            status=entity.status,
            attributes=entity.attributes,
            alignment=entity.alignment,
            abilities=entity.abilities,
            spells=entity.spells,
            tags=entity.tags,
        )

    @staticmethod
    def from_dto(dto: CharacterDTO) -> CharacterEntity:
        return CharacterEntity(
            id=dto.id,
            name=dto.name,
            status=dto.status,
            attributes=dto.attributes,
            alignment=dto.alignment,
            abilities=dto.abilities,
            spells=dto.spells,
            tags=dto.tags,
        )
