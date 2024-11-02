package src

import (
	"reflect"
	"testing"
)

// TODO: Write your test cases here

func Test_multipleOfIndex(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				ints: []int{0, 1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "negative numbers",
			args: args{
				ints: []int{22, -1, -4},
			},
			want: []int{-1, -4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multipleOfIndex(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multipleOfIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
