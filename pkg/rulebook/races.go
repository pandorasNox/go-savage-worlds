package rulebook

//Race defines a race
type Race struct {
	name        string
	description string
	abilities   racialAbilities
}

//Races list of race
type Races []Race

type classification int

//possible values of a classification
const (
	Positive classification = iota
	Negative
)

type racialAbility struct {
	name           string
	description    string
	classification classification
	modifiers      Modifiers
}

type racialAbilities []racialAbility

//Modifiers return the races modifiers
func (r Race) Modifiers() (modifiers Modifiers) {
	m := Modifiers{}

	for _, ab := range r.abilities {
		m = append(m, ab.modifiers...)
	}

	return m
}

//FindRace returns index int and found bool
func (rs Races) FindRace(name string) (index int, found bool) {
	for i, r := range rs {
		if r.name == name {
			return i, true
		}
	}

	return -1, false
}
