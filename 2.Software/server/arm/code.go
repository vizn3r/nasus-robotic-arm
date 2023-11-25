package arm

import "strconv"

// G-code directs the motion and function of the CNC machine, while M-code controls the operations not involving movements

var XYZ []float64 = []float64{0.0, 0.0, 0.0}


func ResolveArgs(args []string) {
	for _, a := range args {
		switch a[:1] {
		case "X":
			d, _ := strconv.ParseFloat(a[1:], 64)
			XYZ[0] = d
		case "Y":
			d, _ := strconv.ParseFloat(a[1:], 64)
			XYZ[1] = d
		case "Z":
			d, _ := strconv.ParseFloat(a[1:], 64)
			XYZ[2] = d
		}
	}
}

// Resolves arm code
func ResolveCode(code string, args []string) {
	ResolveArgs(args)
	switch code {
		// G code
	case "G0":
		MoveXYZ(XYZ[0], XYZ[1], XYZ[2])
	}
}