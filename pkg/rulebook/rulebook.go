package rulebook

//Rulebook defines a rulebook
type Rulebook struct {
	races      Races
	traits     Traits
	hindrances Hindrances
}

//New returns a rulebook with basic SWADE rules and related data
func New(r Races, a Attributes, s Skills, h Hindrances) Rulebook {
	return Rulebook{
		races: r,
		traits: Traits{
			Attributes: a,
			Skills:     s,
		},
		hindrances: h,
	}
}

// Races returns races
func (rb Rulebook) Races() Races {
	return rb.races
}

//Traits returns traits, containing e.g. attributes and skills
func (rb Rulebook) Traits() Traits {
	return rb.traits
}

//Hindrances returns hindrances
func (rb Rulebook) Hindrances() Hindrances {
	return rb.hindrances
}
