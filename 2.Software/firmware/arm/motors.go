package arm

import (
	"fmt"
	"math"
	"time"
)

type Motor struct {
	STEP      int     // STEP pin
	DIR       int     // DIR pin
	ENABLE    int     // ENABLE pin
	StepAngle float64 // Angle of one step
}

func pow9(f float64) float64 {
	return f * math.Pow(10, 10)
}

func (m *Motor) RotateDeg(deg float64, dir int, rpm float64) {
	pStepAngle, pDeg, pRpm := pow9(m.StepAngle), pow9(deg), pow9(rpm)
	sd := pow9(pStepAngle/(pRpm*6))
	sleep := time.Duration((sd/10)*float64(time.Nanosecond))
	t := time.Now()
	for i := 0.0; i <= pDeg; i += pStepAngle {
		// make step functionality

		time.Sleep(sleep)
	}
	fmt.Println(time.Since(t))
}