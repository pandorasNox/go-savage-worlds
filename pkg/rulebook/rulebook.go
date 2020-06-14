package rulebook

//Rulebook defines a rulebook
type Rulebook struct {
	traits Traits
	hindrances Hindrances
}

//New returns a rulebook with basic SWADE rules and related data
func New(a Attributes, s Skills, h Hindrances) Rulebook {
	return Rulebook{
		traits: Traits{
			Attributes: a,
			Skills:     s,
		},
		hindrances: h,
	}
}

//Traits returns traits, containing e.g. attributes and skills
func (rb Rulebook) Traits() Traits {
	return rb.traits
}

//Hindrances returns hindrances
func (rb Rulebook) Hindrances() Hindrances {
	return rb.hindrances
}
