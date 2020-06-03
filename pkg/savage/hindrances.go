package savage

type Hindrance struct {
	name             string
	description      string
	availableDegrees []Degree
}

type Degree struct {
	value     string
	modifiers []Modifier
}

const (
	DEGREE_MAJOR = "major"
	DEGREE_MINOR = "minor"
)

// var allowedDegrees = []string{"DEGREE_MAJOR", "DEGREE_MINOR"}

var hindrances = []Hindrance{
	{name: "Poverty", description: "", availableDegrees: []Degree{{value: DEGREE_MINOR}}},
	{name: "Habit", description: "", availableDegrees: []Degree{{value: DEGREE_MAJOR}, {value: DEGREE_MINOR}}},
	{name: "Mean", description: "", availableDegrees: []Degree{{value: DEGREE_MINOR}}},
	//Can’t Swim (Minor): –2 to swimming (contained in skill Athletiks)
	{name: "Can't Swim", description: "", availableDegrees: []Degree{{value: DEGREE_MINOR}}},
	{
		name:        "Clueless",
		description: "Clueless (Major): –1 to Common Knowledge and Notice rolls.",
		availableDegrees: []Degree{{value: DEGREE_MAJOR,
			modifiers: []Modifier{
				{kind: MODIFIER_KIND_ACCUMULATION, value: -1, selector: Selector{kind: SELECTOR_KIND_SKILL, target: "Common Knowledge"}},
				{kind: MODIFIER_KIND_ACCUMULATION, value: -1, selector: Selector{kind: SELECTOR_KIND_SKILL, target: "Notice"}},
			}}},
	},
	//Clumsy (Major): –2 to Athletics and Stealth rolls.
	//Obese (Minor): Size +1, Pace –1 and running die of d4. Treat Str as one die type lower for Min Str.
	/*
	 * Small (Minor): Size and Toughness are reduced by 1. Size cannot be reduced below –1.
	 * race aquarian +1 toughness
	 */
	//Young (Minor/Major): Minor has 4 attribute points and 10 skill points, extra Benny per session. Major has 3 attribute points, 10 skill points, and two extra Bennies per session.
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

func findDegree(hindrance Hindrance, value string) (int, bool) {
	for i, degree := range hindrance.availableDegrees {
		if degree.value == value {
			return i, true
		}
	}

	return -1, false
}
