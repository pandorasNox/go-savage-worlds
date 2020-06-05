package savage

import (
	"fmt"
	"regexp"
)

//Validate validates a savage world sheet
func Validate(s Sheet) error {
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

	err = validateAttributes(s, availableAttributePoints)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = validateSkills(s, availableSkillPoints)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func validatePermittedHindrances(s Sheet) error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := findHindrance(sheetHindrance.Name)

		if !ok {
			return fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		found := false
		for _, hd := range hindrances[index].availableDegrees {
			if hd.degree.String() == sheetHindrance.Degree {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf(
				"\"%s\" is no valid degree of \"%s\"",
				sheetHindrance.Degree,
				hindrances[index].name,
			)
		}
	}

	return nil
}

func validateAttributes(s Sheet, availableAttributePoints int) error {
	var err error

	err = validateAttributesExist(s)
	if err != nil {
		return err
	}

	err = validateAttributePoints(s, availableAttributePoints)
	if err != nil {
		return err
	}

	return nil
}

func validateAttributesExist(s Sheet) error {
RequiredAttributes:
	for _, attribute := range attributes {
		for _, sheetAttribute := range s.Character.Traits.Attributes {
			if attribute.name == sheetAttribute.Name {
				continue RequiredAttributes
			}
		}

		return fmt.Errorf("\"%s\" is a required attribute", attribute.name)
	}

	return nil
}

func validateAttributePoints(s Sheet, availableAttributePoints int) error {
	aggregatedAttributePoints := 0

	for _, attribute := range s.Character.Traits.Attributes {
		_, ok := findAttribute(attribute.Name)
		if ok == false {
			return fmt.Errorf("\"%s\" is no valid attribute", attribute.Name)
		}

		dice, err := ParseDice(attribute.Dice)
		if err != nil {
			return fmt.Errorf(
				"parsing dice for attribute \"%s\" failed: %s",
				attribute.Name, err,
			)
		}

		aggregatedAttributePoints += dice.value
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

func validateSkills(s Sheet, availableSkillPoints int) error {
	var err error

	err = validateCoreSkillsExist(s)
	if err != nil {
		return err
	}

	err = validatePermittedSkills(s)
	if err != nil {
		return err
	}

	err = validateSkillPoints(s, availableSkillPoints)
	if err != nil {
		return err
	}

	return nil
}

func validateCoreSkillsExist(s Sheet) error {
RequiredCoreSkills:
	for _, coreSkill := range coreSkills() {
		for _, sheetAttr := range s.Character.Traits.Attributes {
			for _, sheetSkill := range sheetAttr.Skills {
				if coreSkill.name == sheetSkill.Name {
					continue RequiredCoreSkills
				}
			}
		}

		return fmt.Errorf("\"%s\" is a required core skill", coreSkill.name)
	}

	return nil
}

func validatePermittedSkills(s Sheet) error {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, ok := findSkill(sheetSkill.Name)

			if !ok {
				return fmt.Errorf("\"%s\" is no valid skill", sheetSkill.Name)
			}

			if skills[index].linkedAttribute != sheetAttr.Name {
				return fmt.Errorf(
					"\"%s\" should belong to attribute \"%s\" and not \"%s\"",
					sheetSkill.Name,
					skills[index].linkedAttribute,
					sheetAttr.Name,
				)
			}

		}
	}

	return nil
}

func validateSkillPoints(s Sheet, availableSkillPoints int) error {

	var re = regexp.MustCompile(`^d(4|6|8|10|12)(\+([1-9][0-9]?))?$`)

	aggregatedSkillPoints := 0

	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			index, _ := findSkill(sheetSkill.Name)
			skill := skills[index]

			found := re.FindAllStringSubmatch(sheetSkill.Dice, -1)

			if found == nil || (len(found[0]) != 2 && len(found[0]) != 4) {
				return fmt.Errorf(
					"validation error: invalid dice value \"%s\" for path \"%s\"",
					sheetSkill.Dice,
					"", //todo: provide path
				)
			}

			pointCostModifier := 1
			if skill.isCore {
				pointCostModifier = 0
			}

			aggregatedSkillPoints += diceValueToPointsUsedMap[found[0][1]] + pointCostModifier
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
