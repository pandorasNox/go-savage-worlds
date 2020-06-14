package rulebook

// SWADE_Attributes which are predefined for the SWADE ruleset
var SWADE_Attributes = []Attribute{
	{Name: "Agility", description: ""},
	{Name: "Smarts", description: ""},
	{Name: "Spirit", description: ""},
	{Name: "Strength", description: ""},
	{Name: "Vigor", description: ""},
}

// SWADE_Skills which are predefined for the SWADE ruleset
var SWADE_Skills = []Skill{
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
