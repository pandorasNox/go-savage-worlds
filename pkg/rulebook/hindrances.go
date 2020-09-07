package rulebook

type Hindrance struct {
	Name             string
	description      string
	AvailableDegrees []HindranceDegree
	//todo: add an `isSituational` flag???
	//(for hind. which has only influencs (+/- something) in special situations)
}

type Hindrances []Hindrance

type HindranceDegree struct {
	Degree    Degree
	Modifiers CharacterAggregationModifiers
}

type Degree int

const (
	Major Degree = iota
	Minor
)

func (d Degree) String() string {
	return []string{"major", "minor"}[d]
}

//FindHindrance returns index int and found bool
func (hs Hindrances) FindHindrance(name string) (index int, found bool) {
	for i, hindrance := range hs {
		if hindrance.Name == name {
			return i, true
		}
	}

	return -1, false
}

//FindDegree returns index int and found bool
func (h Hindrance) FindDegree(degreeName string) (index int, found bool) {
	for i, hd := range h.AvailableDegrees {
		if hd.Degree.String() == degreeName {
			return i, true
		}
	}

	return -1, false
}
