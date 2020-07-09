package rulebook

import "github.com/pandorasNox/go-savage-worlds/pkg/dice"

// SwadeRaces predefined for the SWADE core ruleset
var SwadeRaces = Races{
	{
		name:        "Android",
		description: "pacifistic...",
		abilities: racialAbilities{
			{
				name:           "Pacifist (Major)",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return addHindranceModBuilder("Pacifist", Major, SwadeHindrances, ca)
					},
				},
			},
			{
				name:           "Construct",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					plusShakenRecoveryAdjustmentMod,
					plusShakenRecoveryAdjustmentMod,
				},
			},
		},
	},
	{
		name:        "Aquarius",
		description: "",
		abilities: racialAbilities{
			{
				name:           "Toughness",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					plusBaseToughnessMod,
					plusBaseToughnessMod,
				},
			},
		},
	},
	{
		name:        "Avion",
		description: "has wings",
		abilities: racialAbilities{
			{
				name:           "Frail",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					minusBaseToughnessMod,
				},
			},
			//KEEN SENSES: Avions are more perceptive than most.
			//They begin with a d6 in Notice (instead of d4) and may raise the skill to d12 + 1.
			{
				name:           "Keen Senses",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return skillStartsAtModBuilder("Notice", dice.D6, ca)
					},
				},
			},
		},
	},
	{
		name:        "Dwarfs",
		description: "small and strong",
		abilities: racialAbilities{
			{
				name:           "Tough",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Vigor", dice.D6, ca)
					},
				},
			},
		},
	},
	{
		name:        "Elves",
		description: "Tall, long ears.",
		abilities: racialAbilities{
			{
				name:           "Agile",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Agility", dice.D6, ca)
					},
				},
			},
			{
				name:           "All Thumbs",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return addHindranceModBuilder("All Thumbs", Minor, SwadeHindrances, ca)
					},
				},
			},
		},
	},
	{
		name:        "Half Elves (Adaptable)",
		description: "",
		abilities: racialAbilities{
			{
				name:           "Adaptable",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					freeNoviceEdgeMod,
				},
			},
		},
	},
	{
		name:        "Half Elves (Agile)",
		description: "",
		abilities: racialAbilities{
			{
				name:           "Agile",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Agility", dice.D6, ca)
					},
				},
			},
		},
	},
	{
		name:        "Half-Folk",
		description: "Small but lucky",
		abilities: racialAbilities{
			//LUCK -> extra benny
			//REDUCED PACE: Decrease Pace by 1 and their running die one die type.
			{
				name:           "Size",
				description:    "Reducing Size (& therefore Toughness) by 1.",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					minusBaseToughnessMod,
					minusSizeMod,
				},
			},
			{
				name:           "Spirited",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Spirit", dice.D6, ca)
					},
				},
			},
		},
	},
	{
		name:        "Human",
		description: "boring",
		abilities: racialAbilities{
			{
				name:           "Adaptable",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					freeNoviceEdgeMod,
				},
			},
		},
	},
	{
		name:        "Rakashans",
		description: "animalistic",
		abilities: racialAbilities{
			{
				name:           "Agile",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Agility", dice.D6, ca)
					},
				},
			},
		},
	},
	{
		name:        "Saurians",
		description: "dragon like humanoids",
		abilities: racialAbilities{
			{
				name:           "Armor +2",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					plusArmorMod,
					plusArmorMod,
				},
			},
			//todo:
			//KEEN SENSES (Alertness Edge): Saurians have acute senses, giving them the Alertness Edge.
		},
	},
}

// SwadeAttributes which are predefined for the SWADE ruleset
var SwadeAttributes = Attributes{
	{Name: "Agility", description: ""},
	{Name: "Smarts", description: ""},
	{Name: "Spirit", description: ""},
	{Name: "Strength", description: ""},
	{Name: "Vigor", description: ""},
}

// SwadeSkills which are predefined for the SWADE ruleset
var SwadeSkills = Skills{
	{Name: "Academics", LinkedAttribute: "Smarts", IsCore: false, description: "Academics reflects knowledge of the liberal arts, social sciences, literature, history, archaeology, and similar fields. If an explorer wants to remember when the Mayan calendar ended or cite a line from Macbeth, this is the skill to have."},
	{Name: "Athletics", LinkedAttribute: "Agility", IsCore: true, description: ""},
	{Name: "Battle", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Boating", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Common Knowledge", LinkedAttribute: "Smarts", IsCore: true, description: ""},
	{Name: "Driving", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Electronics", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Faith", LinkedAttribute: "Spirit", IsCore: false, description: ""},
	{Name: "Fighting", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Focus", LinkedAttribute: "Spirit", IsCore: false, description: ""},
	{Name: "Gambling", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Hacking", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Healing", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Intimidation", LinkedAttribute: "Spirit", IsCore: false, description: ""},
	{Name: "Language (X)", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Notice", LinkedAttribute: "Smarts", IsCore: true, description: ""},
	{Name: "Occult", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Performance", LinkedAttribute: "Spirit", IsCore: false, description: ""},
	{Name: "Persuasion", LinkedAttribute: "Spirit", IsCore: true, description: ""},
	{Name: "Piloting", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Psionics", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Repair", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Research", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Riding", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Science", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Shooting", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Spellcasting", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Stealth", LinkedAttribute: "Agility", IsCore: true, description: ""},
	{Name: "Survival", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Taunt", LinkedAttribute: "Smarts", IsCore: false, description: ""},
	{Name: "Thievery", LinkedAttribute: "Agility", IsCore: false, description: ""},
	{Name: "Weird Science", LinkedAttribute: "Smarts", IsCore: false, description: ""},
}

// SwadeHindrances which are predefined for the SWADE ruleset
var SwadeHindrances = Hindrances{
	{Name: "All Thumbs", description: "–2 to use mechanical or electrical devices.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Poverty", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Habit", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
	{Name: "Mean", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	//Can’t Swim (Minor): –2 to swimming (contained in skill Athletiks)
	{Name: "Can't Swim", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{
		Name:        "Clueless",
		description: "Clueless (Major): –1 to Common Knowledge and Notice rolls.",
		AvailableDegrees: []HindranceDegree{{Degree: Major,
			Modifiers: CharacterAggregationModifiers{
				func(ca CharacterAggregation) CharacterAggregation {
					return skillAdjusmentModBuilder("Common Knowledge", -1, SwadeSkills, ca)
				},
				func(ca CharacterAggregation) CharacterAggregation {
					return skillAdjusmentModBuilder("Notice", -1, SwadeSkills, ca)
				},
			}}},
	},
	//Clumsy (Major): –2 to Athletics and Stealth rolls.
	//Obese (Minor): Size +1, Pace –1 and running die of d4. Treat Str as one die type lower for Min Str.
	/*
	 * Small (Minor): Size and Toughness are reduced by 1. Size cannot be reduced below –1.
	 * race aquarian +1 toughness
	 */
	//Young (Minor/Major): Minor has 4 attribute points and 10 skill points, extra Benny per session. Major has 3 attribute points, 10 skill points, and two extra Bennies per session.

	{
		Name:             "Pacifist",
		description:      "Fights only in self-defense as a Minor Hindrance, won’t fight at all as Major.",
		AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}},
	},
}

// SawadeEdges which are predefined for the SWADE ruleset
var SawadeEdges = Edges{
	{
		name:        "Alertness",
		requirement: Requirement{level: Novice},
		modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Notice", 2, SwadeSkills, ca)
			},
		},
	},
}
