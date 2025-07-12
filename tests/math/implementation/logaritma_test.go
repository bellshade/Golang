package math_implementation

import (
	"testing"

	"github.com/bellshade/Golang/math_implementation/logaritma"
)

func TestLogBase32(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{"log2(1) = 0", 1, 0},
		{"log2(128) = 7", 128, 7},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if dapat := logaritma.LogBase2(test.n); dapat != test.want {
				t.Errorf("logBase2() = %v, harusnya %v", dapat, test.want)
			}
		})
	}
}
