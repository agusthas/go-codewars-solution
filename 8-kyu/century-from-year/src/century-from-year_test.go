package src

import "testing"

// TODO: Write your test cases here

func Test_century(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				year: 2000,
			},
			want: 20,
		},
		{
			name: "test 2",
			args: args{
				year: 2001,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := century(tt.args.year); got != tt.want {
				t.Errorf("century() = %v, want %v", got, tt.want)
			}
		})
	}
}
