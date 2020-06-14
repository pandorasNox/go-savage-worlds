package dice

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		dice string
	}
	tests := []struct {
		name    string
		args    args
		want    Dice
		wantErr bool
	}{
		{
			name:    "parse only dice",
			args:    args{dice: "d4"},
			want:    Dice{value: "4", points: 0, adjustment: 0},
			wantErr: false,
		},
		{
			name:    "parse dice and adjustment",
			args:    args{dice: "d12+8"},
			want:    Dice{value: "12", points: 4, adjustment: 8},
			wantErr: false,
		},
		{
			name:    "parse dice missing adjustment",
			args:    args{dice: "d4+"},
			want:    Dice{},
			wantErr: true,
		},
		{
			name:    "parse dice and negative adjustment",
			args:    args{dice: "d10-2"},
			want:    Dice{value: "10", points: 3, adjustment: -2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.dice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDice() = %v, want %v", got, tt.want)
			}
		})
	}
}
