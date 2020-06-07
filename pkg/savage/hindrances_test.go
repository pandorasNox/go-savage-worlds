package savage

import (
	"testing"
)

func Test_findHindrance(t *testing.T) {
	hindrances = []Hindrance{
		{name: "Mock0", description: "", availableDegrees: []HindranceDegree{{degree: Minor}}},
		{name: "Mock1", description: "", availableDegrees: []HindranceDegree{{degree: Major}, {degree: Minor}}},
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
		wantOk    bool
	}{
		// TODO: Add test cases.
		{"find hindrance Mock0", args{"Mock0"}, 0, true},
		{"find hindrance Mock1", args{"Mock1"}, 1, true},
		{"don't find non-existing hindrance", args{"MockFOO"}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findHindrance(tt.args.name)
			if got != tt.wantIndex {
				t.Errorf("findHindrance() got = %v, want %v", got, tt.wantIndex)
			}
			if got1 != tt.wantOk {
				t.Errorf("findHindrance() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func Test_findDegree(t *testing.T) {
	mockHindranceMinor := Hindrance{
		name:             "MockMinor",
		description:      "",
		availableDegrees: []HindranceDegree{{degree: Minor}},
	}
	mockHindranceMajor := Hindrance{
		name:             "MockMajor",
		description:      "",
		availableDegrees: []HindranceDegree{{degree: Major}},
	}
	mockHindranceMinorMajor := Hindrance{
		name:             "MockMinorMajor",
		description:      "",
		availableDegrees: []HindranceDegree{{degree: Minor}, {degree: Major}},
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
			got, got1 := findDegree(tt.args.hindrance, tt.args.degreeName)
			if got != tt.wantIndex {
				t.Errorf("findDegree() got = %v, wantIndex %v", got, tt.wantIndex)
			}
			if got1 != tt.wantOk {
				t.Errorf("findDegree() got1 = %v, wantOk %v", got1, tt.wantOk)
			}
		})
	}
}
