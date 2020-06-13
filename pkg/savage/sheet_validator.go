package savage

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

	err = validatePermittedHindrances(s)
	if err != nil {
		return fmt.Errorf("sheet validation hindrance errors: %s", err)
	}

	modifier := s.collectModifier()
	_ = modifier

	earnedHindrancePoints = s.countHindrancePoints()
	_ = earnedHindrancePoints

	err = validateAttributes(s, rb, availableAttributePoints)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = validateSkills(s, rb, availableSkillPoints)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func validatePermittedHindrances(s Sheet) error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := rulebook.FindHindrance(sheetHindrance.Name)

		if !ok {
			return fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		found := false
		for _, hd := range rulebook.Hindrances[index].AvailableDegrees {
			if hd.Degree.String() == sheetHindrance.Degree {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf(
				"\"%s\" is no valid degree of \"%s\"",
				sheetHindrance.Degree,
				rulebook.Hindrances[index].Name,
			)
		}
	}

	return nil
}

func validateAttributes(s Sheet, rb rulebook.Rulebook, availableAttributePoints int) error {
	var err error

	err = validateAttributesExist(s, rb)
	if err != nil {
		return err
	}

	err = validateAttributePoints(s, rb, availableAttributePoints)
	if err != nil {
		return err
	}

	return nil
}

func validateAttributesExist(s Sheet, rb rulebook.Rulebook) error {
RequiredAttributes:
	for _, attribute := range rb.Traits().Attributes {
		for _, sheetAttribute := range s.Character.Traits.Attributes {
			if attribute.Name == sheetAttribute.Name {
				continue RequiredAttributes
			}
		}

		return fmt.Errorf("\"%s\" is a required attribute", attribute.Name)
	}

	return nil
}

func validateAttributePoints(s Sheet, rb rulebook.Rulebook, availableAttributePoints int) error {
	aggregatedAttributePoints := 0

	for _, attribute := range s.Character.Traits.Attributes {
		_, ok := rb.FindAttribute(attribute.Name)
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

		aggregatedAttributePoints += dice.Value()
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

func validateSkills(s Sheet, rb rulebook.Rulebook, availableSkillPoints int) error {
	var err error

	err = validateCoreSkillsExist(s, rb)
	if err != nil {
		return err
	}

	err = validatePermittedSkills(s, rb)
	if err != nil {
		return err
	}

	err = validateSkillPoints(s, rb, availableSkillPoints)
	if err != nil {
		return err
	}

	return nil
}

func validateCoreSkillsExist(s Sheet, rb rulebook.Rulebook) error {
RequiredCoreSkills:
	for _, coreSkill := range rb.CoreSkills() {
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

func validatePermittedSkills(s Sheet, rb rulebook.Rulebook) error {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, ok := rb.FindSkill(sheetSkill.Name)

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

func validateSkillPoints(s Sheet, rb rulebook.Rulebook, availableSkillPoints int) error {
	aggregatedSkillPoints := 0

	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, _ := rb.FindSkill(sheetSkill.Name)
			skill := rb.Traits().Skills[index]

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

			aggregatedSkillPoints += dice.Value() + pointCostModifier
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
