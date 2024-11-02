package src

import "testing"

// TODO: Write your test cases here

func TestOtherAngle(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "60, 30",
			args: args{
				a: 60,
				b: 30,
			},
			want: 90,
		},
		{
			name: "20, 100",
			args: args{
				a: 20,
				b: 100,
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OtherAngle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("OtherAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}
