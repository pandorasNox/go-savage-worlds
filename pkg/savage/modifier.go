package savage

type Modifier struct {
	kind     ModifierKind
	value    int
	selector Selector
}

type ModifierKind int

const (
	ModifierKindDice ModifierKind = iota
	ModifierKindAccumulator
)

func (mk ModifierKind) String() string {
	return [...]string{"dice", "accumulator"}[mk]
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
