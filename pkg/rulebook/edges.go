package rulebook

import "fmt"

type Rank int

const (
	Novice Rank = iota
	Seasoned
)

var rankToString = []string{"Novice", "Seasoned"}

func (r Rank) String() string {
	return rankToString[r]
}

// aToRank create rank from string
func aToRank(rank string) (Rank, error) {
	for i, v := range rankToString {
		if v == rank {
			return Rank(i), nil
		}
	}

	return Rank(0), fmt.Errorf("\"%s\" is no valid rank", rank)
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

type Requirement struct {
	rank       Rank
	validators validators
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

func minimumRankValidatorBuilder(rank Rank, edgeName string) validator {
	return func(ca CharacterAggregation, s Sheet, rb Rulebook) error {
		sheetRank, err := aToRank(s.Character.Info.Rank)
		if err != nil {
			return err
		}

		if sheetRank < rank {
			return fmt.Errorf("edge \"%s\" requires rank \"%s\"", edgeName, rank.String())
		}

		return nil
	}
}
