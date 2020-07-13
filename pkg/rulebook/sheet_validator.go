package rulebook

import (
	"fmt"
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

		SheetChosenEdges:     Edges{},
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
		ca.CoreValidators["permittedHindrancesValidator"] = permittedHindrancesValidator
		ca.CoreValidators["requiredAttributesValidator"] = requiredAttributesValidator
		ca.CoreValidators["attributePointsValidator"] = attributePointsValidator
		ca.CoreValidators["requiredCoreSkillsValidator"] = requiredCoreSkillsValidator
		ca.CoreValidators["permittedSkillsValidator"] = permittedSkillsValidator
		ca.CoreValidators["skillPointsValidator"] = skillPointsValidator

		return ca
	})

	aggregateMods, err := aggregate(charState, sheet, rb)
	if err != nil {
		return fmt.Errorf("aggregation error: %s", err)
	}
	charState.Updates(aggregateMods)

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
