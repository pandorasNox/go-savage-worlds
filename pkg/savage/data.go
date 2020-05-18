package savage

import (
	"fmt"
	"regexp"
)

type Sheet struct { //playerSheet???
	Version      string       `yaml:"version"`
	RuleSet      string       `yaml:"rule-set"`
	SettingRules SettingRules `yaml:"setting-rules"`
	Character    struct {
		Info   CharacterInfo `yaml:"info"`
		Traits struct {
			Attributes struct {
				Agility struct {
					Dice   string        `yaml:"dice"`
					Skills AgilitySkills `yaml:"skills"`
				} `yaml:"agility"`
				Smarts struct {
					Dice   string       `yaml:"dice"`
					Skills SmartsSkills `yaml:"skills"`
				} `yaml:"smarts"`
				Spirit struct {
					Dice   string       `yaml:"dice"`
					Skills SpiritSkills `yaml:"skills"`
				} `yaml:"spirit"`
				Strenght struct {
					Dice string `yaml:"dice"`
				} `yaml:"strenght"`
				Vigor struct {
					Dice string `yaml:"dice"`
				} `yaml:"vigor"`
			} `yaml:"attributes"`
		} `yaml:"traits"`
		Hindrances        Hindrances        `yaml:"hindrances"`
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
	Name       string `yaml:"name"`
	Race       string `yaml:"race"`
	Gender     string `yaml:"gender"`
	Concept    string `yaml:"concept"`
	Background string `yaml:"background"`
	Height     string `yaml:"height"`
	Weight     string `yaml:"weight"`
	Wealth     int    `yaml:"wealth"`
}

type AgilitySkills struct {
	Athletics string `yaml:"athletics"`
	Boating   string `yaml:"boating"`
	Driving   string `yaml:"driving"`
	Fighting  string `yaml:"fighting"`
	Piloting  string `yaml:"piloting"`
	Riding    string `yaml:"riding"`
	Shooting  string `yaml:"shooting"`
	Stealth   string `yaml:"stealth"`
	Thievery  string `yaml:"thievery"`
}

// type Skill struct {
// 	name string
// 	attributeIBelongTo string
//  isCore bool
// }

// var skillList = []Skill{
// 	{name: "Thievery"},
// }

//todo: is this maybe a task for a code generator???

type SmartsSkills struct {
	Academics       string `yaml:"academics"`
	Battle          string `yaml:"battle"`
	CommonKnowledge string `yaml:"common-knowledge"`
	Electronics     string `yaml:"electronics"`
	Gambling        string `yaml:"gambling"`
	Hacking         string `yaml:"hacking"`
	Healing         string `yaml:"healing"`
	Language        []struct {
		Name string `yaml:"name"`
		Dice string `yaml:"dice"`
	} `yaml:"language"`
	Notice       string `yaml:"notice"`
	Occult       string `yaml:"occult"`
	Psionics     string `yaml:"psionics"`
	Repair       string `yaml:"repair"`
	Research     string `yaml:"research"`
	Science      string `yaml:"science"`
	Spellcasting string `yaml:"spellcasting"`
	Survival     string `yaml:"survival"`
	Taunt        string `yaml:"taunt"`
	WeirdScience string `yaml:"weird-science"`
}

type SpiritSkills struct {
	Faith        string `yaml:"faith"`
	Focus        string `yaml:"focus"`
	Intimidation string `yaml:"intimidation"`
	Performance  string `yaml:"performance"`
	Persuasion   string `yaml:"persuasion"`
}

type Hindrances []struct {
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

//
// var HindrancesList = []struct{
// 	name string
// 	degree string
// 	//??? -3 vigor
// 	//-3 thougness
// 	//-1 strength
// 	//-1 notice
// //adaptValue ???
// }{}

//Validate validates a savage world sheet
func (s Sheet) Validate() error {
	availableAttributePoints := baseAttributePoints
	availableSkillPoints := baseSkillPoints

	var err error

	err = s.validateAttributePoints(availableAttributePoints)
	if err != nil {
		return fmt.Errorf("sheet validation attribute errors: %s", err)
	}

	err = s.validateSkillPoints(availableSkillPoints)
	if err != nil {
		return fmt.Errorf("sheet validation skill errors: %s", err)
	}

	return nil
}

// AttributeDiceToAttributePointsUsedMap
var attributeDiceValues = map[string]int{
	"4":  0,
	"6":  1,
	"8":  2,
	"10": 3,
	"12": 4,
}

func (s Sheet) validateAttributePoints(availableAttributePoints int) error {
	aggregatedAttributePoints := 0

	attrDices := []string{
		s.Character.Traits.Attributes.Agility.Dice,
		// struct{
		//	dice
		// 	diceValMod: 1(non core skill) or 0(core skill)
		// }
		s.Character.Traits.Attributes.Smarts.Dice,
		s.Character.Traits.Attributes.Spirit.Dice,
		s.Character.Traits.Attributes.Strenght.Dice,
		s.Character.Traits.Attributes.Vigor.Dice,
	}

	var re = regexp.MustCompile(`^d(4|6|8|10|12)(\+([1-9][0-9]?))?$`)

	for _, dice := range attrDices {
		found := re.FindAllStringSubmatch(dice, -1)

		if found == nil || (len(found[0]) != 2 && len(found[0]) != 4) {
			return fmt.Errorf(
				"validation error: invalid dice value \"%s\" for path \"%s\"",
				dice,
				"n.a.", //todo: provide path
			)
		}

		aggregatedAttributePoints += attributeDiceValues[found[0][1]]
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

var diceValueToPointsUsedMap = map[string]int{
	"4":  0,
	"6":  1,
	"8":  2,
	"10": 3,
	"12": 4,
}

type SkillField struct {
	dice string
	// modifier int //todo: should always be 0 0r 1
	isCore   bool
	yamlPath string
}

func (s Sheet) validateSkillPoints(availableSkillPoints int) error {
	skillDices := []SkillField{
		{
			dice:     s.Character.Traits.Attributes.Agility.Skills.Athletics,
			isCore:   true,
			yamlPath: "character.traits.attributes.agility.skills.athletics",
		},
		{
			dice:     s.Character.Traits.Attributes.Agility.Skills.Boating,
			isCore:   false,
			yamlPath: "character.traits.attributes.agility.skills.boating",
		},
		{
			dice:     s.Character.Traits.Attributes.Agility.Skills.Driving,
			isCore:   false,
			yamlPath: "character.traits.attributes.agility.skills.driving",
		},
		{
			dice:     s.Character.Traits.Attributes.Agility.Skills.Fighting,
			isCore:   false,
			yamlPath: "character.traits.attributes.agility.skills.fighting",
		},
		{
			dice:     s.Character.Traits.Attributes.Agility.Skills.Piloting,
			isCore:   false,
			yamlPath: "character.traits.attributes.agility.skills.piloting",
		},
	}

	var re = regexp.MustCompile(`^d(4|6|8|10|12)(\+([1-9][0-9]?))?$`)

	aggregatedSkillPoints := 0

	for _, sd := range skillDices {
		found := re.FindAllStringSubmatch(sd.dice, -1)

		if found == nil || (len(found[0]) != 2 && len(found[0]) != 4) {
			return fmt.Errorf(
				"validation error: invalid dice value \"%s\" for path \"%s\"",
				sd.dice,
				sd.yamlPath, //todo: provide path
			)
		}

		coreModifier := 1
		if sd.isCore {
			coreModifier = 0
		}

		aggregatedSkillPoints += diceValueToPointsUsedMap[found[0][1]] + coreModifier
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
