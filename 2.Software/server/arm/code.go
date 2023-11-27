package arm

import (
	"strconv"
	"strings"
)

// G-code directs the motion and function of the CNC machine, while M-code controls the operations not involving movements

func ResolveArgs(args []string) []float64 {
	arr := []float64{0.0, 0.0, 0.0, 0.0}
	for _, a := range args {
		switch a[:1] {
		case "X":
			d, _ := strconv.ParseFloat(a[1:], 64)
			arr[0] = d
		case "Y":
			d, _ := strconv.ParseFloat(a[1:], 64)
			arr[1] = d
		case "Z":
			d, _ := strconv.ParseFloat(a[1:], 64)
			arr[2] = d
		case "F":
			d, _ := strconv.ParseFloat(a[1:], 64)
			arr[3] = d
		}
		
	}
	// ALWAYS IN FORMAT X Y Z F
	return arr
}

// If I want to return smth, just make channel for that
var gcodes = []func(...string) {
	func(args ...string) { // Linear move
		MoveXYZ(ResolveArgs(args))
	},
}

// Resolves arm code
func ResolveCode(args []string) {
	// ResolveArgs(args)
	cLow := strings.ToLower(args[0])
	if strings.HasPrefix(cLow, "g") {
		i, _ := strconv.Atoi(args[0][1:])
		gcodes[i](args[1:]...)
	}
}