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

func (s Sheet) collectModifier(rb Rulebook) (CharacterAggregationModifiers, error) {
	var modifier CharacterAggregationModifiers

	// race modifier

	hindranceMods, err := s.collectHindranceModifier(rb.Hindrances())
	if err != nil {
		return CharacterAggregationModifiers{}, err
	}
	modifier = append(modifier, hindranceMods...)

	// edge modifier

	return modifier, nil
}

func (s Sheet) collectHindranceModifier(rbHinds Hindrances) (CharacterAggregationModifiers, error) {
	var modifier CharacterAggregationModifiers

	for _, sheetHindrance := range s.Character.Hindrances {
		index, foundHin := rbHinds.FindHindrance(sheetHindrance.Name)
		if foundHin == false {
			return CharacterAggregationModifiers{}, fmt.Errorf("hindrance \"%s\" doesn't exist", sheetHindrance.Name)
		}
		matchedHindrance := SwadeHindrances[index]

		index, foundDeg := matchedHindrance.FindDegree(sheetHindrance.Degree)
		if foundDeg == false {
			return CharacterAggregationModifiers{}, fmt.Errorf("hindrance \"%s\" doesn't have a \"%s\" degree", sheetHindrance.Name, sheetHindrance.Degree)
		}
		matchedDegree := matchedHindrance.AvailableDegrees[index]

		modifier = append(modifier, matchedDegree.Modifiers...)
	}

	return modifier, nil
}
