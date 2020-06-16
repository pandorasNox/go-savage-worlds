package sheet

import "github.com/pandorasNox/go-savage-worlds/pkg/rulebook"

//CharacterAggregation reflect the values of rules and sheet
type CharacterAggregation struct {
	AttributePointsAvailable  int
	AttributePointsAggregated int

	SkillPointsAvailable  int
	SkillPointsAggregated int

	HindrancePointsAvailable  int
	HindrancePointsAggregated int
	HindrancesRequired        rulebook.Hindrances
	HindrancesIgnored         rulebook.Hindrances
}

//CharacterAggregationState reflects current state of character aggregation
type CharacterAggregationState struct {
	characterAggregation CharacterAggregation
}

//Update the character aggregation state via the provided method
func (cas *CharacterAggregationState) Update(cam CharacterAggregationModifier) {
	cas.characterAggregation = cam(cas.characterAggregation)
}

//CharacterAggregationModifier the method used to update the character aggregation
type CharacterAggregationModifier func(CharacterAggregation) CharacterAggregation
