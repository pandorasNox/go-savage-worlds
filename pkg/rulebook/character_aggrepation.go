package rulebook

//CharacterAggregation reflect the values of rules and sheet
type CharacterAggregation struct {
	AttributePointsAvailable          int
	AttributePointsUsed               int
	MinimumAttributePointsRequiredFor MinimumAttributePointsRequiredFor

	SkillPointsAvailable          int
	SkillPointsUsed               int
	MinimumSkillPointsRequiredFor MinimumSkillPointsRequiredFor
	SkillsAdjustments             SkillsAdjustments

	HindrancePointsEarnedLimit int
	HindrancePointsEarned      int
	HindrancePointsUsed        int
	HindrancesRequired         AggregatedHindrances
	//ignored for hindrancePoints aggregation
	HindrancesIgnored AggregatedHindrances

	//edges
	SheetChosenEdges     Edges
	MinimumChosenEdges   int
	EdgesRequired        Edges
	CoreValidators       coreValidators
	additionalValidators validators

	//other
	Size                    int
	BaseToughness           int
	Armor                   int
	ShakenRecoveryAdjusment int
}

type HindranceName string
type AggregatedHindrances map[HindranceName]Degree

type SkillName string
type MinimumSkillPointsRequiredFor map[SkillName]int

type AttributeName string
type MinimumAttributePointsRequiredFor map[AttributeName]int

type SkillsAdjustments map[SkillName]int

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

// CharacterAggregation returns CharacterAggregation
func (cas CharacterAggregationState) CharacterAggregation() CharacterAggregation {
	return cas.characterAggregation
}

//SkillPointsAvailable returns available skill points
func (cas CharacterAggregationState) SkillPointsAvailable() (pointsAvailable int) {
	return cas.characterAggregation.SkillPointsAvailable
}

//SkillPointsUsed returns used skill points
func (cas CharacterAggregationState) SkillPointsUsed() (pointsUsed int) {
	return cas.characterAggregation.SkillPointsUsed
}

type validator func(ca CharacterAggregation, s Sheet, rb Rulebook) error

// validators list of functions to validate
type validators []validator

type validatorIdentifier string
type coreValidators map[validatorIdentifier]validator

// Validators returns all validators
func (cas CharacterAggregationState) validators() validators {
	var v validators

	for _, validator := range cas.characterAggregation.CoreValidators {
		v = append(v, validator)
	}

	v = append(v, cas.characterAggregation.additionalValidators...)

	return v
}

// Validate all the things
func (cas CharacterAggregationState) Validate(s Sheet, rb Rulebook) (errors []error) {
	for _, v := range cas.validators() {
		err := v(cas.characterAggregation, s, rb)

		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
