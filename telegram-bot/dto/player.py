from pydantic import BaseModel, Field
from typing import List, Dict

from character import (
    Character,
    CharacterRace,
    Background,
    Inventory,
    CharacterClass,
    SpellSlot,
    Skill,
)


class PlayerCharacterDTO(Character):
    race: CharacterRace = Field(default_factory=CharacterRace, alias="race")
    background: Background = Field(default_factory=Background, alias="background")
    inventory: Inventory = Field(default_factory=Inventory, alias="inventory")
    classes: List[CharacterClass] = Field(default_factory=list, alias="classes")
    level: int = Field(0, alias="level")
    experience: int = Field(0, alias="experience")
    spell_slots: Dict[int, SpellSlot] = Field(default_factory=dict, alias="spell_slots")
    skills: List[Skill] = Field(default_factory=list, alias="skills")