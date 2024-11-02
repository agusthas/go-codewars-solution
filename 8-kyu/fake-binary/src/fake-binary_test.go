package src

import (
	"testing"
)

func TestFakeBin(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				x: "45385593107843568",
			},
			want: "01011110001100111",
		},
		{
			name: "test 1",
			args: args{
				x: "509321967506747",
			},
			want: "101000111101101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FakeBin(tt.args.x); got != tt.want {
				t.Errorf("FakeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
