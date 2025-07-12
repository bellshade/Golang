package math_test

import (
	"testing"

	"github.com/bellshade/Golang/math/prima"
)

func TestPembagianTrial(t *testing.T) {
	testingData := []struct {
		name      string
		n         int64
		ekspetasi bool
	}{
		{"Prima: 2", 2, true},
		{"Prima: 3", 3, true},
		{"Bukan prima: -5", -5, false},
	}

	for _, testingTest := range testingData {
		t.Run(testingTest.name, func(t *testing.T) {
			if getData := prima.PembagianTrial(testingTest.n); getData != testingTest.ekspetasi {
				t.Errorf("PembagianTrial(%v) = %v; ekspetasi %v", testingTest.n, getData, testingTest.ekspetasi)
			}
		})
	}
}
