package rulebook

type Rulebook struct {
	traits Traits
}

// New returns a rulebook with basic SWADE rules and related data
func New() Rulebook {
	return Rulebook{
		traits: Traits{
			Attributes: attributes,
			Skills:     skills,
		},
	}
}

//Traits returns traits, containing e.g. attributes and skills
func (rb Rulebook) Traits() Traits {
	return rb.traits
}
