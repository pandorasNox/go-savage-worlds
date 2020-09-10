package rulebook

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

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
	name        edgeName
	description string
	requirement Requirement
	edgeType    EdgeType
	modifiers   CharacterAggregationModifiers
}

type edgeName string

type Requirement struct {
	rank       Rank			//used to filter/find applied edges
	validators validators
}

// Edges list of multiple edge
type Edges []Edge

//FindEdge returns index int and found bool
func (es Edges) FindEdge(name edgeName) (index int, found bool) {
	for i, edge := range es {
		if edge.name == name {
			return i, true
		}
	}

	return -1, false
}

func minimumRankValidatorBuilder(rank Rank, edgeName edgeName) validator {
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

func minimumAttributeValidatorBuilder(
	attributeName AttributeName,
	minNeededDice dice.Dice,
	edgeName edgeName,
) validator {
	return func(ca CharacterAggregation, s Sheet, rb Rulebook) error {
		_, found := rb.Traits().Attributes.FindAttribute(string(attributeName))
		if found == false {
			return fmt.Errorf("couldn't find attribute \"%s\" in rulebook for min points required validation", attributeName)
		}

		sheetAttribute, err := s.SheetAttribute(attributeName)
		if err != nil {
			return fmt.Errorf("%s: for min points required validation", err)
		}

		aDice, err := dice.Parse(sheetAttribute.Dice)
		if err != nil {
			return fmt.Errorf("couldn't parse dice \"%s\" from sheet attribute \"%s\" for min points required validation", sheetAttribute.Dice, sheetAttribute.Name)
		}

		if minNeededDice.Points() > aDice.Points() {
			return fmt.Errorf("for edge \"%s\" a minimum required dice level the attribute \"%s\" is:\"d%s\"", edgeName, attributeName, minNeededDice.Value())
		}

		return nil
	}
}
