package rulebook

import (
	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

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
						return addRequiredHindranceModBuilder("Pacifist", Major, SwadeHindrances, ca)
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
			{
				name:           "Keen Senses (Skill)",
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
						return addRequiredHindranceModBuilder("All Thumbs", Minor, SwadeHindrances, ca)
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
			{
				name:           "Keen Senses (Edge)",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return addRequiredEdgeModBuilder("Alertness", SawadeEdges, ca)
					},
				},
			},
		},
	},
	{
		name:        "Gnomes",
		description: "",
		abilities: racialAbilities{
			{
				name:           "Smart",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return attributeStartsAtModBuilder("Smarts", dice.D6, ca)
					},
				},
			},
			{
				name:           "Keen Senses (Skill)",
				classification: Positive,
				modifiers: CharacterAggregationModifiers{
					func(ca CharacterAggregation) CharacterAggregation {
						return skillStartsAtModBuilder("Notice", dice.D6, ca)
					},
				},
			},
			{
				name:           "Size -1",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					minusSizeMod,
				},
			},
			{
				name:           "Frail",
				classification: Negative,
				modifiers: CharacterAggregationModifiers{
					minusBaseToughnessMod,
				},
			},
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
	{Name: "Anemic", description: "–2 Vigor when resisting Fatigue", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Arrogant", description: "Likes to dominate his opponent, challenge the most powerful foe in combat.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Bad Eyes", description: "–1 to all Trait rolls dependent on vision, or –2 as a Major Hindrance. Eyewear negates penalty but have a 50% chance of breaking when the hero suffers trauma.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Bad Luck", description: "The characters starts with one less Benny per session.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Big Mouth", description: "Unable to keep secrets and constantly gives away private information.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Blind", description: "–6 to all tasks that require vision (but choice of a free Edge to offset).", AvailableDegrees: []HindranceDegree{{
		Degree: Major,
		Modifiers: CharacterAggregationModifiers{
			freeEdgeMod,
		},
	}}},
	{Name: "Bloodthirsty", description: "Never takes prisoners.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Can't Swim", description: "–2 to swimming (Athletics) rolls; Each inch moved in water costs 3\" of Pace.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Cautious", description: "The character plans extensively and/or is overly careful.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
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
	{Name: "Clumsy", description: "–2 to Athletics and Stealth rolls.", AvailableDegrees: []HindranceDegree{{
		Degree: Major,
		Modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Athletics", -2, SwadeSkills, ca)
			},
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Stealth", -2, SwadeSkills, ca)
			},
		},
	}}},
	{Name: "Code of Honor", description: "The character keeps his word and acts like a gentleman.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Curious", description: "The character wants to know about everything.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Death Wish", description: "The hero wants to die after or while completing some epic task.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Delusional", description: "The individual believes something strange that causes him occasional or frequent trouble.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Doubting Thomas", description: "The character doesn't believe in the supernatural, often exposing him to unnecessary risks.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Driven", description: "The hero’s actions are driven by some important goal or belief.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Elderly", description: "–1 to Pace, running, Agility, Strength, and Vigor. Hero gets 5 extra skill points.", AvailableDegrees: []HindranceDegree{
		{
			Degree: Major,
			Modifiers: CharacterAggregationModifiers{
				func(ca CharacterAggregation) CharacterAggregation {
					return attributeAdjusmentModBuilder(AttributeName("Agility"), -1, SwadeAttributes, ca)
				},
				func(ca CharacterAggregation) CharacterAggregation {
					return attributeAdjusmentModBuilder(AttributeName("Strength"), -1, SwadeAttributes, ca)
				},
				func(ca CharacterAggregation) CharacterAggregation {
					return attributeAdjusmentModBuilder(AttributeName("Vigor"), -1, SwadeAttributes, ca)
				},
				plusSkillPointsAvailableMod,
				plusSkillPointsAvailableMod,
				plusSkillPointsAvailableMod,
				plusSkillPointsAvailableMod,
				plusSkillPointsAvailableMod,
			},
		},
	}},
	{Name: "Enemy", description: "The character has a recurring nemesis.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Greedy", description: "The individual is obsessed with wealth and material possessions.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Habit", description: "Addicted to something, suffers Fatigue if deprived.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Hard of Hearing", description: "–4 to Notice sounds; automatic failure if completely deaf.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Heroic", description: "The character always helps those in need.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Hesitant", description: "Draw two Action Cards and take the lowest (except Jokers, which may be kept).", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Illiterate", description: "The character cannot read or write.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Impulsive", description: "The hero leaps before he looks.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Jealous", description: "The individual covets what others have.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Loyal", description: "The hero is loyal to his friends and allies.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Mean", description: "–1 to Persuasion rolls.", AvailableDegrees: []HindranceDegree{{
		Degree: Minor,
		Modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Persuasion", -1, SwadeSkills, ca)
			},
		},
	}}},
	{Name: "Mild Mannered", description: "–2 to Intimidation rolls.", AvailableDegrees: []HindranceDegree{{
		Degree: Minor,
		Modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Intimidation", -2, SwadeSkills, ca)
			},
		},
	}}},
	{Name: "Mute", description: "The hero cannot speak.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Obese", description: "Size +1, Pace –1 and running die of d4. Treat Str as one die type lower for Min Str.", AvailableDegrees: []HindranceDegree{{
		Degree: Minor,
		Modifiers: CharacterAggregationModifiers{
			plusSizeMod,
			minusArmorRequiredStrenghtPointsCorrectionMod,
		},
	}}},
	{Name: "Obligation", description: "The character has a weekly obligation of 20 (Minor) to 40 (Major) hours.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "One Arm", description: "–4 to tasks (such as Athletics) that require two hands.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "One Eye", description: "–2 to actions at 5′′ (10 yards) or more distance.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{Name: "Outsider", description: "The character doesn't fit in to the local environment and subtracts 2 from Persuasion rolls. As a Major Hindrance she has no legal rights or other serious consequences.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Overconfident", description: "The hero believes she can do anything.", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
	{
		Name:             "Pacifist",
		description:      "Fights only in self-defense as a Minor Hindrance, won’t fight at all as Major.",
		AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}},
	},
	{Name: "Phobia", description: "The character is afraid of something, and subtracts –1/–2 from all Trait rolls in its presence.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Poverty", description: "Half starting funds and the character is always broke.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Quirk", description: "The individual has some minor but persistent foible that often annoys others.", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Ruthless", description: "The character does what it takes to get her way.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Secret", description: "The hero has a dark secret of some kind.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Shamed", description: "The individual is haunted by some tragic event from her past.", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	{Name: "Slow", description: "Pace –1, reduce running die one step. As Major, Pace –2, –2 to Athletics and rolls to resist Athletics. Neither may take the Fleet-Footed Edge.", AvailableDegrees: []HindranceDegree{
		{
			Degree: Minor,
			Modifiers: CharacterAggregationModifiers{
				func(ca CharacterAggregation) CharacterAggregation {
					return addRequiredEdgeModBuilder(edgeName("Fleet-Footed"), SawadeEdges, ca)
				},
			},
		},
		{
			Degree: Major,
			Modifiers: CharacterAggregationModifiers{
				func(ca CharacterAggregation) CharacterAggregation {
					return skillAdjusmentModBuilder(SkillName("Athletics"), -2, SwadeSkills, ca)
				},
				func(ca CharacterAggregation) CharacterAggregation {
					return addRequiredEdgeModBuilder(edgeName("Fleet-Footed"), SawadeEdges, ca)
				},
			},
		},
	}},
	// {Name: "", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
	//todo: continue here completing the list

	/*
	 * Small (Minor): Size and Toughness are reduced by 1. Size cannot be reduced below –1.
	 * race aquarian +1 toughness
	 */
	//Young (Minor/Major): Minor has 4 attribute points and 10 skill points, extra Benny per session. Major has 3 attribute points, 10 skill points, and two extra Bennies per session.
}

