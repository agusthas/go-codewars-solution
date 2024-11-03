package src

import (
	"reflect"
	"testing"
)

func TestCalculateYears(t *testing.T) {
	type args struct {
		years int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [3]int
	}{
		{
			name: "test 1",
			args: args{
				years: 1,
			},
			wantResult: [3]int{1, 15, 15},
		},
		{
			name: "test 2",
			args: args{
				years: 2,
			},
			wantResult: [3]int{2, 24, 24},
		},
		{
			name: "test 3",
			args: args{
				years: 10,
			},
			wantResult: [3]int{10, 56, 64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CalculateYears(tt.args.years); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CalculateYears() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
