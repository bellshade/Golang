package searching

import (
	"math"
	"testing"

	"github.com/bellshade/Golang/algoritma/searching"
)

const EPS = 1e-6

func equal(a, b float64) bool {
	return math.Abs(a-b) <= EPS
}

func TestTernaryMax(t *testing.T) {
	tests := []struct {
		f         func(x float64) float64
		a         float64
		b         float64
		ekspetasi float64
	}{
		{f: func(x float64) float64 { return -x * x }, a: 1, b: -1, ekspetasi: 0},
	}

	for _, test := range tests {
		hasil, error := searching.TernaryMaxSearching(test.a, test.b, EPS, test.f)
		if error != nil {
			t.Errorf("error, data: `%v`", error)
		}

		if !equal(hasil, test.ekspetasi) {
			t.Errorf("hasil salah, ekspetasi: `%v`, return aktual: `%v`", test.ekspetasi, hasil)
		}
	}
}

func TestTernaryMin(t *testing.T) {
	tests := []struct {
		f         func(x float64) float64
		a         float64
		b         float64
		ekspetasi float64
	}{
		{f: func(x float64) float64 { return x * x }, a: -1, b: 1, ekspetasi: 0},
		{f: func(x float64) float64 { return 2*x*x + x + 1 }, a: -1, b: 1, ekspetasi: 0.875},
	}

	for _, test := range tests {
		hasil, error := searching.TernaryMinSearching(test.a, test.b, EPS, test.f)
		if error != nil {
			t.Errorf("error, data: `%v`", error)
		}

		if !equal(hasil, test.ekspetasi) {
			t.Errorf("hasil salah, ekspetasi: `%v`, return aktual: `%v`", test.ekspetasi, hasil)
		}
	}
}
