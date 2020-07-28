package rulebook

import (
	"testing"
)

func Test_aggregateAttributePointsUsed(t *testing.T) {
	type args struct {
		s          Sheet
		attributes Attributes
	}
	tests := []struct {
		name           string
		args           args
		wantPointsUsed int
		wantErr        bool
	}{
		{
			name: "returns 0 if no attributes set",
			args: args{
				Sheet{},
				Attributes{},
			},
			wantPointsUsed: 0,
			wantErr:        false,
		},
		{
			name: "one attribute default dice",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{Name: "mock1", Dice: "d4"},
				}}}},
				Attributes{{
					Name: "mock1",
				}},
			},
			wantPointsUsed: 0,
			wantErr:        false,
		},
		{
			name: "one attribute dice with points",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{Name: "mock1", Dice: "d8"},
				}}}},
				Attributes{{
					Name: "mock1",
				}},
			},
			wantPointsUsed: 2,
			wantErr:        false,
		},
		{
			name: "multiple attributes dice with points",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{Name: "mock1", Dice: "d6"},
					{Name: "mock2", Dice: "d8"},
				}}}},
				Attributes{
					{Name: "mock1"},
					{Name: "mock2"},
				},
			},
			wantPointsUsed: 3,
			wantErr:        false,
		},
		{
			name: "non existing attribute in sheet",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{Name: "mock1", Dice: "d8"},
					{Name: "mock2", Dice: "d8"},
				}}}},
				Attributes{{
					Name: "mock1",
				}},
			},
			wantPointsUsed: 0,
			wantErr:        true,
		},
		{
			name: "invalid dice for attribute",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{Name: "mock1", Dice: "invalid dice"},
				}}}},
				Attributes{{
					Name: "mock1",
				}},
			},
			wantPointsUsed: 0,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPointsUsed, err := aggregateAttributePointsUsed(tt.args.s, tt.args.attributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("aggregateAttributePointsUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPointsUsed != tt.wantPointsUsed {
				t.Errorf("aggregateAttributePointsUsed() = %v, want %v", gotPointsUsed, tt.wantPointsUsed)
			}
		})
	}
}

func Test_aggregateSkillPointsUsed(t *testing.T) {
	type args struct {
		s      Sheet
		skills Skills
	}
	tests := []struct {
		name           string
		args           args
		wantPointsUsed int
		wantErr        bool
	}{
		{
			name: "invalid dice for skill",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d4",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "invalid dice"},
						},
					},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 0,
			wantErr:        true,
		},
		{
			name: "invalid skill in sheet",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d4",
						Skills: []SheetSkill{
							{Name: "invalid skill", Dice: "d6"},
						},
					},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 0,
			wantErr:        true,
		},
		{
			name: "aggregate multiple mixed skill points",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d6",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						},
					},
					{
						Name: "mock2",
						Dice: "d6",
						Skills: []SheetSkill{
							{Name: "mockSkill2", Dice: "d6"},
						},
					},
				}}}},
				Skills{
					{Name: "mockSkill1", IsCore: true},
					{Name: "mockSkill2"},
				},
			},
			wantPointsUsed: 3,
			wantErr:        false,
		},
		{
			name: "aggregate single core skill points",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d6",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1", IsCore: true},
				},
			},
			wantPointsUsed: 1,
			wantErr:        false,
		},
		{
			name: "aggregate single non core skill points",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d6",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 2,
			wantErr:        false,
		},
		{
			name: "no points when no skills",
			args: args{
				Sheet{},
				Skills{},
			},
			wantPointsUsed: 0,
			wantErr:        false,
		},
		{
			name: "skill higher than attribute cost 2 more",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d4",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 3,
			wantErr:        false,
		},
		{
			name: "core skill upgrade cost more",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d4",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1", IsCore: true},
				},
			},
			wantPointsUsed: 2,
			wantErr:        false,
		},
		{
			name: "skill upgrade equals cost one per lvl",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d8",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d8"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 3,
			wantErr:        false,
		},
		{
			name: "invalid attribute dice",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "invalid",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d8"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1"},
				},
			},
			wantPointsUsed: 0,
			wantErr:        true,
		},
		{
			name: "core skill below attribute lvl",
			args: args{
				Sheet{Character: SheetCharacter{Traits: SheetTraits{Attributes: []SheetAttribute{
					{
						Name: "mock1",
						Dice: "d8",
						Skills: []SheetSkill{
							{Name: "mockSkill1", Dice: "d6"},
						}},
				}}}},
				Skills{
					{Name: "mockSkill1", IsCore: true},
				},
			},
			wantPointsUsed: 1,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPointsUsed, err := aggregateSkillPointsUsed(tt.args.s, tt.args.skills)
			if (err != nil) != tt.wantErr {
				t.Errorf("aggregateSkillPointsUsed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPointsUsed != tt.wantPointsUsed {
				t.Errorf("aggregateSkillPointsUsed() = %v, want %v", gotPointsUsed, tt.wantPointsUsed)
			}
		})
	}
}

