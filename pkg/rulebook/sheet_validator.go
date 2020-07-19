package rulebook

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

//Validate validates a savage world sheet
func Validate(sheet Sheet, rb Rulebook) error {
	var err error

	initCharAggegation := CharacterAggregation{
		AttributePointsAvailable:          BaseAttributePoints,
		AttributePointsUsed:               0,
		MinimumAttributePointsRequiredFor: make(MinimumAttributePointsRequiredFor),

		SkillPointsAvailable:          BaseSkillPoints,
		SkillPointsUsed:               0,
		MinimumSkillPointsRequiredFor: make(MinimumSkillPointsRequiredFor),
		SkillsAdjustments:             make(SkillsAdjustments),

		HindrancePointsEarnedLimit: 4,
		HindrancePointsEarned:      0,
		HindrancePointsUsed:        0,
		HindrancesRequired:         make(AggregatedHindrances),
		HindrancesIgnored:          make(AggregatedHindrances),

		MinimumChosenEdges:   0,
		EdgesRequired:        Edges{},
		CoreValidators:       make(coreValidators),
		additionalValidators: validators{},

		Size:                    0,
		BaseToughness:           0,
		ShakenRecoveryAdjusment: 0,
	}

	charState := CharacterAggregationState{}
	//todo: init func instead?!
	charState.Update(func(_ CharacterAggregation) CharacterAggregation {
		return initCharAggegation
	})

	charState.Update(func(ca CharacterAggregation) CharacterAggregation {
		ca.CoreValidators["requiredAttributesValidator"] = requiredAttributesValidator
		ca.CoreValidators["attributePointsValidator"] = attributePointsValidator

		ca.CoreValidators["requiredCoreSkillsValidator"] = requiredCoreSkillsValidator
		ca.CoreValidators["permittedSkillsValidator"] = permittedSkillsValidator
		ca.CoreValidators["skillPointsValidator"] = skillPointsValidator

		ca.CoreValidators["permittedHindrancesValidator"] = permittedHindrancesValidator
		ca.CoreValidators["hindrancePointsUsedValidator"] = hindrancePointsUsedValidator

		ca.CoreValidators["minimumAttributePointsRequiredForValidator"] = minimumAttributePointsRequiredForValidator
		ca.CoreValidators["minimumSkillPointsRequiredForValidator"] = minimumSkillPointsRequiredForValidator

		return ca
	})

	err = aggregateAndUpdate(&charState, sheet, rb)
	if err != nil {
		return fmt.Errorf("aggregation error: %s", err)
	}

	errors := charState.Validate(sheet, rb)
	if errors != nil {
		var sErrors string = ""

		for _, err := range errors {
			sErrors = fmt.Sprintf("%s  %s\n", sErrors, err)
		}

		return fmt.Errorf("aggregation validation failed:\n%s", sErrors)
	}

	return nil
}

func permittedHindrancesValidator(ca CharacterAggregation, s Sheet, rb Rulebook) error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := rb.Hindrances().FindHindrance(sheetHindrance.Name)

		if !ok {
			return fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		found := false
		for _, hd := range SwadeHindrances[index].AvailableDegrees {
			if hd.Degree.String() == sheetHindrance.Degree {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf(
				"\"%s\" is no valid degree of \"%s\"",
				sheetHindrance.Degree,
				SwadeHindrances[index].Name,
			)
		}
	}

	return nil
}

