package sheet

import (
	"fmt"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
	"github.com/pandorasNox/go-savage-worlds/pkg/rulebook"
)

//Validate validates a savage world sheet
func Validate(s Sheet, rb rulebook.Rulebook) error {
	availableAttributePoints := baseAttributePoints
	availableSkillPoints := baseSkillPoints
	earnedHindrancePoints := 0

	var err error

	err = validatePermittedHindrances(s, rb.Hindrances())
	if err != nil {
		return fmt.Errorf("sheet validation hindrance errors: %s", err)
	}

	modifier := s.collectModifier(rb)
	_ = modifier

	earnedHindrancePoints = s.countHindrancePoints()
	_ = earnedHindrancePoints

	err = validateAttributes(s, rb.Traits().Attributes, availableAttributePoints)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = validateSkills(s, rb.Traits().Skills, availableSkillPoints)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func validatePermittedHindrances(s Sheet, rbHinds rulebook.Hindrances) error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := rbHinds.FindHindrance(sheetHindrance.Name)

		if !ok {
			return fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		found := false
		for _, hd := range rulebook.SwadeHindrances[index].AvailableDegrees {
			if hd.Degree.String() == sheetHindrance.Degree {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf(
				"\"%s\" is no valid degree of \"%s\"",
				sheetHindrance.Degree,
				rulebook.SwadeHindrances[index].Name,
			)
		}
	}

	return nil
}

func validateAttributes(s Sheet, rbAttrs rulebook.Attributes, availableAttributePoints int) error {
	var err error

	err = validateAttributesExist(s, rbAttrs)
	if err != nil {
		return err
	}

	err = validateAttributePoints(s, rbAttrs, availableAttributePoints)
	if err != nil {
		return err
	}

	return nil
}

func validateAttributesExist(s Sheet, rbAttrs rulebook.Attributes) error {
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

func validateAttributePoints(s Sheet, rbAttrs rulebook.Attributes, availableAttributePoints int) error {
	aggregatedAttributePoints := 0

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

		aggregatedAttributePoints += dice.Points()
	}

	if aggregatedAttributePoints > availableAttributePoints {
		return fmt.Errorf(
			"validation error: Used %d of %d available attribute points",
			aggregatedAttributePoints,
			availableAttributePoints,
		)
	}

	return nil
}

func validateSkills(s Sheet, rbs rulebook.Skills, availableSkillPoints int) error {
	var err error

	err = validateCoreSkillsExist(s, rbs)
	if err != nil {
		return err
	}

	err = validatePermittedSkills(s, rbs)
	if err != nil {
		return err
	}

	err = validateSkillPoints(s, rbs, availableSkillPoints)
	if err != nil {
		return err
	}

	return nil
}

func validateCoreSkillsExist(s Sheet, rbs rulebook.Skills) error {
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

func validatePermittedSkills(s Sheet, rbs rulebook.Skills) error {
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

func validateSkillPoints(s Sheet, rbs rulebook.Skills, availableSkillPoints int) error {
	aggregatedSkillPoints := 0

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

			aggregatedSkillPoints += dice.Points() + pointCostModifier
		}
	}

	if aggregatedSkillPoints > availableSkillPoints {
		return fmt.Errorf(
			"validation error: Used %d of %d available skill points",
			aggregatedSkillPoints,
			availableSkillPoints,
		)
	}

	return nil
}
