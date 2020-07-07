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
		coreValidators:             make(coreValidators),
		additionalValidators:       validators{},
	}

	charState := CharacterAggregationState{}
	//todo: init func instead?!
	charState.Update(func(_ CharacterAggregation) CharacterAggregation {
		return initCharAggegation
	})

	aggregateMod, err := aggregate(charState, sheet, rb)
	if err != nil {
		return fmt.Errorf("aggregation error: %s", err)
	}
	charState.Update(aggregateMod)

	err = validatePermittedHindrances(sheet, rb.Hindrances())
	if err != nil {
		return fmt.Errorf("sheet validation hindrance errors: %s", err)
	}

	charState.Update(func(currentState CharacterAggregation) CharacterAggregation {
		currentState.HindrancePointsEarned = sheet.countHindrancePoints()
		return currentState
	})

	modifiers := sheet.collectModifier(rb)
	charState.Updates(modifiers)

	errors := charState.Validate()
	if errors != nil {
		var sErrors string = ""

		for _, err := range errors {
			sErrors = fmt.Sprintf("%s  %s\n", sErrors, err)
		}

		return fmt.Errorf("aggregation validation failed:\n%s", sErrors)
	}

	err = validateAttributes(sheet, rb.Traits().Attributes, charState)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = validateSkills(sheet, rb.Traits().Skills, charState)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func aggregate(cas CharacterAggregationState, s Sheet, rb Rulebook) (CharacterAggregationModifier, error) {
	var err error
	emptyFn := func(ca CharacterAggregation) CharacterAggregation {
		return ca
	}
	var fn CharacterAggregationModifier

	attributePointsUsed, err := aggregateAttributePointsUsed(s, rb.Traits().Attributes)
	if err != nil {
		return emptyFn, err
	}

	skillPointsUsed, err := aggregateSkillPointsUsed(s, rb.Traits().Skills)
	if err != nil {
		return emptyFn, err
	}

	fn = func(currentState CharacterAggregation) CharacterAggregation {
		currentState.AttributePointsUsed = attributePointsUsed
		currentState.SkillPointsUsed = skillPointsUsed
		return currentState
	}

	return fn, nil
}

func aggregateAttributePointsUsed(sheet Sheet, attributes Attributes) (pointsUsed int, err error) {
	attributePointsUsed := 0
	for _, attribute := range sheet.Character.Traits.Attributes {
		_, ok := attributes.FindAttribute(attribute.Name)
		if ok == false {
			return 0, fmt.Errorf("\"%s\" is no valid attribute", attribute.Name)
		}

		dice, err := dice.Parse(attribute.Dice)
		if err != nil {
			return 0, fmt.Errorf(
				"parsing dice for attribute \"%s\" failed: %s",
				attribute.Name, err,
			)
		}

		attributePointsUsed += dice.Points()
	}

	return attributePointsUsed, nil
}

func aggregateSkillPointsUsed(sheet Sheet, skills Skills) (pointsUsed int, err error) {
	skillPointsUsed := 0

	for _, sheetAttr := range sheet.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, _ := skills.FindSkill(sheetSkill.Name)
			skill := skills[index]

			dice, err := dice.Parse(sheetSkill.Dice)
			if err != nil {
				return 0, fmt.Errorf(
					"parsing dice for skill \"%s\" failed: %s",
					sheetSkill.Name, err,
				)
			}

			pointCostModifier := 1
			if skill.IsCore {
				pointCostModifier = 0
			}

			skillPointsUsed += dice.Points() + pointCostModifier
		}
	}

	return skillPointsUsed, nil
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
	if charState.SkillPointsUsed() > charState.SkillPointsAvailable() {
		return fmt.Errorf(
			"validation error: Used %d of %d available skill points",
			charState.SkillPointsUsed(),
			charState.SkillPointsAvailable(),
		)
	}

	return nil
}
