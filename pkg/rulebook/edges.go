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
