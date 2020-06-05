package savage

type Hindrance struct {
	name             string
	description      string
	availableDegrees []HindranceDegree
}

type HindranceDegree struct {
	degree    Degree
	modifiers []Modifier
}

type Degree int

const (
	Major Degree = iota
	Minor
)

func (d Degree) String() string {
	return []string{"major", "minor"}[d]
}

var hindrances = []Hindrance{
	{name: "Poverty", description: "", availableDegrees: []HindranceDegree{{degree: Minor}}},
	{name: "Habit", description: "", availableDegrees: []HindranceDegree{{degree: Major}, {degree: Minor}}},
	{name: "Mean", description: "", availableDegrees: []HindranceDegree{{degree: Minor}}},
	//Can’t Swim (Minor): –2 to swimming (contained in skill Athletiks)
	{name: "Can't Swim", description: "", availableDegrees: []HindranceDegree{{degree: Minor}}},
	{
		name:        "Clueless",
		description: "Clueless (Major): –1 to Common Knowledge and Notice rolls.",
		availableDegrees: []HindranceDegree{{degree: Major,
			modifiers: []Modifier{
				{kind: ModifierKindDiceAdjustment, value: -1, selector: Selector{kind: SelectorKindSkill, target: "Common Knowledge"}},
				{kind: ModifierKindDiceAdjustment, value: -1, selector: Selector{kind: SelectorKindSkill, target: "Notice"}},
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

func findDegree(hindrance Hindrance, degreeName string) (int, bool) {
	for i, hd := range hindrance.availableDegrees {
		if hd.degree.String() == degreeName {
			return i, true
		}
	}

	return -1, false
}
