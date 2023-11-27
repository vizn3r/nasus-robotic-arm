package arm

import (
	"math"
)

// Calculate the required angle to get from source to destination with step and return steps (rounded)
func AngleToSteps(deg float64, destDeg float64, step float64) float64 {
	res := (deg - destDeg) / step
	return math.Round(res)
}

// Convert steps to degrees with step
func StepsToAngle(deg float64, steps float64, step float64) float64 {
	return deg + (steps * step)
}

//dano was here