package artificialIntelligence

import (
	"math"
	"testing"

	"github.com/bellshade/Golang/artificialIntelligence"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSum(t *testing.T) {
	independet := []float64{1, 2, 3}
	dependent := []float64{4, 5, 6}
	stats := artificialIntelligence.DataStatistik{}

	artificialIntelligence.Sum(independet, dependent, &stats)

	if stats.SumX != 6 {
		t.Errorf("expected SumX = 6, dapat nilai %v", stats.SumX)
	}

	if stats.SumY != 15 {
		t.Errorf("expected SumY = 15, got %v", stats.SumY)
	}
}

func TestMultipy(t *testing.T) {
	independet := []float64{1, 2, 3}
	dependent := []float64{4, 5, 6}
	stats := artificialIntelligence.DataStatistik{}

	artificialIntelligence.Multiply(independet, dependent, &stats)
	expectedMultiply := 32.0

	if !almostEqual(stats.Multiply, expectedMultiply, 1e-6) {
		t.Errorf("ekspetasi multiply = %.2f, dapat nilai %.2f", expectedMultiply, stats.Multiply)
	}
}

func TestRegressiLinear(t *testing.T) {
	independet := []float64{1, 2}
	// dependent := []float64{2, 4}
	stats := artificialIntelligence.DataStatistik {
		SumX: 3,
		SumY: 6,
		XSquared: 5,
		YSquared: 20,
		Multiply: 11,
	}

	var slope, intercept float64
	n := float64(len(independet))

	sumSquareXY := stats.Multiply - (stats.SumX * stats.SumY) / n
	sumSquareXX := stats.XSquared - (math.Pow(stats.SumX, 2)) / n
	
	slope = sumSquareXY / sumSquareXX
	intercept = (stats.SumY - (slope * stats.SumX)) / n

	if !almostEqual(slope, 4.0, 1e-6) {
		t.Errorf("ekspetasi slope = 2.0, dapat nilai %.6f", slope)
	}

	if !almostEqual(intercept, -3.0, 1e-6) {
		t.Errorf("ekspetasi intercept = 0.0, dapat nilai %.6f", intercept)
	}
}
