package rulebook

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
					{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
					{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
					{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				},
			},
			wantCoreSkills: []Skill{},
		},
		{
			name: "find all",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: true, description: ""},
					{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
					{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: true, description: ""},
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
			},
		},
		{
			name: "find one",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
					{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
					{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
			},
		},
		{
			name: "find some",
			fields: fields{
				attributes: []Attribute{},
				skills: []Skill{
					{Name: "Skill0", LinkedAttribute: "Attribute1", IsCore: false, description: ""},
					{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
					{Name: "Skill2", LinkedAttribute: "Attribute2", IsCore: false, description: ""},
					{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
				},
			},
			wantCoreSkills: []Skill{
				{Name: "Skill1", LinkedAttribute: "Attribute2", IsCore: true, description: ""},
				{Name: "Skill3", LinkedAttribute: "Attribute4", IsCore: true, description: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Traits{
				Attributes: tt.fields.attributes,
				Skills:     tt.fields.skills,
			}
			if gotCoreSkills := tr.CoreSkills(); !reflect.DeepEqual(gotCoreSkills, tt.wantCoreSkills) {
				t.Errorf("Traits.coreSkills() = %v, want %v", gotCoreSkills, tt.wantCoreSkills)
			}
		})
	}
}
