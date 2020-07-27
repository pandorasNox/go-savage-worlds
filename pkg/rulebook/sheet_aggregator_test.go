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
