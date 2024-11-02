package src

import (
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{-1, -2, -3, -4, -5},
		},
		{
			name: "test 2",
			args: args{
				arr: []int{-1, 2, -3, 4, -5},
			},
			want: []int{1, -2, 3, -4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
