package savage

import (
	"reflect"
	"testing"
)

func Test_coreSkills(t *testing.T) {
	skills = []Skill{
		{name: "Skill0", linkedAttribute: "Attribute1", isCore: true, description: ""},
		{name: "Skill1", linkedAttribute: "Attribute2", isCore: false, description: ""},
		{name: "Skill2", linkedAttribute: "Attribute2", isCore: false, description: ""},
	}
	tests := []struct {
		name           string
		wantCoreSkills []Skill
	}{
		{
			"find core skills",
			[]Skill{
				{name: "Skill0", linkedAttribute: "Attribute1", isCore: true, description: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCoreSkills := coreSkills(); !reflect.DeepEqual(gotCoreSkills, tt.wantCoreSkills) {
				t.Errorf("coreSkills() = %v, want %v", gotCoreSkills, tt.wantCoreSkills)
			}
		})
	}
}
