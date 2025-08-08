from pydantic import BaseModel, Field
from typing import List, Dict

from character import (
    CharacterStatus,
    Attributes,
    Ability,
    Spell,
    Inventory,
    Background,
    CharacterRace,
    CharacterClass,
    SpellSlot,
    Skill
)


class PlayerChar(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    status: CharacterStatus = Field(default_factory=CharacterStatus)
    attributes: Attributes = Field(default_factory=Attributes)
    alignment: str = Field("")
    abilities: List[Ability] = Field(default_factory=list)
    spells: List[Spell] = Field(default_factory=list)
    tags: List[str] = Field(default_factory=list)

    race: CharacterRace = Field(default_factory=CharacterRace)
    background: Background = Field(default_factory=Background)
    inventory: Inventory = Field(default_factory=Inventory)
    classes: List[CharacterClass] = Field(default_factory=list)
    level: int = Field(0)
    experience: int = Field(0)
    spell_slots: Dict[int, SpellSlot] = Field(default_factory=dict)
    skills: List[Skill] = Field(default_factory=list)
