package rulebook

type Traits struct {
	Attributes []Attribute
	Skills     []Skill
}

type Attribute struct {
	Name        string
	description string
}

// FindAttribute returns index int and ok bool
func (rb Rulebook) FindAttribute(name string) (int, bool) {
	for i, attribute := range rb.Traits().Attributes {
		if attribute.Name == name {
			return i, true
		}
	}

	return -1, false
}

type Skill struct {
	Name            string
	LinkedAttribute string
	IsCore          bool
	description     string
}

func (rb Rulebook) CoreSkills() (coreSkills []Skill) {
	coreSkills = []Skill{}

	for _, skill := range rb.Traits().Skills {
		if skill.IsCore {
			coreSkills = append(coreSkills, skill)
		}
	}

	return coreSkills
}

// FindSkill returns index int and ok bool
func (rb Rulebook) FindSkill(name string) (int, bool) {
	for i, skill := range rb.Traits().Skills {
		if skill.Name == name {
			return i, true
		}
	}

	return -1, false
}
