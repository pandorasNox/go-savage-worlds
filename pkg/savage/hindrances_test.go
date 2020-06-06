package savage

import "testing"

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
