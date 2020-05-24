package savage

type Attribute struct {
	name        string
	description string
}

var attributes = []Attribute{
	{name: "Agility", description: ""},
	{name: "Smarts", description: ""},
	{name: "Spirit", description: ""},
	{name: "Strength", description: ""},
	{name: "Vigor", description: ""},
}

// findAttribute returns index int and ok bool
func findAttribute(name string) (int, bool) {
	for i, attribute := range attributes {
		if attribute.name == name {
			return i, true
		}
	}

	return -1, false
}

type Skill struct {
	name            string
	linkedAttribute string
	isCore          bool
	description     string
}

var skills = []Skill{
	{name: "Academics", linkedAttribute: "Smarts", isCore: false, description: "Academics reflects knowledge of the liberal arts, social sciences, literature, history, archaeology, and similar fields. If an explorer wants to remember when the Mayan calendar ended or cite a line from Macbeth, this is the skill to have."},
	{name: "Athletics", linkedAttribute: "Agility", isCore: true, description: ""},
	{name: "Battle", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Boating", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Common Knowledge", linkedAttribute: "Smarts", isCore: true, description: ""},
	{name: "Driving", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Electronics", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Faith", linkedAttribute: "Spirit", isCore: false, description: ""},
	{name: "Fighting", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Focus", linkedAttribute: "Spirit", isCore: false, description: ""},
	{name: "Gambling", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Hacking", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Healing", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Intimidation", linkedAttribute: "Spirit", isCore: false, description: ""},
	{name: "Language (X)", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Notice", linkedAttribute: "Smarts", isCore: true, description: ""},
	{name: "Occult", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Performance", linkedAttribute: "Spirit", isCore: false, description: ""},
	{name: "Persuasion", linkedAttribute: "Spirit", isCore: true, description: ""},
	{name: "Piloting", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Psionics", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Repair", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Research", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Riding", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Science", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Shooting", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Spellcasting", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Stealth", linkedAttribute: "Agility", isCore: true, description: ""},
	{name: "Survival", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Taunt", linkedAttribute: "Smarts", isCore: false, description: ""},
	{name: "Thievery", linkedAttribute: "Agility", isCore: false, description: ""},
	{name: "Weird Science", linkedAttribute: "Smarts", isCore: false, description: ""},
}

func coreSkills() (coreSkills []Skill) {
	for _, skill := range skills {
		if skill.isCore {
			coreSkills = append(coreSkills, skill)
		}
	}

	return coreSkills
}

// findSkill returns index int and ok bool
func findSkill(name string) (int, bool) {
	for i, skill := range skills {
		if skill.name == name {
			return i, true
		}
	}

	return -1, false
}
