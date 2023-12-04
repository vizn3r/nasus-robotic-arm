package arm

import (
	"fmt"
	"math"
	"time"

	"firmware/gpio"
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

func (m *Motor) RotateDeg(deg float64, dir int, ms int) {
	err := gpio.Open(m.STEP)
	if err != nil {
		return
	}
	defer gpio.Close()
	fmt.Println(time.Duration(ms)*time.Microsecond)
	pStepAngle, pDeg := pow9(m.StepAngle), pow9(deg)
	t := time.Now()
	for i := 0.0; i < pDeg; i += pStepAngle {
		gpio.Write(m.STEP, "1")
		time.Sleep(time.Duration(ms)*time.Microsecond)
		gpio.Write(m.STEP, "0")	
		time.Sleep(time.Duration(ms)*time.Microsecond)
	}
	fmt.Println(time.Since(t))
}