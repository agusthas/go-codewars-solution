package src

import (
	"reflect"
	"testing"
)

// TODO: Write your test cases here

func TestIncrementer(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				n: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{2, 4, 6, 8, 0, 2, 4, 6, 8},
		},
		{
			name: "empty array",
			args: args{
				n: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Incrementer(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Incrementer() = %v, want %v", got, tt.want)
			}
		})
	}
}
