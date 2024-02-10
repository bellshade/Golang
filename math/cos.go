package math

import "math"

func Cos(x float64) float64 {
	tp := 1.0 / (2.0 * math.Pi)
	x *= tp
	x -= 0.25 + math.Floor(x+0.25)
	x *= 16.0 * (math.Abs(x) - 0.5)
	x += 0.255 * x * (math.Abs(x) - 1.0)
	return x
}
