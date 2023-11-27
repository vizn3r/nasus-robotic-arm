package arm

import (
	"fmt"
)

// Movement settings
var SPEED int = 1
var CURRENT_POS []float64 = []float64{0.0, 0.0, 0.0}

func MoveXYZ(xyz []float64) {
	fmt.Println("Moving to: " + fmt.Sprintf("%v", xyz[0]) + " " + fmt.Sprintf("%v", xyz[1]) + " " + fmt.Sprintf("%v", xyz[2]))
}