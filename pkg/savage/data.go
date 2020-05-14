package savage

type Sheet struct { //player??
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
	Parry        string `yaml:"parry"`
	Toughness    struct {
		Base  int `yaml:"base"`
		Armor int `yaml:"armor"`
	} `yaml:"toughness"`
}
