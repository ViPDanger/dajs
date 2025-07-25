from pydantic import BaseModel, Field
from typing import List, Dict, Any


class Skill(BaseModel):
    proficient: bool = Field(False, alias="proficient")
    modifier: int = Field(0, alias="modifier")


class Skills(BaseModel):
    acrobatics: Skill = Field(default_factory=Skill, alias="acrobatics")
    animal_handling: Skill = Field(default_factory=Skill, alias="animal_handling")
    arcana: Skill = Field(default_factory=Skill, alias="arcana")
    athletics: Skill = Field(default_factory=Skill, alias="athletics")
    deception: Skill = Field(default_factory=Skill, alias="deception")
    history: Skill = Field(default_factory=Skill, alias="history")
    insight: Skill = Field(default_factory=Skill, alias="insight")
    intimidation: Skill = Field(default_factory=Skill, alias="intimidation")
    investigation: Skill = Field(default_factory=Skill, alias="investigation")
    medicine: Skill = Field(default_factory=Skill, alias="medicine")
    nature: Skill = Field(default_factory=Skill, alias="nature")
    perception: Skill = Field(default_factory=Skill, alias="perception")
    performance: Skill = Field(default_factory=Skill, alias="performance")
    persuasion: Skill = Field(default_factory=Skill, alias="persuasion")
    religion: Skill = Field(default_factory=Skill, alias="religion")
    sleight_of_hand: Skill = Field(default_factory=Skill, alias="sleight_of_hand")
    stealth: Skill = Field(default_factory=Skill, alias="stealth")
    survival: Skill = Field(default_factory=Skill, alias="survival")


class Attributes(BaseModel):
    strength: int = Field(0, alias="strength")
    dexterity: int = Field(0, alias="dexterity")
    constitution: int = Field(0, alias="constitution")
    intelligence: int = Field(0, alias="intelligence")
    wisdom: int = Field(0, alias="wisdom")
    charisma: int = Field(0, alias="charisma")
    skills: Skills = Field(default_factory=Skills, alias="skills")


class Ability(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    description: str = Field("", alias="description")
    level_gained: int = Field(0, alias="level_gained")


class Spell(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    description: str = Field("", alias="description")
    level: int = Field(0, alias="level")


class SpellSlot(BaseModel):
    max: int = Field(0, alias="max")
    current: int = Field(0, alias="current")


class Currency(BaseModel):
    copper: int = Field(0, alias="copper")
    silver: int = Field(0, alias="silver")
    gold: int = Field(0, alias="gold")


class Item(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    description: str = Field("", alias="description")
    weight: float = Field(0.0, alias="weight")
    tags: List[str] = Field(default_factory=list, alias="tags")


class Inventory(BaseModel):
    items: List[Item] = Field(default_factory=list, alias="items")
    total_weight: float = Field(0.0, alias="total_weight")
    currency: Currency = Field(default_factory=Currency, alias="currency")


class CharacterStatus(BaseModel):
    hp: int = Field(0, alias="hp")
    max_hp: int = Field(0, alias="max_hp")
    temporary_hp: int = Field(0, alias="temporary_hp")
    armor_class: int = Field(0, alias="armor_class")
    speed: int = Field(0, alias="speed")
    initiative: int = Field(0, alias="initiative")


class Background(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    abilities: List[Ability] = Field(default_factory=list, alias="abilities")
    proficiencies: List[str] = Field(default_factory=list, alias="proficiencies")
    extra_fields: Dict[str, Any] = Field(default_factory=dict, alias="extra_fields")


class CharacterClass(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    level: int = Field(0, alias="level")
    hit_dice: str = Field("", alias="hit_dice")
    abilities: List[Ability] = Field(default_factory=list, alias="abilities")
    extra_fields: Dict[str, Any] = Field(default_factory=dict, alias="extra_fields")


class CharacterRace(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    subrace: str = Field("", alias="subrace")
    traits: List[str] = Field(default_factory=list, alias="traits")
    abilities: List[Ability] = Field(default_factory=list, alias="abilities")
    extra_fields: Dict[str, Any] = Field(default_factory=dict, alias="extra_fields")


class Character(BaseModel):
    id: str = Field("", alias="id")
    name: str = Field("", alias="name")
    status: CharacterStatus = Field(default_factory=CharacterStatus, alias="status")
    attributes: Attributes = Field(default_factory=Attributes, alias="attributes")
    alignment: str = Field("", alias="alignment")
    abilities: List[Ability] = Field(default_factory=list, alias="abilities")
    spells: List[Spell] = Field(default_factory=list, alias="spells")
    tags: List[str] = Field(default_factory=list, alias="tags")
