package rulebook

import "fmt"

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
	Name           string `yaml:"name" isCore:"true"`
	Rank           string `yaml:"rank"`
	Race           string `yaml:"race"`
	Gender         string `yaml:"gender"`
	Concept        string `yaml:"concept"`
	Background     string `yaml:"background"`
	Height         string `yaml:"height"`
	Weight         string `yaml:"weight"`
	Wealth         int    `yaml:"wealth"`
	ShakenRecovery string `yaml:"shaken-recovery"`
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
	Size         int    `yaml:"size"`
	Toughness    struct {
		Base  int `yaml:"base"`
		Armor int `yaml:"armor"`
	} `yaml:"toughness"`
}

// SheetAttribute returns SheetAttribute by attributeName
func (s Sheet) SheetAttribute(attributeName AttributeName) (sheetAttribute SheetAttribute, err error) {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		if sheetAttr.Name == string(attributeName) {
			return sheetAttr, nil
		}
	}

	return SheetAttribute{}, fmt.Errorf("couldn't find attribute %s in sheet", attributeName)
}

// SheetSkill returns sheetSkill by skillName
func (s Sheet) SheetSkill(skillName SkillName) (sheetSkill SheetSkill, err error) {
	for _, sheetAttr := range s.Character.Traits.Attributes {
		for _, sheetSkill := range sheetAttr.Skills {
			if sheetSkill.Name == string(skillName) {
				return sheetSkill, nil
			}
		}
	}

	return SheetSkill{}, fmt.Errorf("couldn't find skill %s in sheet", skillName)
}
