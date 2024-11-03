package src

import "testing"

// TODO: Write your test cases here

func TestPoints(t *testing.T) {
	type args struct {
		games []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				games: []string{"3:1", "2:2", "1:3"}, // 3 + 1 + 0
			},
			want: 4,
		},
		{
			name: "test 2",
			args: args{
				games: []string{"1:1", "2:2", "3:3"}, // 1 + 1 + 1
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Points(tt.args.games); got != tt.want {
				t.Errorf("Points() = %v, want %v", got, tt.want)
			}
		})
	}
}
