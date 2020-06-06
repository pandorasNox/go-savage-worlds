package savage

import (
	"reflect"
	"testing"
)

func TestParseDice(t *testing.T) {
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
			want:    Dice{value: 0, adjustment: 0},
			wantErr: false,
		},
		{
			name:    "parse dice and adjustment",
			args:    args{dice: "d12+8"},
			want:    Dice{value: 4, adjustment: 8},
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
			want:    Dice{value: 3, adjustment: -2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDice(tt.args.dice)
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
