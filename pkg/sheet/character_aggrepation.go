package sheet

import "github.com/pandorasNox/go-savage-worlds/pkg/rulebook"

//CharacterAggregation reflect the values of rules and sheet
type CharacterAggregation struct {
	AttributePointsAvailable int
	AttributePointsUsed      int

	SkillPointsAvailable int
	SkillPointsUsed      int

	HindrancePointsLimit  int
	HindrancePointsEarned int
	HindrancePointsUsed   int
	HindrancesRequired    rulebook.Hindrances
	//ignored for hindrancePoints aggregation
	HindrancesIgnored rulebook.Hindrances
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

//AttributePointsAvailable returns available attribute points
func (cas CharacterAggregationState) AttributePointsAvailable() (pointsAvailable int) {
	return cas.characterAggregation.AttributePointsAvailable
}

//AttributePointsUsed returns used attribute points
func (cas CharacterAggregationState) AttributePointsUsed() (pointsUsed int) {
	return cas.characterAggregation.AttributePointsUsed
}

//SkillPointsAvailable returns available skill points
func (cas CharacterAggregationState) SkillPointsAvailable() (pointsAvailable int) {
	return cas.characterAggregation.SkillPointsAvailable
}

//SkillPointsUsed returns used skill points
func (cas CharacterAggregationState) SkillPointsUsed() (pointsUsed int) {
	return cas.characterAggregation.SkillPointsUsed
}
