package rulebook

import "log"

type Modifier struct {
	kind     ModifierKind
	value    int
	selector Selector
}

type Modifiers []Modifier

type ModifierKind int

const (
	ModifierKindDiceValue ModifierKind = iota
	ModifierKindDiceAdjustment
)

func (mk ModifierKind) String() string {
	return [...]string{"diceValue", "diceAdjustment"}[mk]
}

type Selector struct {
	kind   SelectorKind
	target string
}

type SelectorKind int

const (
	SelectorKindAttribute SelectorKind = iota
	SelectorKindSkill
)

func (sk SelectorKind) String() string {
	return [...]string{"attribute", "skill"}[sk]
}

func modifierAddIgnoredPacifistHindrance(ca CharacterAggregation) CharacterAggregation {
	hindranceName := HindranceName("Pacifist")
	wantDegree := Major

	hIndex, hFound := SwadeHindrances.FindHindrance(string(hindranceName))
	if hFound == false {
		log.Fatalf("couldn't find %s hindrance in application data for modifierAddIgnoredPacifistHindrance function", hindranceName)
	}

	pacifistHindrance := SwadeHindrances[hIndex]

	_, dFound := pacifistHindrance.FindDegree(wantDegree.String())
	if dFound == false {
		log.Fatalf("couldn't find degree %s for %s hindrance in application data for modifierAddIgnoredPacifistHindrance function", wantDegree.String(), hindranceName)
	}

	ca.HindrancesRequired[hindranceName] = wantDegree
	ca.HindrancesIgnored[hindranceName] = wantDegree

	return ca
}
