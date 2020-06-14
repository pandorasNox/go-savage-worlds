package rulebook

import (
	"testing"
)

func TestHindrances_FindHindrance(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		hs        Hindrances
		args      args
		wantIndex int
		wantFound bool
	}{
		{
			name: "find hindrance Mock0",
			hs: []Hindrance{
				{Name: "Mock0", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				{Name: "Mock1", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
			},
			args:      args{name: "Mock0"},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "find hindrance Mock1",
			hs: []Hindrance{
				{Name: "Mock0", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				{Name: "Mock1", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
			},
			args:      args{name: "Mock1"},
			wantIndex: 1,
			wantFound: true,
		},
		{
			name: "dont find non-existing hindrance",
			hs: []Hindrance{
				{Name: "Mock0", description: "", AvailableDegrees: []HindranceDegree{{Degree: Minor}}},
				{Name: "Mock1", description: "", AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}}},
			},
			args:      args{name: "Mock foo"},
			wantIndex: -1,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := tt.hs.FindHindrance(tt.args.name)
			if gotIndex != tt.wantIndex {
				t.Errorf("Hindrances.FindHindrance() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Hindrances.FindHindrance() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestHindrance_FindDegree(t *testing.T) {
	type fields struct {
		Name             string
		description      string
		AvailableDegrees []HindranceDegree
	}
	type args struct {
		degreeName string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
		wantFound bool
	}{
		{
			name: "find degree",
			fields: fields{
				Name:             "MockMinor",
				description:      "",
				AvailableDegrees: []HindranceDegree{{Degree: Minor}},
			},
			args:      args{degreeName: Minor.String()},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "find degree",
			fields: fields{
				Name:             "MockMinor",
				description:      "",
				AvailableDegrees: []HindranceDegree{{Degree: Minor}},
			},
			args:      args{degreeName: Minor.String()},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "do not find degree (empty hindrance)",
			fields: fields{
				Name:             "MockEmpty",
				description:      "",
				AvailableDegrees: []HindranceDegree{},
			},
			args:      args{degreeName: Minor.String()},
			wantIndex: -1,
			wantFound: false,
		},
		{

			name: "do not find degree (missing degree)",
			fields: fields{
				Name:             "MockMajor",
				description:      "",
				AvailableDegrees: []HindranceDegree{{Degree: Major}},
			},
			args:      args{degreeName: Minor.String()},
			wantIndex: -1,
			wantFound: false,
		},
		{

			name: "find degree (having  multiple degrees)",
			fields: fields{
				Name:             "MockMajorMinor",
				description:      "",
				AvailableDegrees: []HindranceDegree{{Degree: Major}, {Degree: Minor}},
			},
			args:      args{degreeName: Minor.String()},
			wantIndex: 1,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hindrance{
				Name:             tt.fields.Name,
				description:      tt.fields.description,
				AvailableDegrees: tt.fields.AvailableDegrees,
			}
			gotIndex, gotFound := h.FindDegree(tt.args.degreeName)
			if gotIndex != tt.wantIndex {
				t.Errorf("Hindrance.FindDegree() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Hindrance.FindDegree() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
