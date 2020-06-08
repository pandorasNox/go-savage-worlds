package savage

import (
	"reflect"
	"testing"
)

func TestTraits_coreSkills(t *testing.T) {
	type fields struct {
		attributes []Attribute
		skills     []Skill
	}
	tests := []struct {
		name           string
		fields         fields
		wantCoreSkills []Skill
	}{
		// TODO: Add test cases.
		{
			name: "find none",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{name: "Skill0", linkedAttribute: "Attribute1", isCore: false, description: ""},
					{name: "Skill1", linkedAttribute: "Attribute2", isCore: false, description: ""},
					{name: "Skill2", linkedAttribute: "Attribute2", isCore: false, description: ""},
				},
			},
			wantCoreSkills: []Skill{},
		},
		{
			name: "find all",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{name: "Skill0", linkedAttribute: "Attribute1", isCore: true, description: ""},
					{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
					{name: "Skill2", linkedAttribute: "Attribute2", isCore: true, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{name: "Skill0", linkedAttribute: "Attribute1", isCore: true, description: ""},
				{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
				{name: "Skill2", linkedAttribute: "Attribute2", isCore: true, description: ""},
			},
		},
		{
			name: "find one",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{name: "Skill0", linkedAttribute: "Attribute1", isCore: false, description: ""},
					{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
					{name: "Skill2", linkedAttribute: "Attribute2", isCore: false, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
			},
		},
		{
			name: "find some",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{name: "Skill0", linkedAttribute: "Attribute1", isCore: false, description: ""},
					{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
					{name: "Skill2", linkedAttribute: "Attribute2", isCore: false, description: ""},
					{name: "Skill3", linkedAttribute: "Attribute4", isCore: true, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{name: "Skill1", linkedAttribute: "Attribute2", isCore: true, description: ""},
				{name: "Skill3", linkedAttribute: "Attribute4", isCore: true, description: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Traits{
				attributes: tt.fields.attributes,
				skills:     tt.fields.skills,
			}
			if gotCoreSkills := tr.coreSkills(); !reflect.DeepEqual(gotCoreSkills, tt.wantCoreSkills) {
				t.Errorf("Traits.coreSkills() = %v, want %v", gotCoreSkills, tt.wantCoreSkills)
			}
		})
	}
}