func Test_aggregateHindrancePointsEarned(t *testing.T) {
	type args struct {
		s  Sheet
		hs Hindrances
	}
	tests := []struct {
		name             string
		args             args
		wantPointsEarned int
		wantErr          bool
	}{
		{
			name: "no hindrances no points",
			args: args{
				Sheet{},
				Hindrances{},
			},
			wantPointsEarned: 0,
			wantErr:          false,
		},
		{
			name: "invalid hindrance in sheet",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "Invalid SheetHindrance", Degree: "minor"},
				}}},
				Hindrances{},
			},
			wantPointsEarned: 0,
			wantErr:          true,
		},
		{
			name: "invalid hindrance degree in sheet",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance1", Degree: "invalid"},
				}}},
				Hindrances{
					{Name: "mockHindrance1", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 0,
			wantErr:          true,
		},
		{
			name: "minor hindrance eq 1 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance1", Degree: "minor"},
				}}},
				Hindrances{
					{Name: "mockHindrance1", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				},
			},
			wantPointsEarned: 1,
			wantErr:          false,
		},
		{
			name: "major hindrance eq 2 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance2", Degree: "major"},
				}}},
				Hindrances{
					{Name: "mockHindrance2", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 2,
			wantErr:          false,
		},
		{
			name: "minor + major hindrance eq 3 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance1", Degree: "minor"},
					{Name: "mockHindrance2", Degree: "major"},
				}}},
				Hindrances{
					{Name: "mockHindrance1", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance2", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 3,
			wantErr:          false,
		},
		{
			name: "4x minor hindrance eq 4 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance10", Degree: "minor"},
					{Name: "mockHindrance20", Degree: "minor"},
					{Name: "mockHindrance30", Degree: "minor"},
					{Name: "mockHindrance40", Degree: "minor"},
				}}},
				Hindrances{
					{Name: "mockHindrance10", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance20", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance30", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance40", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				},
			},
			wantPointsEarned: 4,
			wantErr:          false,
		},
		{
			name: "5x minor hindrance eq 5 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance10", Degree: "minor"},
					{Name: "mockHindrance20", Degree: "minor"},
					{Name: "mockHindrance30", Degree: "minor"},
					{Name: "mockHindrance40", Degree: "minor"},
					{Name: "mockHindrance50", Degree: "minor"},
				}}},
				Hindrances{
					{Name: "mockHindrance10", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance20", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance30", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance40", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance50", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				},
			},
			wantPointsEarned: 5,
			wantErr:          false,
		},
		{
			name: "2x major hindrance eq 4 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance11", Degree: "major"},
					{Name: "mockHindrance21", Degree: "major"},
				}}},
				Hindrances{
					{Name: "mockHindrance11", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
					{Name: "mockHindrance21", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 4,
			wantErr:          false,
		},
		{
			name: "3x major hindrance eq 6 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance11", Degree: "major"},
					{Name: "mockHindrance21", Degree: "major"},
					{Name: "mockHindrance31", Degree: "major"},
				}}},
				Hindrances{
					{Name: "mockHindrance11", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
					{Name: "mockHindrance21", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
					{Name: "mockHindrance31", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 6,
			wantErr:          false,
		},
		{
			name: "2x minor + 3x major hindrance eq 8 point",
			args: args{
				Sheet{Character: SheetCharacter{Hindrances: SheetHindrances{
					{Name: "mockHindrance11", Degree: "minor"},
					{Name: "mockHindrance11", Degree: "major"},
					{Name: "mockHindrance20", Degree: "minor"},
					{Name: "mockHindrance21", Degree: "major"},
					{Name: "mockHindrance31", Degree: "major"},
				}}},
				Hindrances{
					{Name: "mockHindrance11", AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}}},
					{Name: "mockHindrance20", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
					{Name: "mockHindrance21", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
					{Name: "mockHindrance31", AvailableDegrees: []HindranceDegree{{Degree: Major}}},
				},
			},
			wantPointsEarned: 8,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPointsEarned, err := aggregateHindrancePointsEarned(tt.args.s, tt.args.hs)
			if (err != nil) != tt.wantErr {
				t.Errorf("aggregateHindrancePointsEarned() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPointsEarned != tt.wantPointsEarned {
				t.Errorf("aggregateHindrancePointsEarned() = %v, want %v", gotPointsEarned, tt.wantPointsEarned)
			}
		})
	}
}
