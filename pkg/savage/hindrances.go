package savage

type Hindrance struct {
	name             string
	availableDegrees []string
}

var hindrances = []Hindrance{
	{name: "Poverty", availableDegrees: []string{"minor"}},
	{name: "Habit", availableDegrees: []string{"major", "minor"}},
	{name: "Mean", availableDegrees: []string{"minor"}},
}

// findHindrance returns index int and ok bool
func findHindrance(name string) (int, bool) {
	for i, hindrance := range hindrances {
		if hindrance.name == name {
			return i, true
		}
	}

	return -1, false
}
