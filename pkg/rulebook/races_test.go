package rulebook

import (
	"fmt"
	"testing"
)

func TestRace_Modifiers(t *testing.T) {
	m1 := func(ca CharacterAggregation) CharacterAggregation {
		return ca
	}
	m2 := func(ca CharacterAggregation) CharacterAggregation {
		return ca
	}
	m3 := func(ca CharacterAggregation) CharacterAggregation {
		return ca
	}

	type fields struct {
		name        string
		description string
		abilities   racialAbilities
	}
	tests := []struct {
		name          string
		fields        fields
		wantModifiers CharacterAggregationModifiers
	}{
		{
			name: "nothing found when empty modifiers",
			fields: fields{
				name:        "RaceMock0",
				description: "",
				abilities:   racialAbilities{},
			},
			wantModifiers: CharacterAggregationModifiers{},
		},
		{
			name: "found single",
			fields: fields{
				name:        "RaceMock1",
				description: "",
				abilities: racialAbilities{
					racialAbility{
						modifiers: CharacterAggregationModifiers{m1},
					},
				},
			},
			wantModifiers: CharacterAggregationModifiers{m1},
		},
		{
			name: "found multiple",
			fields: fields{
				name:        "RaceMock2",
				description: "",
				abilities: racialAbilities{
					racialAbility{
						modifiers: CharacterAggregationModifiers{m1},
					},
					racialAbility{
						modifiers: CharacterAggregationModifiers{m2, m3},
					},
				},
			},
			wantModifiers: CharacterAggregationModifiers{m1, m2, m3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				name:        tt.fields.name,
				description: tt.fields.description,
				abilities:   tt.fields.abilities,
			}

			if gotModifiers := fmt.Sprintf("%v", r.Modifiers()); gotModifiers != fmt.Sprintf("%v", tt.wantModifiers) {
				t.Errorf("Race.Modifiers() = %v, want %v", gotModifiers, fmt.Sprintf("%v", tt.wantModifiers))
			}
		})
	}
}

func TestRaces_FindRace(t *testing.T) {
	racesFixture := Races{
		{name: "race1"},
		{name: "race2"},
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		rs        Races
		args      args
		wantIndex int
		wantFound bool
	}{
		{
			name:      "race found",
			rs:        racesFixture,
			args:      args{name: "race2"},
			wantIndex: 1,
			wantFound: true,
		},
		{
			name:      "race not found",
			rs:        racesFixture,
			args:      args{name: "notexisting"},
			wantIndex: -1,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := tt.rs.FindRace(tt.args.name)
			if gotIndex != tt.wantIndex {
				t.Errorf("Races.FindRace() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Races.FindRace() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
