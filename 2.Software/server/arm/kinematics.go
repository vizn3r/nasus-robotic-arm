package arm

import (
	"fmt"
)

// Movement settings
var SPEED int = 1
var CURRENT_POS []float64 = []float64{0.0, 0.0, 0.0}

func MoveXYZ(x float64, y float64, z float64) {
	fmt.Println("Moving to: " + fmt.Sprintf("%v", x) + " " + fmt.Sprintf("%v", y) + " " + fmt.Sprintf("%v", z))
}