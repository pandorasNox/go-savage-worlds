package rulebook

import (
	"testing"
)

func Test_findDegree(t *testing.T) {
	mockHindranceMinor := Hindrance{
		Name:             "MockMinor",
		description:      "",
		AvailableDegrees: []HindranceDegree{{Degree: Minor}},
	}
	mockHindranceMajor := Hindrance{
		Name:             "MockMajor",
		description:      "",
		AvailableDegrees: []HindranceDegree{{Degree: Major}},
	}
	mockHindranceMinorMajor := Hindrance{
		Name:             "MockMinorMajor",
		description:      "",
		AvailableDegrees: []HindranceDegree{{Degree: Minor}, {Degree: Major}},
	}
	type args struct {
		hindrance  Hindrance
		degreeName string
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantOk    bool
	}{
		{
			"find degree",
			args{hindrance: mockHindranceMinor, degreeName: Minor.String()},
			0,
			true,
		},
		{
			"do not find degree (empty hindrance)",
			args{hindrance: Hindrance{}, degreeName: Minor.String()},
			-1,
			false,
		},
		{
			"do not find degree (missing degree)",
			args{hindrance: mockHindranceMajor, degreeName: Minor.String()},
			-1,
			false,
		},
		{
			"find degree (having  multiple degrees)",
			args{hindrance: mockHindranceMinorMajor, degreeName: Major.String()},
			1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindDegree(tt.args.hindrance, tt.args.degreeName)
			if got != tt.wantIndex {
				t.Errorf("findDegree() got = %v, wantIndex %v", got, tt.wantIndex)
			}
			if got1 != tt.wantOk {
				t.Errorf("findDegree() got1 = %v, wantOk %v", got1, tt.wantOk)
			}
		})
	}
}

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
