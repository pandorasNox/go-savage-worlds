package rulebook

import (
	"reflect"
	"testing"
)

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
