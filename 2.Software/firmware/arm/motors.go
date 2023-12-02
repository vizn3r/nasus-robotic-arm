package arm

import (
	"fmt"
	"math"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
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
	err := rpio.Open()
	if err != nil {
		return
	}
	defer rpio.Close()
	pin := rpio.Pin(m.STEP)
	pStepAngle, pDeg, pRpm := pow9(m.StepAngle), pow9(deg), pow9(rpm)
	sd := pow9(pStepAngle/(pRpm*6))
	sleep := time.Duration((sd/10)*float64(time.Nanosecond) / 2)
	t := time.Now()
	for i := 0.0; i <= pDeg; i += pStepAngle {
		pin.High()
		time.Sleep(sleep)
		pin.Low()
		time.Sleep(sleep)

	}
	fmt.Println(time.Since(t))
}