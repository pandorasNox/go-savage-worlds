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