// SawadeEdges which are predefined for the SWADE ruleset
var SawadeEdges = Edges{
	{
		name:     "Alertness",
		edgeType: BackgroundEdge,
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Alertness")),
			},
		},
		modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Notice", 2, SwadeSkills, ca)
			},
		},
	},
	{
		name:     "Ambidextrous",
		edgeType: BackgroundEdge,
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Alertness")),
				minimumAttributeValidatorBuilder(AttributeName("Agility"), dice.D8, edgeName("Ambidextrous")),
			},
		},
		modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Notice", 2, SwadeSkills, ca)
			},
		},
	},
	{
		name:        "Arcane Background",
		edgeType:    BackgroundEdge,
		description: "Allows access to the Arcane Backgrounds listed in Chapter Five.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Arcane Resistance")),
			},
		},
	},
	{
		name:        "Arcane Resistance",
		edgeType:    BackgroundEdge,
		description: "Arcane skills targeting the hero suffer a −2 penalty (even if cast by allies!); magical damage is reduced by 2.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Arcane Resistance")),
				minimumAttributeValidatorBuilder(AttributeName("Spirit"), dice.D8, edgeName("Arcane Resistance")),
			},
		},
	},
	{
		name:        "Improved Arcane",
		edgeType:    BackgroundEdge,
		description: "As Arcane Resistance except penalty is increased to −4 and magical damage is reduced by 4.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Improved Arcane")),
				minimumAttributeValidatorBuilder(AttributeName("Spirit"), dice.D8, edgeName("Improved Arcane")),
			},
		},
	},
	{
		name:        "Aristocrat",
		edgeType:    BackgroundEdge,
		description: "+2 to Common Knowledge and networking with upper class.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Aristocrat")),
			},
		},
	},
	{
		name:        "Attractive",
		edgeType:    BackgroundEdge,
		description: "It's no secret people are more willing to help those they find physically attractive. Your character adds +1 to Performance and Persuasion rolls if the target is attracted to his general type (gender, sex, species, etc.).",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Attractive")),
				minimumAttributeValidatorBuilder(AttributeName("Vigor"), dice.D6, edgeName("Attractive")),
			},
		},
	},
	{
		name:        "Very Attractive",
		edgeType:    BackgroundEdge,
		description: "Your hero is drop-dead gorgeous. He increases his Performance and Persuasion bonus to +2.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Very Attractive")),
			},
		},
		modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Performance", 2, SwadeSkills, ca)
			},
			func(ca CharacterAggregation) CharacterAggregation {
				return skillAdjusmentModBuilder("Persuasion", 2, SwadeSkills, ca)
			},
		},
	},
	{
		name:        "Berserk",
		edgeType:    BackgroundEdge,
		description: "After being Shaken or Wounded, melee attacks must be Wild Attacks, +1 die type to Strength, +2 to Toughness, ignore one level of Wound penalties, Critical Failure on Fighting roll hits random target. Take Fatigue after every five consecutive rounds, may choose to end rage with Smarts roll –2.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Berserk")),
			},
		},
	},
	{
		name:        "Brave",
		edgeType:    BackgroundEdge,
		description: "Those with this Edge have learned to master their fear, or have dealt with so many horrors they've become jaded. These valiant explorers add +2 to Fear checks and subtract 2 from Fear Table results (see page 124).",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Brave")),
				minimumAttributeValidatorBuilder(AttributeName("Spirit"), dice.D6, edgeName("Brave")),
			},
		},
	},
	{
		name:        "Brawny",
		edgeType:    BackgroundEdge,
		description: "Your bruiser is very large or very fit. Her Size increases by +1 (and therefore Toughness by 1) and she treats her Strength as one die type higher when determining Encumbrance (page 67) and Minimum Strength to use armor, weapons, and equipment without a penalty (page 66). Brawny can't increase a character's Size above +3.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Brawny")),
				minimumAttributeValidatorBuilder(AttributeName("Vigor"), dice.D6, edgeName("Brawny")),
				minimumAttributeValidatorBuilder(AttributeName("Strength"), dice.D6, edgeName("Brawny")),
			},
		},
		modifiers: CharacterAggregationModifiers{
			func(ca CharacterAggregation) CharacterAggregation {
				return plusSizeMod(ca)
			},
			func(ca CharacterAggregation) CharacterAggregation {
				return plusBaseToughnessMod(ca)
			},
		},
	},
	{
		name:        "Brute",
		edgeType:    BackgroundEdge,
		description: "Link Athletics to Strength instead of Agility (including resistance). Short Range of any thrown item increased by +1. Double that for the adjusted Medium Range, and double again for Long Range.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Brute")),
				minimumAttributeValidatorBuilder(AttributeName("Vigor"), dice.D6, edgeName("Brute")),
				minimumAttributeValidatorBuilder(AttributeName("Strength"), dice.D6, edgeName("Brute")),
			},
		},
	},
	{
		name:        "Charismatic",
		edgeType:    BackgroundEdge,
		description: "Your hero is likable for some reason. She may be trustworthy or kind, or might just exude confidence and goodwill. You get one free reroll on Persuasion rolls.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Charismatic")),
				minimumAttributeValidatorBuilder(AttributeName("Spirit"), dice.D8, edgeName("Charismatic")),
			},
		},
	},
	{
		name:        "Elan",
		edgeType:    BackgroundEdge,
		description: "Elan means energy or spirit. +2 when spending a Benny to reroll a Trait roll.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Elan")),
				minimumAttributeValidatorBuilder(AttributeName("Spirit"), dice.D8, edgeName("Elan")),
			},
		},
	},
	{
		name:        "Fame",
		edgeType:    BackgroundEdge,
		description: "+1 Persuasion rolls when recognized (Common Knowledge), double usual fee for Performance.",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Fame")),
			},
		},
	},
	{
		name:        "Famous",
		edgeType:    BackgroundEdge,
		description: "+2 Persuasion when recognized, 5 × or more usual fee for Performance.",
		requirement: Requirement{
			rank: Seasoned,
			validators: validators{
				minimumRankValidatorBuilder(Seasoned, edgeName("Famous")),
			},
		},
	},
	{
		name:        "Fast Healer",
		edgeType:    BackgroundEdge,
		description: "+2 Vigor when rolling for natural healing; check every 3 days (instead of 5). (see Healing, page 96)",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumRankValidatorBuilder(Novice, edgeName("Fast Healer")),
				minimumAttributeValidatorBuilder(AttributeName("Vigor"), dice.D8, edgeName("Fast Healer")),
			},
		},
	},
	{
		name:        "Fleet-Footed",
		description: "The hero’s Pace is increased by +2 and his running die increases one step (from d6 to d8, for example).",
		requirement: Requirement{
			rank: Novice,
			validators: validators{
				minimumAttributeValidatorBuilder(AttributeName("Agility"), dice.D6, "Fleet-Footed"),
			},
		},
	},
	//Linguist
	//Luck
	////Great Luck
	//Quick
	//Rich
	////Filthy Rich
}
