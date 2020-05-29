package savage

type Modifier struct {
	value    int
	selector Selector
}

type Selector struct {
	kind   string
	target string
}
