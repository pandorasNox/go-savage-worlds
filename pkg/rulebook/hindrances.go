package rulebook

type Hindrance struct {
	Name             string
	description      string
	AvailableDegrees []HindranceDegree
}

type HindranceDegree struct {
	Degree    Degree
	Modifiers []Modifier
}

type Degree int

const (
	Major Degree = iota
	Minor
)

func (d Degree) String() string {
	return []string{"major", "minor"}[d]
}

//find vs indexOf
// findHindrance returns index int and ok bool
func FindHindrance(name string) (index int, ok bool) {
	for i, hindrance := range SwadeHindrances {
		if hindrance.Name == name {
			return i, true
		}
	}

	return -1, false
}

func FindDegree(hindrance Hindrance, degreeName string) (int, bool) {
	for i, hd := range hindrance.AvailableDegrees {
		if hd.Degree.String() == degreeName {
			return i, true
		}
	}

	return -1, false
}
