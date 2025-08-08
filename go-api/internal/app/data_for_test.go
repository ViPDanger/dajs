package app_test

import (
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

var users = [...]dto.UserDTO{
	{Username: "alice", Password: "password123"},
	{Username: "bob", Password: "qwerty456"},
	{Username: "charlie", Password: "letmein789"},
	{Username: "diana", Password: "passw0rd"},
	{Username: "edward", Password: "secure!456"},
	{Username: "fiona", Password: "helloWorld1"},
	{Username: "george", Password: "abcDEF123"},
	{Username: "hannah", Password: "sunshine99"},
	{Username: "ian", Password: "hunter2"},
	{Username: "julia", Password: "admin@123"},
}

var characters = [...]dto.CharacterDTO{
	{
		ID:        "",
		Name:      "Thalor",
		Alignment: "Chaotic Good",
		Status:    dto.CharacterStatusDTO{HP: 35, MaxHP: 35, TemporaryHP: 0, ArmorClass: 17, Speed: 30, Initiative: 3},
		Attributes: dto.AttributesDTO{
			Strength:     16,
			Dexterity:    12,
			Constitution: 14,
			Intelligence: 10,
			Wisdom:       13,
			Charisma:     15,
			SkillsDTO: dto.SkillsDTO{
				Athletics:    dto.SkillDTO{Proficient: true, Modifier: 5},
				Persuasion:   dto.SkillDTO{Proficient: true, Modifier: 4},
				Insight:      dto.SkillDTO{Proficient: false, Modifier: 1},
				Perception:   dto.SkillDTO{Proficient: true, Modifier: 3},
				History:      dto.SkillDTO{Proficient: false, Modifier: 0},
				Intimidation: dto.SkillDTO{Proficient: true, Modifier: 4},
			},
		},
		Abilities: []dto.AbilityDTO{
			{ID: "ab1", Name: "Divine Smite", Description: "Radiant burst damage", LevelGained: 2},
		},
		Spells: []dto.SpellDTO{
			{ID: "sp1", Name: "Bless", Description: "Add d4 to attacks/saves", Level: 1},
			{ID: "sp2", Name: "Cure Wounds", Description: "Restore HP", Level: 1},
		},
		Tags: []string{"tank", "support"},
	},
	{
		ID:        "",
		Name:      "Sylra Moonshade",
		Alignment: "Neutral Good",
		Status:    dto.CharacterStatusDTO{HP: 27, MaxHP: 27, TemporaryHP: 0, ArmorClass: 14, Speed: 35, Initiative: 5},
		Attributes: dto.AttributesDTO{
			Strength:     10,
			Dexterity:    18,
			Constitution: 12,
			Intelligence: 14,
			Wisdom:       15,
			Charisma:     11,
			SkillsDTO: dto.SkillsDTO{
				Stealth:        dto.SkillDTO{Proficient: true, Modifier: 6},
				Nature:         dto.SkillDTO{Proficient: true, Modifier: 4},
				Survival:       dto.SkillDTO{Proficient: true, Modifier: 5},
				Perception:     dto.SkillDTO{Proficient: true, Modifier: 5},
				AnimalHandling: dto.SkillDTO{Proficient: false, Modifier: 2},
			},
		},
		Abilities: []dto.AbilityDTO{
			{ID: "ab2", Name: "Favored Enemy", Description: "Bonus vs chosen creature type", LevelGained: 1},
		},
		Spells: []dto.SpellDTO{
			{ID: "sp3", Name: "Hunter's Mark", Description: "Mark enemy for extra damage", Level: 1},
		},
		Tags: []string{"scout", "ranged"},
	},
}
var playerChars = []dto.PlayerCharDTO{
	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "p1",
			Name:      "Aric the Brave",
			Alignment: "Lawful Good",
			Status: dto.CharacterStatusDTO{
				HP:          45,
				MaxHP:       45,
				TemporaryHP: 0,
				ArmorClass:  18,
				Speed:       30,
				Initiative:  2,
			},
			Attributes: dto.AttributesDTO{
				Strength:     16,
				Dexterity:    12,
				Constitution: 14,
				Intelligence: 10,
				Wisdom:       13,
				Charisma:     15,
				SkillsDTO: dto.SkillsDTO{
					Athletics: dto.SkillDTO{Proficient: true, Modifier: 5},
					Insight:   dto.SkillDTO{Proficient: false, Modifier: 1},
					History:   dto.SkillDTO{Proficient: true, Modifier: 2},
				},
			},
			Abilities: []dto.AbilityDTO{
				{ID: "smite", Name: "Divine Smite", Description: "Radiant melee damage", LevelGained: 2},
			},
			Spells: []dto.SpellDTO{
				{ID: "bless", Name: "Bless", Description: "Buff allies", Level: 1},
			},
			Tags: []string{"paladin", "tank"},
		},
		CharacterRaceDTO: dto.CharacterRaceDTO{
			ID:        "human",
			Name:      "Human",
			Subrace:   "",
			Traits:    []string{"Versatile"},
			Abilities: []dto.AbilityDTO{},
		},
		BackgroundDTO: dto.BackgroundDTO{
			ID:            "noble",
			Name:          "Noble",
			Abilities:     []dto.AbilityDTO{{ID: "retainers", Name: "Retainers", Description: "Three loyal retainers", LevelGained: 1}},
			Proficiencies: []string{"History", "Persuasion"},
		},
		InventoryDTO: dto.InventoryDTO{
			Items: []dto.ItemDTO{
				{
					Id:                "i1",
					Name:              "Longsword",
					OrigName:          "Longsword",
					Weight:            3.0,
					WeaponFormula:     "1d8",
					WeaponAttackBonus: 3,
					Tags:              []string{"weapon,melee"},
					WeaponDamageType:  1,
					PropertyWeapon:    "Versatile",
				},
				{
					Id:            "i2",
					Name:          "Chain Mail",
					OrigName:      "Chain Mail",
					Weight:        20.0,
					Tags:          []string{"armor,heavy"},
					ArmorInt:      16,
					StrArmor:      13,
					NoDexArmor:    true,
					PropertyArmor: "Disadvantage on Stealth",
				},
			},
			TotalWeight: 23.0,
			Currency:    dto.CurrencyDTO{Gold: 50, Silver: 10, Copper: 5},
		},
		Classes: []dto.CharacterClassDTO{
			{
				ID:        "paladin",
				Name:      "Paladin",
				Level:     5,
				HitDice:   "1d10",
				Abilities: []dto.AbilityDTO{{ID: "lay_on_hands", Name: "Lay on Hands", Description: "Healing pool", LevelGained: 1}},
			},
		},
		Level:      5,
		Experience: 6500,
		SpellSlots: map[int]dto.SpellSlotDTO{
			1: {Max: 4, Current: 3},
			2: {Max: 2, Current: 1},
		},
		Skills: []dto.SkillDTO{
			{Proficient: true, Modifier: 5},
			{Proficient: false, Modifier: 1},
		},
	},

	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "p2",
			Name:      "Lyra Moonshadow",
			Alignment: "Chaotic Good",
			Status: dto.CharacterStatusDTO{
				HP:          30,
				MaxHP:       30,
				TemporaryHP: 0,
				ArmorClass:  14,
				Speed:       35,
				Initiative:  4,
			},
			Attributes: dto.AttributesDTO{
				Strength:     8,
				Dexterity:    18,
				Constitution: 12,
				Intelligence: 14,
				Wisdom:       11,
				Charisma:     13,
				SkillsDTO: dto.SkillsDTO{
					Stealth:    dto.SkillDTO{Proficient: true, Modifier: 6},
					Acrobatics: dto.SkillDTO{Proficient: true, Modifier: 6},
					Arcana:     dto.SkillDTO{Proficient: false, Modifier: 2},
				},
			},
			Abilities: []dto.AbilityDTO{
				{ID: "sneak", Name: "Sneak Attack", Description: "Bonus damage", LevelGained: 1},
			},
			Spells: []dto.SpellDTO{},
			Tags:   []string{"rogue", "stealth"},
		},
		CharacterRaceDTO: dto.CharacterRaceDTO{
			ID:        "elf",
			Name:      "Elf",
			Subrace:   "Wood Elf",
			Traits:    []string{"Darkvision", "Fey Ancestry"},
			Abilities: []dto.AbilityDTO{},
		},
		BackgroundDTO: dto.BackgroundDTO{
			ID:            "urchin",
			Name:          "Urchin",
			Abilities:     []dto.AbilityDTO{{ID: "city_secrets", Name: "City Secrets", Description: "Navigate cities quickly", LevelGained: 1}},
			Proficiencies: []string{"Sleight of Hand", "Stealth"},
		},
		InventoryDTO: dto.InventoryDTO{
			Items: []dto.ItemDTO{
				{
					Id:       "i3",
					Name:     "Dagger",
					OrigName: "Dagger",
					Weight:   1.0,

					Tags:              []string{"weapon,light,finesse"},
					WeaponFormula:     "1d4",
					WeaponAttackBonus: 4,
					WeaponDamageType:  1,
				},
				{
					Id:       "i4",
					Name:     "Thieves' Tools",
					OrigName: "Thieves' Tools",
					Weight:   0.5,
					Tags:     []string{"tool,utility"},
				},
			},
			TotalWeight: 1.5,
			Currency:    dto.CurrencyDTO{Gold: 12, Silver: 5, Copper: 30},
		},
		Classes: []dto.CharacterClassDTO{
			{
				ID:        "rogue",
				Name:      "Rogue",
				Level:     4,
				HitDice:   "1d8",
				Abilities: []dto.AbilityDTO{{ID: "cunning", Name: "Cunning Action", Description: "Bonus action utility", LevelGained: 2}},
			},
		},
		Level:      4,
		Experience: 3200,
		SpellSlots: map[int]dto.SpellSlotDTO{},
		Skills: []dto.SkillDTO{
			{Proficient: true, Modifier: 6},  // Stealth
			{Proficient: false, Modifier: 2}, // Arcana
		},
	},

	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "p3",
			Name:      "Thorn Earthspeaker",
			Alignment: "Neutral",
			Status: dto.CharacterStatusDTO{
				HP:          38,
				MaxHP:       38,
				TemporaryHP: 0,
				ArmorClass:  16,
				Speed:       30,
				Initiative:  1,
			},
			Attributes: dto.AttributesDTO{
				Strength:     12,
				Dexterity:    10,
				Constitution: 14,
				Intelligence: 11,
				Wisdom:       17,
				Charisma:     9,
				SkillsDTO: dto.SkillsDTO{
					Nature:   dto.SkillDTO{Proficient: true, Modifier: 3},
					Medicine: dto.SkillDTO{Proficient: true, Modifier: 5},
					Survival: dto.SkillDTO{Proficient: false, Modifier: 3},
				},
			},
			Abilities: []dto.AbilityDTO{
				{ID: "wild_shape", Name: "Wild Shape", Description: "Transform into beast", LevelGained: 2},
			},
			Spells: []dto.SpellDTO{
				{ID: "entangle", Name: "Entangle", Description: "Restrains enemies", Level: 1},
				{ID: "healing_word", Name: "Healing Word", Description: "Heals at range", Level: 1},
			},
			Tags: []string{"druid", "nature"},
		},
		CharacterRaceDTO: dto.CharacterRaceDTO{
			ID:        "dwarf",
			Name:      "Dwarf",
			Subrace:   "Hill Dwarf",
			Traits:    []string{"Darkvision", "Dwarven Resilience"},
			Abilities: []dto.AbilityDTO{},
		},
		BackgroundDTO: dto.BackgroundDTO{
			ID:            "hermit",
			Name:          "Hermit",
			Abilities:     []dto.AbilityDTO{{ID: "discovery", Name: "Discovery", Description: "Major revelation", LevelGained: 1}},
			Proficiencies: []string{"Medicine", "Religion"},
		},
		InventoryDTO: dto.InventoryDTO{
			Items: []dto.ItemDTO{
				{
					Id:                "i5",
					Name:              "Quarterstaff",
					OrigName:          "Quarterstaff",
					Weight:            4.0,
					Tags:              []string{"weapon,versatile"},
					WeaponFormula:     "1d8",
					WeaponAttackBonus: 2,
					WeaponDamageType:  1,
				},
				{
					Id:       "i6",
					Name:     "Herbalism Kit",
					OrigName: "Herbalism Kit",
					Weight:   2.0,
					Tags:     []string{"tool"},
				},
			},
			TotalWeight: 6.0,
			Currency:    dto.CurrencyDTO{Gold: 10, Silver: 2, Copper: 40},
		},
		Classes: []dto.CharacterClassDTO{
			{
				ID:        "druid",
				Name:      "Druid",
				Level:     5,
				HitDice:   "1d8",
				Abilities: []dto.AbilityDTO{{ID: "wild_shape", Name: "Wild Shape", Description: "Transform into beast", LevelGained: 2}},
			},
		},
		Level:      5,
		Experience: 7000,
		SpellSlots: map[int]dto.SpellSlotDTO{
			1: {Max: 4, Current: 3},
			2: {Max: 2, Current: 2},
		},
		Skills: []dto.SkillDTO{
			{Proficient: true, Modifier: 5},  // Medicine
			{Proficient: false, Modifier: 3}, // Survival
		},
	},
}
var items = []dto.ItemDTO{
	{
		Id:                "w1",
		Name:              "Longsword",
		OrigName:          "Longsword",
		Comment:           "Versatile martial melee weapon",
		Price:             15.0,
		Money:             1500,
		Weight:            3.0,
		HtmlText:          "<b>1d8 slashing</b>",
		Tags:              []string{"weapon", "melee", "versatile"},
		WeaponFormula:     "1d8",
		WeaponAttackBonus: 2,
		WeaponDamageType:  1,
		IsFencing:         false,
		HasDescription:    true,
	},
	{
		Id:             "a1",
		Name:           "Chain Mail",
		OrigName:       "Chain Mail",
		Comment:        "Heavy armor, requires STR 13",
		Price:          75.0,
		Money:          7500,
		Weight:         20.0,
		HtmlText:       "<b>AC 16</b>",
		Tags:           []string{"armor", "heavy"},
		ArmorInt:       16,
		StrArmor:       13,
		StealthDis:     true,
		NoDexArmor:     true,
		PropertyArmor:  "Disadvantage on Stealth",
		HasDescription: true,
	},
	{
		Id:             "am1",
		Name:           "Arrows (20)",
		OrigName:       "Arrows",
		Comment:        "Standard ammunition for bows",
		Price:          1.0,
		Money:          100,
		Weight:         1.0,
		Tags:           []string{"ammo", "arrow"},
		IsCounting:     true,
		DefaultCount:   20,
		CustomTags:     "piercing",
		HasDescription: false,
	},
	{
		Id:             "c1",
		Name:           "Backpack",
		OrigName:       "Backpack",
		Comment:        "Useful for carrying gear",
		Price:          2.0,
		Money:          200,
		Weight:         5.0,
		Tags:           []string{"container", "gear"},
		HasDescription: true,
		List: []dto.ItemDTO{
			{
				Id:                "w2",
				Name:              "Dagger",
				OrigName:          "Dagger",
				Comment:           "Small, light melee weapon",
				Price:             2.0,
				Money:             200,
				Weight:            1.0,
				Tags:              []string{"weapon", "light", "finesse"},
				WeaponFormula:     "1d4",
				WeaponAttackBonus: 1,
				WeaponDamageType:  1,
				IsFencing:         true,
			},
		},
	},
	{
		Id:             "a2",
		Name:           "Leather Armor",
		OrigName:       "Leather Armor",
		Comment:        "Basic light armor",
		Price:          10.0,
		Money:          1000,
		Weight:         11.0,
		Tags:           []string{"armor", "light"},
		ArmorInt:       11,
		ShortDexArmor:  true,
		PropertyArmor:  "Allows full Dex bonus",
		HasDescription: true,
	},
}
var npcs = []dto.NPCdto{
	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "npc001",
			Name:      "Thargrim Stoneforge",
			Alignment: "Lawful Good",
			Status: dto.CharacterStatusDTO{
				HP:          20,
				MaxHP:       20,
				TemporaryHP: 0,
				ArmorClass:  16,
				Speed:       25,
				Initiative:  0,
			},
			Attributes: dto.AttributesDTO{
				Strength:     15,
				Dexterity:    10,
				Constitution: 14,
				Intelligence: 11,
				Wisdom:       12,
				Charisma:     8,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "smith1", Name: "Forge Mastery", Description: "Expert in crafting armor and weapons", LevelGained: 1},
			},
			Spells: nil,
			Tags:   []string{"dwarf", "blacksmith"},
		},
		Occupation: "Blacksmith",
		Notes:      "Owns the village forge. Knows local gossip.",
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "npc002",
			Name:      "Elandra the White",
			Alignment: "Neutral Good",
			Status: dto.CharacterStatusDTO{
				HP:          18,
				MaxHP:       18,
				TemporaryHP: 0,
				ArmorClass:  13,
				Speed:       30,
				Initiative:  1,
			},
			Attributes: dto.AttributesDTO{
				Strength:     8,
				Dexterity:    12,
				Constitution: 10,
				Intelligence: 14,
				Wisdom:       16,
				Charisma:     13,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "healing1", Name: "Prayer of Healing", Description: "Heals allies", LevelGained: 1},
			},
			Spells: []dto.SpellDTO{
				{ID: "heal", Name: "Cure Wounds", Description: "Restores hit points", Level: 1},
			},
			Tags: []string{"elf", "priestess", "healer"},
		},
		Occupation: "Priestess",
		Notes:      "Caretaker of the temple. Soft-spoken and wise.",
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "npc003",
			Name:      "Bram Copperfoot",
			Alignment: "Chaotic Neutral",
			Status: dto.CharacterStatusDTO{
				HP:          12,
				MaxHP:       12,
				TemporaryHP: 0,
				ArmorClass:  12,
				Speed:       25,
				Initiative:  2,
			},
			Attributes: dto.AttributesDTO{
				Strength:     10,
				Dexterity:    14,
				Constitution: 10,
				Intelligence: 13,
				Wisdom:       9,
				Charisma:     14,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "trick1", Name: "Street Smarts", Description: "Advantage on deception in cities", LevelGained: 1},
			},
			Spells: nil,
			Tags:   []string{"halfling", "merchant"},
		},
		Occupation: "Merchant",
		Notes:      "Smooth talker. Deals in rare curiosities.",
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:        "npc004",
			Name:      "Ser Reginald Crowne",
			Alignment: "Lawful Neutral",
			Status: dto.CharacterStatusDTO{
				HP:          22,
				MaxHP:       22,
				TemporaryHP: 0,
				ArmorClass:  17,
				Speed:       30,
				Initiative:  1,
			},
			Attributes: dto.AttributesDTO{
				Strength:     16,
				Dexterity:    12,
				Constitution: 14,
				Intelligence: 11,
				Wisdom:       10,
				Charisma:     13,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "guard1", Name: "Shield Wall", Description: "Can guard adjacent allies", LevelGained: 2},
			},
			Spells: nil,
			Tags:   []string{"human", "guard", "knight"},
		},
		Occupation: "City Guard Captain",
		Notes:      "Disciplined veteran of the local militia.",
	},
}
var monsters = []dto.MonsterDTO{
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m1",
			Name: "Goblin",
			Status: dto.CharacterStatusDTO{
				HP:          7,
				MaxHP:       7,
				TemporaryHP: 0,
				ArmorClass:  15,
				Speed:       30,
				Initiative:  2,
			},
			Attributes: dto.AttributesDTO{
				Strength:     8,
				Dexterity:    14,
				Constitution: 10,
				Intelligence: 10,
				Wisdom:       8,
				Charisma:     8,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "nimble_escape", Name: "Nimble Escape", Description: "Can take the Disengage or Hide action as a bonus action", LevelGained: 1},
			},
			Spells: nil,
			Tags:   []string{"goblin", "humanoid"},
		},
		Type:      "humanoid",
		Challenge: 0.25,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Scimitar", Description: "Melee Weapon Attack", Damage: "1d6+2"},
			{Name: "Shortbow", Description: "Ranged Weapon Attack", Damage: "1d6+2"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m2",
			Name: "Orc",
			Status: dto.CharacterStatusDTO{
				HP:          15,
				MaxHP:       15,
				TemporaryHP: 0,
				ArmorClass:  13,
				Speed:       30,
				Initiative:  1,
			},
			Attributes: dto.AttributesDTO{
				Strength:     16,
				Dexterity:    12,
				Constitution: 16,
				Intelligence: 7,
				Wisdom:       11,
				Charisma:     10,
			},
			Abilities: []dto.AbilityDTO{
				{ID: "aggressive", Name: "Aggressive", Description: "Can move as bonus action toward enemy", LevelGained: 1},
			},
			Tags: []string{"orc", "brute"},
		},
		Type:      "humanoid",
		Challenge: 0.5,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Greataxe", Description: "Melee Weapon Attack", Damage: "1d12+3"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m3",
			Name: "Young Red Dragon",
			Status: dto.CharacterStatusDTO{
				HP:          178,
				MaxHP:       178,
				TemporaryHP: 0,
				ArmorClass:  18,
				Speed:       40,
				Initiative:  2,
			},
			Attributes: dto.AttributesDTO{
				Strength:     23,
				Dexterity:    10,
				Constitution: 21,
				Intelligence: 14,
				Wisdom:       11,
				Charisma:     19,
			},
			Tags: []string{"dragon", "fire"},
		},
		Type:      "dragon",
		Challenge: 10,
		Legendary: true,
		Actions: []dto.ActionDTO{
			{Name: "Multiattack", Description: "1 bite and 2 claw attacks"},
			{Name: "Fire Breath", Description: "Exhales fire in 30ft cone", Damage: "12d6"},
		},
		LegendaryActions: []dto.ActionDTO{
			{Name: "Wing Attack", Description: "Beats wings to knock enemies prone", Damage: "2d6+6"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m4",
			Name: "Skeleton",
			Status: dto.CharacterStatusDTO{
				HP:         13,
				MaxHP:      13,
				ArmorClass: 13,
				Speed:      30,
				Initiative: 0,
			},
			Attributes: dto.AttributesDTO{
				Strength:     10,
				Dexterity:    14,
				Constitution: 15,
				Intelligence: 6,
				Wisdom:       8,
				Charisma:     5,
			},
			Tags: []string{"undead", "skeleton"},
		},
		Type:      "undead",
		Challenge: 0.25,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Shortsword", Description: "Melee Weapon Attack", Damage: "1d6+2"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m5",
			Name: "Dire Wolf",
			Status: dto.CharacterStatusDTO{
				HP:         37,
				MaxHP:      37,
				ArmorClass: 14,
				Speed:      50,
				Initiative: 2,
			},
			Attributes: dto.AttributesDTO{
				Strength:     17,
				Dexterity:    15,
				Constitution: 15,
				Intelligence: 3,
				Wisdom:       12,
				Charisma:     7,
			},
			Tags: []string{"beast", "wolf"},
		},
		Type:      "beast",
		Challenge: 1.0,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Bite", Description: "Melee Weapon Attack", Damage: "2d6+3"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m6",
			Name: "Zombie",
			Status: dto.CharacterStatusDTO{
				HP:         22,
				MaxHP:      22,
				ArmorClass: 8,
				Speed:      20,
			},
			Attributes: dto.AttributesDTO{
				Strength:     13,
				Dexterity:    6,
				Constitution: 16,
				Intelligence: 3,
				Wisdom:       6,
				Charisma:     5,
			},
			Tags: []string{"undead", "zombie"},
		},
		Type:      "undead",
		Challenge: 0.25,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Slam", Description: "Melee Weapon Attack", Damage: "1d6+1"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m7",
			Name: "Troll",
			Status: dto.CharacterStatusDTO{
				HP:         84,
				MaxHP:      84,
				ArmorClass: 15,
				Speed:      30,
			},
			Attributes: dto.AttributesDTO{
				Strength:     18,
				Dexterity:    13,
				Constitution: 20,
				Intelligence: 7,
				Wisdom:       9,
				Charisma:     7,
			},
			Tags: []string{"giant", "regeneration"},
		},
		Type:      "giant",
		Challenge: 5.0,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Multiattack", Description: "Three attacks: one bite and two claws"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m8",
			Name: "Banshee",
			Status: dto.CharacterStatusDTO{
				HP:         58,
				MaxHP:      58,
				ArmorClass: 12,
				Speed:      0,
			},
			Attributes: dto.AttributesDTO{
				Strength:     1,
				Dexterity:    14,
				Constitution: 10,
				Intelligence: 12,
				Wisdom:       11,
				Charisma:     17,
			},
			Tags: []string{"undead", "ghost"},
		},
		Type:      "undead",
		Challenge: 4.0,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Wail", Description: "All creatures within 30ft make CHA save or drop to 0 HP"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m9",
			Name: "Beholder",
			Status: dto.CharacterStatusDTO{
				HP:         180,
				MaxHP:      180,
				ArmorClass: 18,
				Speed:      0,
			},
			Attributes: dto.AttributesDTO{
				Strength:     10,
				Dexterity:    14,
				Constitution: 18,
				Intelligence: 17,
				Wisdom:       15,
				Charisma:     17,
			},
			Tags: []string{"aberration", "eye"},
		},
		Type:      "aberration",
		Challenge: 13.0,
		Legendary: true,
		Actions: []dto.ActionDTO{
			{Name: "Eye Rays", Description: "Shoots 3 eye rays per turn"},
		},
		LegendaryActions: []dto.ActionDTO{
			{Name: "Legendary Ray", Description: "Shoots additional ray at end of another's turn"},
		},
	},
	{
		CharacterDTO: dto.CharacterDTO{
			ID:   "m10",
			Name: "Gelatinous Cube",
			Status: dto.CharacterStatusDTO{
				HP:         84,
				MaxHP:      84,
				ArmorClass: 6,
				Speed:      15,
			},
			Attributes: dto.AttributesDTO{
				Strength:     14,
				Dexterity:    3,
				Constitution: 20,
				Intelligence: 1,
				Wisdom:       6,
				Charisma:     1,
			},
			Tags: []string{"ooze", "cube"},
		},
		Type:      "ooze",
		Challenge: 2.0,
		Legendary: false,
		Actions: []dto.ActionDTO{
			{Name: "Pseudopod", Description: "Melee Weapon Attack", Damage: "3d6+2"},
		},
	},
}
var glossarys = []dto.GlossaryDTO{
	{
		ID: "g001",
		Text: `**Armor Class (AC)** represents how difficult it is to hit a creature with an attack. It is calculated based on the creature’s armor, Dexterity modifier, and other bonuses.

Higher AC means better defense. A creature with AC 18 is much harder to hit than one with AC 12.`,
	},
	{
		ID: "g002",
		Text: `**Hit Points (HP)** measure a creature’s ability to survive damage. When reduced to 0 HP, a creature falls unconscious or dies.

Temporary HP can absorb damage but do not stack and cannot heal you.`,
	},
	{
		ID: "g003",
		Text: `**Initiative** determines the order of actions during combat. Each participant rolls a d20 and adds their Dexterity modifier.

The DM organizes all combatants into a turn order based on these results.`,
	},
	{
		ID: "g004",
		Text: `**Spell Slots** are the resource a spellcaster uses to cast spells. Each spell has a level, and casting it consumes a slot of equal or higher level.

Spell slots are regained after long rests (or short rests for some classes like Warlocks).`,
	},
	{
		ID: "g005",
		Text: `**Saving Throws** are special checks made to resist harmful effects such as poison, magic, or traps. They are tied to your abilities (e.g., Constitution save).

Some classes have proficiency in certain saves, giving them a bonus.`,
	},
	{
		ID: "g006",
		Text: `**Conditions** are status effects that alter a creature’s capabilities. Examples include Blinded, Charmed, Frightened, Paralyzed, and Stunned.

Each condition has specific rules that affect how a creature behaves or what it can do.`,
	},
	{
		ID: "g007",
		Text: `**Advantage and Disadvantage**: When you have advantage, you roll two d20s and use the higher. Disadvantage means you roll two and take the lower.

They cancel each other out — if you have both, roll normally.`,
	},
	{
		ID: "g008",
		Text: `**Ability Checks** represent efforts to accomplish tasks, such as climbing, persuading, or recalling lore. They involve rolling a d20 and adding a relevant modifier.

The DM sets a Difficulty Class (DC), and you succeed if your result equals or exceeds it.`,
	},
	{
		ID: "g009",
		Text: `**Attunement** is required to use some powerful magic items. A character may only be attuned to three items at a time.

Attunement typically takes a short rest and must meet any prerequisites set by the item.`,
	},
	{
		ID: "g010",
		Text: `**Exhaustion** represents a cumulative penalty for overexertion or exposure to extreme conditions. It has six levels, ranging from disadvantage on ability checks to death.

Rest, magic, or specific features are required to remove levels of exhaustion.`,
	},
}
