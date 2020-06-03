package savage

type Modifier struct {
	kind     string
	value    int
	selector Selector
}

const (
	MODIFIER_KIND_DICE         = "dice"
	MODIFIER_KIND_ACCUMULATION = "accumulation"
)

type Selector struct {
	kind   string
	target string
}

const (
	SELECTOR_KIND_SKILL     = "skill"
	SELECTOR_KIND_ATTRIBUTE = "attribute"
)
