package savage

import (
	"fmt"
	"regexp"
)

type Sheet struct {
	Version      string       `yaml:"version"`
	RuleSet      string       `yaml:"rule-set"`
	SettingRules SettingRules `yaml:"setting-rules"`
	Character    struct {
		Info   CharacterInfo `yaml:"info"`
		Traits struct {
			Attributes []SheetAttribute `yaml:"attributes"`
		} `yaml:"traits"`
		Hindrances        []SheetHindrance  `yaml:"hindrances"`
		Edges             []string          `yaml:"edges"`
		DerivedStatistics DerivedStatistics `yaml:"derived-statistics"`
		Gear              []string          `yaml:"gear"`
		Inventory         []string          `yaml:"inventory"`
	} `yaml:"character"`
}

type SettingRules struct {
	StartingWealth    int  `yaml:"starting-wealth"`
	BornAHero         bool `yaml:"born-a-hero"`
	MultipleLanguages bool `yaml:"multiple-languages"`
}

type CharacterInfo struct {
	Name       string `yaml:"name" isCore:"true"`
	Race       string `yaml:"race"`
	Gender     string `yaml:"gender"`
	Concept    string `yaml:"concept"`
	Background string `yaml:"background"`
	Height     string `yaml:"height"`
	Weight     string `yaml:"weight"`
	Wealth     int    `yaml:"wealth"`
}

type SheetAttribute struct {
	Name   string       `yaml:"name"`
	Dice   string       `yaml:"dice"`
	Skills []SheetSkill `yaml:"skills"`
}

type SheetSkill struct {
	Name string `yaml:"name"`
	Dice string `yaml:"dice"`
}

type SheetHindrance struct {
	Name   string `yaml:"name"`
	Degree string `yaml:"degree"`
}

type DerivedStatistics struct {
	StandardPace string `yaml:"standard-pace"`
	Parry        int    `yaml:"parry"`
	Toughness    struct {
		Base  int `yaml:"base"`
		Armor int `yaml:"armor"`
	} `yaml:"toughness"`
}

const (
	baseAttributePoints int = 5
	baseSkillPoints     int = 12
)

//Validate validates a savage world sheet
func (s Sheet) Validate() error {
	availableAttributePoints := baseAttributePoints
	availableSkillPoints := baseSkillPoints
	earnedHindrancePoints := 0

	var err error

	err = s.validatePermittedHindrances()
	if err != nil {
		return fmt.Errorf("sheet validation hindrance errors: %s", err)
	}

	earnedHindrancePoints = s.countHindrancePoints()
	_ = earnedHindrancePoints

	err = s.validateAttributes(availableAttributePoints)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = s.validateSkills(availableSkillPoints)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

func (s Sheet) validatePermittedHindrances() error {
	for _, sheetHindrance := range s.Character.Hindrances {
		index, ok := findHindrance(sheetHindrance.Name)

		if !ok {
			return fmt.Errorf("\"%s\" is no valid hindrance", sheetHindrance.Name)
		}

		found := false
		for _, degree := range hindrances[index].availableDegrees {
			if degree.value == sheetHindrance.Degree {
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

func (s Sheet) countHindrancePoints() int {
	hindrancePoints := 0

	for _, hindrance := range s.Character.Hindrances {
		if hindrance.Degree == "minor" {
			hindrancePoints += 1
		}

		if hindrance.Degree == "major" {
			hindrancePoints += 2
		}
	}

	return hindrancePoints
}

var diceValueToPointsUsedMap = map[string]int{
	"4":  0,
	"6":  1,
	"8":  2,
	"10": 3,
	"12": 4,
}

func (s Sheet) validateAttributes(availableAttributePoints int) error {
	var err error

	err = s.validateAttributesExist()
	if err != nil {
		return err
	}

	err = s.validateAttributePoints(availableAttributePoints)
	if err != nil {
		return err
	}

	return nil
}

func (s Sheet) validateAttributesExist() error {
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

func (s Sheet) validateAttributePoints(availableAttributePoints int) error {
	aggregatedAttributePoints := 0

	var re = regexp.MustCompile(`^d(4|6|8|10|12)(\+([1-9][0-9]?))?$`)

	for _, attribute := range s.Character.Traits.Attributes {
		_, ok := findAttribute(attribute.Name)
		if ok == false {
			return fmt.Errorf("\"%s\" is no valid attribute", attribute.Name)
		}

		found := re.FindAllStringSubmatch(attribute.Dice, -1)

		if found == nil || (len(found[0]) != 2 && len(found[0]) != 4) {
			return fmt.Errorf(
				"validation error: invalid dice value \"%s\" for path \"%s\"",
				attribute.Dice,
				"n.a.", //todo: provide path
			)
		}

		aggregatedAttributePoints += diceValueToPointsUsedMap[found[0][1]]
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

func (s Sheet) validateSkills(availableSkillPoints int) error {
	var err error

	err = s.validateCoreSkillsExist()
	if err != nil {
		return err
	}

	err = s.validatePermittedSkills()
	if err != nil {
		return err
	}

	err = s.validateSkillPoints(availableSkillPoints)
	if err != nil {
		return err
	}

	return nil
}

func (s Sheet) validateCoreSkillsExist() error {
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

func (s Sheet) validatePermittedSkills() error {
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

func (s Sheet) validateSkillPoints(availableSkillPoints int) error {

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
