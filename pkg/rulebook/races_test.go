package rulebook

import (
	"reflect"
	"testing"
)

func TestRace_Modifiers(t *testing.T) {
	type fields struct {
		name        string
		description string
		abilities   racialAbilities
	}
	tests := []struct {
		name          string
		fields        fields
		wantModifiers Modifiers
	}{
		{
			name: "nothing found when empty modifiers",
			fields: fields{
				name:        "RaceMock0",
				description: "",
				abilities:   racialAbilities{},
			},
			wantModifiers: Modifiers{},
		},
		{
			name: "found single",
			fields: fields{
				name:        "RaceMock1",
				description: "",
				abilities: racialAbilities{
					racialAbility{
						modifiers: Modifiers{
							Modifier{
								kind:     ModifierKindDiceAdjustment,
								value:    3,
								selector: Selector{},
							},
						},
					},
				},
			},
			wantModifiers: Modifiers{
				Modifier{
					kind:     ModifierKindDiceAdjustment,
					value:    3,
					selector: Selector{},
				},
			},
		},
		{
			name: "found multiple",
			fields: fields{
				name:        "RaceMock2",
				description: "",
				abilities: racialAbilities{
					racialAbility{
						modifiers: Modifiers{
							Modifier{
								kind:     ModifierKindDiceAdjustment,
								value:    3,
								selector: Selector{},
							},
						},
					},
					racialAbility{
						modifiers: Modifiers{
							Modifier{
								kind:     ModifierKindDiceAdjustment,
								value:    1,
								selector: Selector{},
							},
							Modifier{
								kind:     ModifierKindDiceValue,
								value:    1,
								selector: Selector{},
							},
						},
					},
				},
			},
			wantModifiers: Modifiers{
				Modifier{
					kind:     ModifierKindDiceAdjustment,
					value:    3,
					selector: Selector{},
				},
				Modifier{
					kind:     ModifierKindDiceAdjustment,
					value:    1,
					selector: Selector{},
				},
				Modifier{
					kind:     ModifierKindDiceValue,
					value:    1,
					selector: Selector{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				name:        tt.fields.name,
				description: tt.fields.description,
				abilities:   tt.fields.abilities,
			}
			if gotModifiers := r.Modifiers(); !reflect.DeepEqual(gotModifiers, tt.wantModifiers) {
				t.Errorf("Race.Modifiers() = %v, want %v", gotModifiers, tt.wantModifiers)
			}
		})
	}
}
