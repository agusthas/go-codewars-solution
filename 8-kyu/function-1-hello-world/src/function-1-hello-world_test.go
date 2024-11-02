package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "hello world!",
			want: "hello world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, greet())
		})
	}
}
