package rulebook

//CharacterAggregation reflect the values of rules and sheet
type CharacterAggregation struct {
	AttributePointsAvailable          int
	AttributePointsUsed               int
	MinimumAttributePointsRequiredFor MinimumAttributePointsRequiredFor

	SkillPointsAvailable          int
	SkillPointsUsed               int
	MinimumSkillPointsRequiredFor MinimumSkillPointsRequiredFor

	HindrancePointsLimit  int
	HindrancePointsEarned int
	HindrancePointsUsed   int
	HindrancesRequired    AggregatedHindrances
	//ignored for hindrancePoints aggregation
	HindrancesIgnored AggregatedHindrances

	ToughnessAdjustment int

	ShakenRecoveryAdjusment int
}

type HindranceName string
type AggregatedHindrances map[HindranceName]Degree

type SkillName string
type MinimumSkillPointsRequiredFor map[SkillName]int

type AttributeName string
type MinimumAttributePointsRequiredFor map[AttributeName]int

//CharacterAggregationState reflects current state of character aggregation
type CharacterAggregationState struct {
	characterAggregation CharacterAggregation
}

//Update the character aggregation state via the provided functions
func (cas *CharacterAggregationState) Update(cam CharacterAggregationModifier) {
	cas.characterAggregation = cam(cas.characterAggregation)
}

//Updates the character aggregation state via the provided functions
func (cas *CharacterAggregationState) Updates(cams CharacterAggregationModifiers) {
	for _, m := range cams {
		cas.Update(m)
	}
}

//CharacterAggregationModifier the function used to update the character aggregation
type CharacterAggregationModifier func(CharacterAggregation) CharacterAggregation

//CharacterAggregationModifiers is a list of CharacterAggregationModifier
type CharacterAggregationModifiers []CharacterAggregationModifier

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
