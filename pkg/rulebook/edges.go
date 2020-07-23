package rulebook

type Rank int

const (
	Novice Rank = iota
)

// Requirement to apply the edge
type Requirement struct {
	rank Rank
	//edgeDeps
	//skill & attr level (probbably mod)

	// todo: transform to additinalValidator
}

// EdgeType group
type EdgeType int

const (
	BackgroundEdge EdgeType = iota
)

//Edge defines the edges
type Edge struct {
	name        string
	description string
	requirement Requirement
	edgeType    EdgeType
	modifiers   CharacterAggregationModifiers
}

// Edges list of multiple edge
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
