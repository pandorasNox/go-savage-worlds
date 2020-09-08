package rulebook

import (
	"testing"

	"github.com/pandorasNox/go-savage-worlds/pkg/dice"
)

func Test_minimumAttributeValidatorBuilder(t *testing.T) {
	type args struct {
		attributeName AttributeName
		minNeededDice dice.Dice
		edgeName      string
	}
	type validatorInput struct {
		ca CharacterAggregation
		s  Sheet
		rb Rulebook
	}
	sheetFixture := Sheet{
		Character: SheetCharacter{
			Traits: SheetTraits{
				Attributes: []SheetAttribute{
					{"Agility", "d6+1", []SheetSkill{}},
				},
			},
		},
	}
	sheetFixtureInvalidDice := Sheet{
		Character: SheetCharacter{
			Traits: SheetTraits{
				Attributes: []SheetAttribute{
					{"Agility", "invalid dice", []SheetSkill{}},
				},
			},
		},
	}
	rulebookFixture := New(
		Races{},
		Hindrances{},
		SwadeAttributes,
		Skills{},
		Edges{},
	)

	tests := []struct {
		name              string
		args              args
		validatorInput    validatorInput
		wantValidationErr bool
	}{
		{
			name: "empty",
			args: args{"Agility", dice.D6, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixture,
				Rulebook{},
			},
			wantValidationErr: true,
		},
		{
			name: "attribute not in rulebook",
			args: args{"unknown attribute", dice.D6, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixture,
				rulebookFixture,
			},
			wantValidationErr: true,
		},
		{
			name: "invalid required dice level",
			args: args{"unknown attribute", dice.D6, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixtureInvalidDice,
				rulebookFixture,
			},
			wantValidationErr: true,
		},
		{
			name: "missing attrbute in sheet",
			args: args{"Smarts", dice.D6, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixture,
				rulebookFixture,
			},
			wantValidationErr: true,
		},
		{
			name: "has not required attribute level",
			args: args{"Agility", dice.D8, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixture,
				rulebookFixture,
			},
			wantValidationErr: true,
		},
		{
			name: "has required attribute level",
			args: args{"Agility", dice.D6, "edgeName"},
			validatorInput: validatorInput{
				CharacterAggregation{},
				sheetFixture,
				rulebookFixture,
			},
			wantValidationErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := minimumAttributeValidatorBuilder(tt.args.attributeName, tt.args.minNeededDice, tt.args.edgeName)
			if got := val(tt.validatorInput.ca, tt.validatorInput.s, tt.validatorInput.rb); !((got != nil) == tt.wantValidationErr) {
				t.Errorf("minimumAttributeValidatorBuilder() = %v, want %v", got, tt.wantValidationErr)
			}
		})
	}
}
