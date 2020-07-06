package rulebook

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

//Validate validates a savage world sheet
func Validate(s Sheet, rb Rulebook) error {
	initCharAggegation := CharacterAggregation{
		AttributePointsAvailable: BaseAttributePoints,
		SkillPointsAvailable:     BaseSkillPoints,
		HindrancePointsLimit:     4,
		HindrancePointsEarned:    0,
		HindrancePointsUsed:      0,
		SkillsAdjustments:        make(SkillsAdjustments),
		coreValidators:           make(coreValidators),
		additionalValidators:     validators{},
	}

	charState := CharacterAggregationState{}
	//todo: init func instead?!
	charState.Update(func(_ CharacterAggregation) CharacterAggregation {
		return initCharAggegation
	})

	var err error

	err = validatePermittedHindrances(s, rb.Hindrances())
	if err != nil {
		return fmt.Errorf("sheet validation hindrance errors: %s", err)
	}

	charState.Update(func(currentState CharacterAggregation) CharacterAggregation {
		currentState.HindrancePointsEarned = s.countHindrancePoints()
		return currentState
	})

	modifiers := s.collectModifier(rb)
	charState.Updates(modifiers)

	errors := charState.Validate()
	if errors != nil {
		var sErrors string = ""

		for _, err := range errors {
			sErrors = fmt.Sprintf("%s  %s\n", sErrors, err)
		}

		return fmt.Errorf("aggregation validation failed:\n%s", sErrors)
	}

	err = validateAttributes(s, rb.Traits().Attributes, charState)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = validateSkills(s, rb.Traits().Skills, charState)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func validatePermittedHindrances(s Sheet, rbHinds Hindrances) error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := rbHinds.FindHindrance(sheetHindrance.Name)

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

func validateAttributes(s Sheet, rbAttrs Attributes, charState CharacterAggregationState) error {
	var err error

	err = validateAttributesExist(s, rbAttrs)
	if err != nil {
		return err
	}

	err = validateAttributePoints(s, rbAttrs, charState)
	if err != nil {
		return err
	}

	return nil
}

func validateAttributesExist(s Sheet, rbAttrs Attributes) error {
RequiredAttributes:
	for _, attribute := range rbAttrs {
		for _, sheetAttribute := range s.Character.Traits.Attributes {
			if attribute.Name == sheetAttribute.Name {
				continue RequiredAttributes
			}
		}

		return fmt.Errorf("\"%s\" is a required attribute", attribute.Name)
	}

	return nil
}

func validateAttributePoints(s Sheet, rbAttrs Attributes, charState CharacterAggregationState) error {
	for _, attribute := range s.Character.Traits.Attributes {
		_, ok := rbAttrs.FindAttribute(attribute.Name)
		if ok == false {
			return fmt.Errorf("\"%s\" is no valid attribute", attribute.Name)
		}

		dice, err := dice.Parse(attribute.Dice)
		if err != nil {
			return fmt.Errorf(
				"parsing dice for attribute \"%s\" failed: %s",
				attribute.Name, err,
			)
		}

		charState.Update(func(currentState CharacterAggregation) CharacterAggregation {
			currentState.AttributePointsUsed += dice.Points()
			return currentState
		})
	}

	if charState.AttributePointsUsed() > charState.AttributePointsAvailable() {
		return fmt.Errorf(
			"validation error: Used %d of %d available attribute points",
			charState.AttributePointsUsed(),
			charState.AttributePointsAvailable(),
		)
	}

	return nil
}

func validateSkills(s Sheet, rbs Skills, charState CharacterAggregationState) error {
	var err error

	err = validateCoreSkillsExist(s, rbs)
	if err != nil {
		return err
	}

	err = validatePermittedSkills(s, rbs)
	if err != nil {
		return err
	}

	err = validateSkillPoints(s, rbs, charState)
	if err != nil {
		return err
	}

	return nil
}

func validateCoreSkillsExist(s Sheet, rbs Skills) error {
RequiredCoreSkills:
	for _, coreSkill := range rbs.CoreSkills() {
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

func validatePermittedSkills(s Sheet, rbs Skills) error {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, ok := rbs.FindSkill(sheetSkill.Name)

			if !ok {
				return fmt.Errorf("\"%s\" is no valid skill", sheetSkill.Name)
			}

			if rbs[index].LinkedAttribute != sheetAttr.Name {
				return fmt.Errorf(
					"\"%s\" should belong to attribute \"%s\" and not \"%s\"",
					sheetSkill.Name,
					rbs[index].LinkedAttribute,
					sheetAttr.Name,
				)
			}

		}
	}

	return nil
}

func validateSkillPoints(s Sheet, rbs Skills, charState CharacterAggregationState) error {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, _ := rbs.FindSkill(sheetSkill.Name)
			skill := rbs[index]

			dice, err := dice.Parse(sheetSkill.Dice)
			if err != nil {
				return fmt.Errorf(
					"parsing dice for skill \"%s\" failed: %s",
					sheetSkill.Name, err,
				)
			}

			pointCostModifier := 1
			if skill.IsCore {
				pointCostModifier = 0
			}

			charState.Update(func(currentState CharacterAggregation) CharacterAggregation {
				currentState.SkillPointsUsed += dice.Points() + pointCostModifier
				return currentState
			})
		}
	}

	if charState.SkillPointsUsed() > charState.SkillPointsAvailable() {
		return fmt.Errorf(
			"validation error: Used %d of %d available skill points",
			charState.SkillPointsUsed(),
			charState.SkillPointsAvailable(),
		)
	}

	return nil
}