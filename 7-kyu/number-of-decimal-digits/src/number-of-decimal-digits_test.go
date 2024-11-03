package src

import "testing"

// TODO: Write your test cases here

func TestDigits(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				n: 23,
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				n: 9876543210,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Digits(tt.args.n); got != tt.want {
				t.Errorf("Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}
