package src

import "testing"

// TODO: Write your test cases here

func TestScale(t *testing.T) {
	type args struct {
		s string
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "abcd\nefgh\nijkl\nmnop",
				k: 2,
				n: 3,
			},
			want: "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp",
		},
		{
			name: "test 2",
			args: args{
				s: "abcd\nefgh\nijkl\nmnop",
				k: 1,
				n: 1,
			},
			want: "abcd\nefgh\nijkl\nmnop",
		},
		{
			name: "test 3",
			args: args{
				s: "",
				k: 2,
				n: 3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scale(tt.args.s, tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}
