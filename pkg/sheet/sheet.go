package sheet

import "github.com/pandorasNox/go-savage-worlds/pkg/rulebook"

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

func (s Sheet) collectModifier() []rulebook.Modifier {
	var modifier []rulebook.Modifier

	// race modifier

	modifier = append(modifier, s.collectHindranceModifier()...)

	// endge modifier

	return modifier
}

func (s Sheet) collectHindranceModifier() []rulebook.Modifier {
	modifier := []rulebook.Modifier{}

	for _, sheetHindrance := range s.Character.Hindrances {
		index, _ := rulebook.FindHindrance(sheetHindrance.Name)
		matchedHindrance := rulebook.Hindrances[index]

		index, ok := rulebook.FindDegree(matchedHindrance, sheetHindrance.Degree)
		if !ok {
			continue
		}
		matchedDegree := matchedHindrance.AvailableDegrees[index]

		modifier = append(modifier, matchedDegree.Modifiers...)
	}

	return modifier
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
