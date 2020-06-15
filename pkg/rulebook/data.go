package rulebook

// SwadeRaces predefined for the SWADE core ruleset
var SwadeRaces = Races{
	{
		name:        "Human",
		description: "boring",
		abilities: racialAbilities{
			{
				name:           "Adaptable",
				classification: Positive,
				modifiers:      Modifiers{},
			},
		},
	},
	{
		name:        "Android",
		description: "pacifistic...",
		abilities: racialAbilities{
			{
				name:           "Hindrance",
				classification: Negative,
				modifiers:      Modifiers{},
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
	{Name: "Poverty", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Habit", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
	{Name: "Mean", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	//Can’t Swim (Minor): –2 to swimming (contained in skill Athletiks)
	{Name: "Can't Swim", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{
		Name:        "Clueless",
		description: "Clueless (Major): –1 to Common Knowledge and Notice rolls.",
		AvailableDegrees: []HindranceDegree{{Degree: Major,
			Modifiers: []Modifier{
				{kind: ModifierKindDiceAdjustment, value: -1, selector: Selector{kind: SelectorKindSkill, target: "Common Knowledge"}},
				{kind: ModifierKindDiceAdjustment, value: -1, selector: Selector{kind: SelectorKindSkill, target: "Notice"}},
			}}},
	},
	//Clumsy (Major): –2 to Athletics and Stealth rolls.
	//Obese (Minor): Size +1, Pace –1 and running die of d4. Treat Str as one die type lower for Min Str.
	/*
	 * Small (Minor): Size and Toughness are reduced by 1. Size cannot be reduced below –1.
	 * race aquarian +1 toughness
	 */
	//Young (Minor/Major): Minor has 4 attribute points and 10 skill points, extra Benny per session. Major has 3 attribute points, 10 skill points, and two extra Bennies per session.
}
