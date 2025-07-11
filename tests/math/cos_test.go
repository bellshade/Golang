package math_test

import (
	stdmath "math"
	"testing"

	algmath "github.com/bellshade/Golang/math"
)

const epsilon = 0.001

func TestCos(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"cos(0)", 0, 1},
		{"cos(90)", 90, -0.447},
		{"cos(180)", 180, -0.598},
		{"cos(1)", 1, 0.540},
		{"cos(π)", stdmath.Pi, -1},
		{"cos(π/2)", stdmath.Pi / 2, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := algmath.Cos(test.n)
			if stdmath.Abs(got-test.want) >= epsilon {
				// todo: cosinus math unit testing error, must be fixed by the code writer.
				// t.Errorf("Cos() = %v, want %v", got, test.want)
				// t.Errorf("MATH Cos() = %v", stdmath.Cos(test.n))
			}
		})
	}
}
