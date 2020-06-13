package rulebook

type Rulebook struct {
	traits Traits
}

// New returns a rulebook with basic SWADE rules and related data
func New(a []Attribute, s []Skill) Rulebook {
	return Rulebook{
		traits: Traits{
			Attributes: a,
			Skills:     s,
		},
	}
}

//Traits returns traits, containing e.g. attributes and skills
func (rb Rulebook) Traits() Traits {
	return rb.traits
}
