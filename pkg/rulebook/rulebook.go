package rulebook

//Rulebook defines a rulebook
type Rulebook struct {
	races      Races
	hindrances Hindrances
	traits     Traits
	edges      Edges
}

//New returns a rulebook with basic SWADE rules and related data
func New(r Races, h Hindrances, a Attributes, s Skills, e Edges) Rulebook {
	return Rulebook{
		races:      r,
		hindrances: h,
		traits: Traits{
			Attributes: a,
			Skills:     s,
		},
		edges: e,
	}
}

// Races returns races
func (rb Rulebook) Races() Races {
	return rb.races
}

// Hindrances returns hindrances
func (rb Rulebook) Hindrances() Hindrances {
	return rb.hindrances
}

// Traits returns traits, containing e.g. attributes and skills
func (rb Rulebook) Traits() Traits {
	return rb.traits
}

// Edges returns edges
func (rb Rulebook) Edges() Edges {
	return rb.edges
}
