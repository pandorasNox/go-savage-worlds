package rulebook

type CharacterLevel int

const (
	Novice CharacterLevel = iota
)

type Requirement struct {
	level CharacterLevel
	//edgeDeps
	//skill & attr level (probbably mod)

	// todo: transform to additinalValidator
}

type EdgeClassification int

const (
	BackgroundEdge EdgeClassification = iota
)

//Edge defines the edges
type Edge struct {
	name           string
	description    string
	requirement    Requirement
	classification EdgeClassification
	modifiers      CharacterAggregationModifiers
}

type Edges []Edge

//FindEdge returns index int and found bool
func (es Edges) FindEdge(name string) (index int, found bool) {
	for i, edge := range es {
		if edge.name == name {
			return i, true
		}
	}

	return -1, false
}
