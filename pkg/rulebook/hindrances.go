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

var Hindrances = []Hindrance{
	{Name: "Poverty", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{Name: "Habit", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
	{Name: "Mean", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	//Can’t Swim (Minor): –2 to swimming (contained in skill Athletiks)
	{Name: "Can't Swim", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
	{
		Name:        "Clueless",
		description: "Clueless (Major): –1 to Common Knowledge and Notice rolls.",
		AvailableDegrees: []HindranceDegree{{Degree: Major,
			Modifiers: []Modifier{
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

//find vs indexOf
// findHindrance returns index int and ok bool
func FindHindrance(name string) (index int, ok bool) {
	for i, hindrance := range Hindrances {
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
