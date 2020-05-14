package savage

type Sheet struct { //player??
	Version      string       `yaml:"version"`
	RuleSet      string       `yaml:"rule-set"`
	SettingRules SettingRules `yaml:"setting-rules"`
	Character    struct {
		Info   CharacterInfo `yaml:"info"`
		Traits struct {
			// 	Attributes struct {
			// 		Agility struct {
			// 			Dice   string `yaml:"dice"`
			// 			Skills []struct {
			// 				Athletics int    `yaml:"athletics,omitempty"`
			// 				Boating   int    `yaml:"boating,omitempty"`
			// 				Driving   int    `yaml:"driving,omitempty"`
			// 				Fighting  string `yaml:"fighting,omitempty"`
			// 				Piloting  int    `yaml:"piloting,omitempty"`
			// 				Riding    int    `yaml:"riding,omitempty"`
			// 				Shooting  int    `yaml:"shooting,omitempty"`
			// 				Stealth   int    `yaml:"stealth,omitempty"`
			// 				Thievery  int    `yaml:"thievery,omitempty"`
			// 			} `yaml:"skills"`
			// 		} `yaml:"agility"`
			// 		Smarts struct {
			// 			Dice   string `yaml:"dice"`
			// 			Skills []struct {
			// 				Academics       int `yaml:"academics,omitempty"`
			// 				Battle          int `yaml:"battle,omitempty"`
			// 				CommonKnowledge int `yaml:"common-knowledge,omitempty"`
			// 				Electronics     int `yaml:"electronics,omitempty"`
			// 				Gambling        int `yaml:"gambling,omitempty"`
			// 				Hacking         int `yaml:"hacking,omitempty"`
			// 				Healing         int `yaml:"healing,omitempty"`
			// 				Language        []struct {
			// 					Common int `yaml:"common"`
			// 				} `yaml:"language,omitempty"`
			// 				Notice       int `yaml:"notice,omitempty"`
			// 				Occult       int `yaml:"occult,omitempty"`
			// 				Psionics     int `yaml:"psionics,omitempty"`
			// 				Repair       int `yaml:"repair,omitempty"`
			// 				Research     int `yaml:"research,omitempty"`
			// 				Science      int `yaml:"science,omitempty"`
			// 				Spellcasting int `yaml:"spellcasting,omitempty"`
			// 				Survival     int `yaml:"survival,omitempty"`
			// 				Taunt        int `yaml:"taunt,omitempty"`
			// 				WeirdScience int `yaml:"weird-science,omitempty"`
			// 			} `yaml:"skills"`
			// 		} `yaml:"smarts"`
			// 		Spirit struct {
			// 			Dice   string `yaml:"dice"`
			// 			Skills []struct {
			// 				Faith        int `yaml:"faith,omitempty"`
			// 				Focus        int `yaml:"focus,omitempty"`
			// 				Intimidation int `yaml:"intimidation,omitempty"`
			// 				Performance  int `yaml:"performance,omitempty"`
			// 				Persuasion   int `yaml:"persuasion,omitempty"`
			// 			} `yaml:"skills"`
			// 		} `yaml:"spirit"`
			// 		Strenght struct {
			// 			Dice string `yaml:"dice"`
			// 		} `yaml:"strenght"`
			// 		Vigor struct {
			// 			Dice string `yaml:"dice"`
			// 		} `yaml:"vigor"`
			// 	} `yaml:"attributes"`
		} `yaml:"traits"`
		// Hindrances []struct {
		// 	Name   string `yaml:"name"`
		// 	Degree string `yaml:"degree"`
		// } `yaml:"hindrances"`
		// Edges             []string `yaml:"edges"`
		// DerivedStatistics struct {
		// 	StandardPace string `yaml:"standard-pace"`
		// 	Parry        string `yaml:"parry"`
		// 	Toughness    struct {
		// 		Base  int `yaml:"base"`
		// 		Armor int `yaml:"armor"`
		// 	} `yaml:"toughness"`
		// } `yaml:"derived-statistics"`
		// Gear      []string `yaml:"gear"`
		// Inventory []string `yaml:"inventory"`
	} `yaml:"character"`
}

type SettingRules struct {
	StartingWealth    int  `yaml:"starting-wealth,omitempty"`
	BornAHero         bool `yaml:"born-a-hero,omitempty"`
	MultipleLanguages bool `yaml:"multiple-languages,omitempty"`
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
