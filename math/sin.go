package math

import "math"

// fungsi dari sin sederhana

func Sin(x float64) float64 {
  return math.Cos((math.Pi / 2) - x)
}
