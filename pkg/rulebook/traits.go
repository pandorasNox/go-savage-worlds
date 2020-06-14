package rulebook

type Traits struct {
	Attributes Attributes
	Skills     Skills
}

type Attribute struct {
	Name        string
	description string
}

type Attributes []Attribute

// FindAttribute returns index and found bool
func (attrs Attributes) FindAttribute(name string) (index int, found bool) {
	for i, attribute := range attrs {
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

type Skills []Skill

//CoreSkills returns slice of skills which have isCore true
func (ss Skills) CoreSkills() (coreSkills Skills) {
	coreSkills = Skills{}

	for _, skill := range ss {
		if skill.IsCore {
			coreSkills = append(coreSkills, skill)
		}
	}

	return coreSkills
}

// FindSkill returns index int and found bool
func (ss Skills) FindSkill(name string) (index int, found bool) {
	for i, skill := range ss {
		if skill.Name == name {
			return i, true
		}
	}

	return -1, false
}
