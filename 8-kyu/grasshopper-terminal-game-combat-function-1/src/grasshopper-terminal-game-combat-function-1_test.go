package src

import "testing"

// TODO: Write your test cases here

func Test_combat(t *testing.T) {
	type args struct {
		health float64
		damage float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{
				health: 100.0,
				damage: 50.0,
			},
			want: 50.0,
		},
		{
			name: "resulting health below 0",
			args: args{
				health: 100.0,
				damage: 120.0,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combat(tt.args.health, tt.args.damage); got != tt.want {
				t.Errorf("combat() = %v, want %v", got, tt.want)
			}
		})
	}
}
