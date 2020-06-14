package rulebook

import (
	"reflect"
	"testing"
)

func TestAttributes_FindAttribute(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		attrs     Attributes
		args      args
		wantIndex int
		wantFound bool
	}{
		// TODO: Add test cases.
		{
			name: "find first",
			attrs: Attributes{
				Attribute{Name: "attribute0", description: ""},
				Attribute{Name: "attribute1", description: ""},
				Attribute{Name: "attribute2", description: ""},
				Attribute{Name: "attribute3", description: ""},
				Attribute{Name: "attribute4", description: ""},
			},
			args:      args{name: "attribute0"},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "find last",
			attrs: Attributes{
				Attribute{Name: "attribute0", description: ""},
				Attribute{Name: "attribute1", description: ""},
				Attribute{Name: "attribute2", description: ""},
				Attribute{Name: "attribute3", description: ""},
				Attribute{Name: "attribute4", description: ""},
			},
			args:      args{name: "attribute4"},
			wantIndex: 4,
			wantFound: true,
		},
		{
			name: "find inbetween",
			attrs: Attributes{
				Attribute{Name: "attribute0", description: ""},
				Attribute{Name: "attribute1", description: ""},
				Attribute{Name: "attribute2", description: ""},
				Attribute{Name: "attribute3", description: ""},
				Attribute{Name: "attribute4", description: ""},
			},
			args:      args{name: "attribute2"},
			wantIndex: 2,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := tt.attrs.FindAttribute(tt.args.name)
			if gotIndex != tt.wantIndex {
				t.Errorf("Attributes.FindAttribute() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Attributes.FindAttribute() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestSkills_CoreSkills(t *testing.T) {
	tests := []struct {
		name           string
		skills         Skills
		wantCoreSkills []Skill
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "find none",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
			},
			wantCoreSkills: []Skill{},
		},
		{
			name: "find all",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: true, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
			},
			wantCoreSkills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: true, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
			},
		},
		{
			name: "find one",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
			},

			wantCoreSkills: []Skill{
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
			},
		},
		{
			name: "find some",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
			wantCoreSkills: []Skill{
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCoreSkills := tt.skills.CoreSkills(); !reflect.DeepEqual(gotCoreSkills, tt.wantCoreSkills) {
				t.Errorf("Skills.CoreSkills() = %v, want %v", gotCoreSkills, tt.wantCoreSkills)
			}
		})
	}
}

func TestSkills_FindSkill(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		skills    Skills
		args      args
		wantIndex int
		wantFound bool
	}{
		// TODO: Add test cases.
		{
			name: "find first",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
			args:      args{name: "Skill0"},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "find last",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
			args:      args{name: "Skill3"},
			wantIndex: 3,
			wantFound: true,
		},
		{
			name: "find inbetween",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
			args:      args{name: "Skill2"},
			wantIndex: 2,
			wantFound: true,
		},
		{
			name: "find none",
			skills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
			args:      args{name: "SkillNotExisting"},
			wantIndex: -1,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := tt.skills.FindSkill(tt.args.name)
			if gotIndex != tt.wantIndex {
				t.Errorf("Skills.FindSkill() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Skills.FindSkill() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
