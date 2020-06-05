package savage

type Modifier struct {
	kind     ModifierKind
	value    int
	selector Selector
}

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
