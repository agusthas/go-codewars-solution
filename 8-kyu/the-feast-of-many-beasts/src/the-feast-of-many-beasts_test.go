package src

import "testing"

// TODO: Write your test cases here

func TestFeast(t *testing.T) {
	type args struct {
		beast string
		dish  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "macan sumatra",
			args: args{
				beast: "macan sumatra",
				dish:  "m a",
			},
			want: true,
		},
		{
			name: "burung",
			args: args{
				beast: "burung",
				dish:  "apel",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Feast(tt.args.beast, tt.args.dish); got != tt.want {
				t.Errorf("Feast() = %v, want %v", got, tt.want)
			}
		})
	}
}
