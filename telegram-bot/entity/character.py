from pydantic import BaseModel, Field
from typing import List, Dict, Any


# Skills
class Skill(BaseModel):
    proficient: bool = Field(False)
    modifier: int = Field(0)


class Skills(BaseModel):
    acrobatics: Skill = Field(default_factory=Skill)
    animal_handling: Skill = Field(default_factory=Skill)
    arcana: Skill = Field(default_factory=Skill)
    athletics: Skill = Field(default_factory=Skill)
    deception: Skill = Field(default_factory=Skill)
    history: Skill = Field(default_factory=Skill)
    insight: Skill = Field(default_factory=Skill)
    intimidation: Skill = Field(default_factory=Skill)
    investigation: Skill = Field(default_factory=Skill)
    medicine: Skill = Field(default_factory=Skill)
    nature: Skill = Field(default_factory=Skill)
    perception: Skill = Field(default_factory=Skill)
    performance: Skill = Field(default_factory=Skill)
    persuasion: Skill = Field(default_factory=Skill)
    religion: Skill = Field(default_factory=Skill)
    sleight_of_hand: Skill = Field(default_factory=Skill)
    stealth: Skill = Field(default_factory=Skill)
    survival: Skill = Field(default_factory=Skill)


# Attributes
class Attributes(BaseModel):
    strength: int = Field(0)
    dexterity: int = Field(0)
    constitution: int = Field(0)
    intelligence: int = Field(0)
    wisdom: int = Field(0)
    charisma: int = Field(0)
    skills: Skills = Field(default_factory=Skills)


# Status
class CharacterStatus(BaseModel):
    hp: int = Field(0)
    max_hp: int = Field(0)
    temporary_hp: int = Field(0)
    armor_class: int = Field(0)
    speed: int = Field(0)
    initiative: int = Field(0)


# Ability & Spell
class Ability(BaseModel):
    id: str = Field("")
    name: str = Field("")
    description: str = Field("")
    level_gained: int = Field(0)


class Spell(BaseModel):
    id: str = Field("")
    name: str = Field("")
    description: str = Field("")
    level: int = Field(0)


# Currency & Inventory
class Currency(BaseModel):
    copper: int = Field(0)
    silver: int = Field(0)
    gold: int = Field(0)


class Item(BaseModel):
    id: str = Field("")
    name: str = Field("")
    description: str = Field("")
    weight: float = Field(0.0)
    tags: List[str] = Field(default_factory=list)


class Inventory(BaseModel):
    items: List[Item] = Field(default_factory=list)
    total_weight: float = Field(0.0)
    currency: Currency = Field(default_factory=Currency)


# Race & Background
class CharacterRace(BaseModel):
    id: str = Field("")
    name: str = Field("")
    subrace: str = Field("")
    traits: List[str] = Field(default_factory=list)
    abilities: List[Ability] = Field(default_factory=list)
    extra_fields: Dict[str, Any] = Field(default_factory=dict)


class Background(BaseModel):
    id: str = Field("")
    name: str = Field("")
    abilities: List[Ability] = Field(default_factory=list)
    proficiencies: List[str] = Field(default_factory=list)
    extra_fields: Dict[str, Any] = Field(default_factory=dict)


# Class & SpellSlots
class CharacterClass(BaseModel):
    id: str = Field("")
    name: str = Field("")
    level: int = Field(0)
    hit_dice: str = Field("")
    abilities: List[Ability] = Field(default_factory=list)
    extra_fields: Dict[str, Any] = Field(default_factory=dict)


class SpellSlot(BaseModel):
    max: int = Field(0)
    current: int = Field(0)


# === Character Entity ===
class Character(BaseModel):
    id: str = Field("")
    name: str = Field("")
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
