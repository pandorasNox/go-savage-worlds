package rulebook

import (
	"reflect"
	"testing"
)

func TestSheet_SheetSkill(t *testing.T) {
	sheetCharFixture := SheetCharacter{
		Traits: SheetTraits{
			Attributes: []SheetAttribute{
				{Skills: []SheetSkill{
					{Name: "skill1"},
					{Name: "skill2"},
				}},
				{Skills: []SheetSkill{
					{Name: "skill3"},
				}},
			},
		},
	}
	type fields struct {
		Character SheetCharacter
	}
	type args struct {
		skillName SkillName
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantSheetSkill SheetSkill
		wantErr        bool
	}{
		{
			name:           "skill not found",
			fields:         fields{Character: sheetCharFixture},
			args:           args{skillName: "nonexisting"},
			wantSheetSkill: SheetSkill{},
			wantErr:        true,
		},
		{
			name:           "skill not found in fst attr",
			fields:         fields{Character: sheetCharFixture},
			args:           args{skillName: "skill1"},
			wantSheetSkill: SheetSkill{Name: "skill1"},
			wantErr:        false,
		},
		{
			name:           "skill not found in snd attr",
			fields:         fields{Character: sheetCharFixture},
			args:           args{skillName: "skill3"},
			wantSheetSkill: SheetSkill{Name: "skill3"},
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sheet{
				Character: tt.fields.Character,
			}
			gotSheetSkill, err := s.SheetSkill(tt.args.skillName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sheet.SheetSkill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSheetSkill, tt.wantSheetSkill) {
				t.Errorf("Sheet.SheetSkill() = %v, want %v", gotSheetSkill, tt.wantSheetSkill)
			}
		})
	}
}