func hindrancePointsUsedValidator(ca CharacterAggregation, _ Sheet, _ Rulebook) error {
	hindrancePointsAvailable := min(ca.HindrancePointsEarned, ca.HindrancePointsEarnedLimit)
	if ca.HindrancePointsUsed > hindrancePointsAvailable {
		return fmt.Errorf(
			"Used %d of %d available hindrance points",
			ca.HindrancePointsUsed,
			hindrancePointsAvailable,
		)
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func requiredAttributesValidator(_ CharacterAggregation, s Sheet, rb Rulebook) error {
RequiredAttributes:
	for _, attribute := range rb.Traits().Attributes {
		for _, sheetAttribute := range s.Character.Traits.Attributes {
			if attribute.Name == sheetAttribute.Name {
				continue RequiredAttributes
			}
		}

		return fmt.Errorf("Attribute \"%s\" is a required attribute", attribute.Name)
	}

	return nil
}

func attributePointsValidator(ca CharacterAggregation, _ Sheet, _ Rulebook) error {
	if ca.AttributePointsUsed > ca.AttributePointsAvailable {
		return fmt.Errorf(
			"validation error: Used %d of %d available attribute points",
			ca.AttributePointsUsed,
			ca.AttributePointsAvailable,
		)
	}

	return nil
}

func requiredCoreSkillsValidator(_ CharacterAggregation, s Sheet, rb Rulebook) error {
RequiredCoreSkills:
	for _, coreSkill := range rb.traits.Skills.CoreSkills() {
		for _, sheetAttr := range s.Character.Traits.Attributes {
			for _, sheetSkill := range sheetAttr.Skills {
				if coreSkill.Name == sheetSkill.Name {
					continue RequiredCoreSkills
				}
			}
		}

		return fmt.Errorf("\"%s\" is a required core skill", coreSkill.Name)
	}

	return nil
}

func permittedSkillsValidator(_ CharacterAggregation, s Sheet, rb Rulebook) error {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, ok := rb.Traits().Skills.FindSkill(sheetSkill.Name)

			if !ok {
				return fmt.Errorf("\"%s\" is no valid skill", sheetSkill.Name)
			}

			if rb.Traits().Skills[index].LinkedAttribute != sheetAttr.Name {
				return fmt.Errorf(
					"\"%s\" should belong to attribute \"%s\" and not \"%s\"",
					sheetSkill.Name,
					rb.Traits().Skills[index].LinkedAttribute,
					sheetAttr.Name,
				)
			}

		}
	}

	return nil
}

func skillPointsValidator(ca CharacterAggregation, _ Sheet, _ Rulebook) error {
	if ca.SkillPointsUsed > ca.SkillPointsAvailable {
		return fmt.Errorf(
			"validation error: Used %d of %d available skill points",
			ca.SkillPointsUsed,
			ca.SkillPointsAvailable,
		)
	}

	return nil
}

func minimumAttributePointsRequiredForValidator(ca CharacterAggregation, s Sheet, rb Rulebook) error {
	for attributeName, minRequiredPoints := range ca.MinimumAttributePointsRequiredFor {
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

		if minRequiredPoints > aDice.Points() {
			minDice, err := dice.FromPoints(minRequiredPoints)
			return fmt.Errorf("for attribute \"%s\" a minimum dice of \"d%s\" is required. other err: %s", attributeName, minDice.Value(), err)
		}
	}

	return nil
}

func minimumSkillPointsRequiredForValidator(ca CharacterAggregation, s Sheet, rb Rulebook) error {
	for skillName, minRequiredPoints := range ca.MinimumSkillPointsRequiredFor {
		_, found := rb.Traits().Skills.FindSkill(string(skillName))
		if found == false {
			return fmt.Errorf("couldn't find skill \"%s\" in rulebook for min points required validation", skillName)
		}

		sheetSkill, err := s.SheetSkill(skillName)
		if err != nil {
			return fmt.Errorf("%s: for min points required validation", err)
		}

		sDice, err := dice.Parse(sheetSkill.Dice)
		if err != nil {
			return fmt.Errorf("couldn't parse dice \"%s\" from sheet skill \"%s\" for min points required validation", sheetSkill.Dice, sheetSkill.Name)
		}

		if minRequiredPoints > sDice.Points() {
			minDice, err := dice.FromPoints(minRequiredPoints)
			return fmt.Errorf("for skill \"%s\" a minimum dice of \"d%s\" is required. other err: %s", skillName, minDice.Value(), err)
		}
	}

	return nil
}
