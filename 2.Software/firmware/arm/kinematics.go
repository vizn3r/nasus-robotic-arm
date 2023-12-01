package arm

import (
	"fmt"
	"math"
)

type Arm struct {
	t []float64 // Theta 1 - 6 in DEGREES
	r []Matrix // Rotation matrixes
}

// Return THETA in RADIANS
func (a *Arm) T(i int) float64 {
	return (a.t[i] / 180) * math.Pi
}

// Set THETA in RADIANS
func (a *Arm) SetT(i int, rad float64) {
	a.t[i] = rad
}

// Set THETA in DEGREES
func (a *Arm) SetTDeg(i int, deg float64) {
	a.t[i] = (deg / 180) * math.Pi
}

func (*Arm) RadToDeg(rad float64) float64 {
	return math.Acos(rad) * (180 / math.Pi)
}

func (a *Arm) CalcRotation() {
	R0_1 := NewMtxArr([][]float64{
		{math.Cos(a.T(0)), 0, math.Sin(a.T(0))}, 
		{math.Sin(a.T(0)), 0, -math.Cos(a.T(0))}, 
		{0, 1, 0}})
	R1_2 := NewMtxArr([][]float64{
		{math.Cos(a.T(1)), -math.Sin(a.T(1)), 0}, 
		{math.Sin(a.T(1)), math.Cos(a.T(1)), 0}, 
		{0, 0, 1}})
	R2_3 := NewMtxArr([][]float64{
		{-math.Sin(a.T(2)), 0, math.Cos(a.T(2))}, 
		{math.Cos(a.T(2)), 0, math.Sin(a.T(2))}, 
		{0, 1, 0}})
	R3_4 := NewMtxArr([][]float64{
		{math.Cos(a.T(3)), -math.Sin(a.T(3)), 0}, 
		{math.Sin(a.T(3)), math.Cos(a.T(3)), 0}, 
		{0, 0, 1}})
	R4_5 := NewMtxArr([][]float64{
		{math.Cos(a.T(4)), 0, -math.Sin(a.T(4))}, 
		{math.Sin(a.T(4)), 0, math.Cos(a.T(4))}, 
		{0, -1, 0}})
	R5_6 := NewMtxArr([][]float64{
		{math.Cos(a.T(5)), -math.Sin(a.T(5)), 0}, 
		{math.Sin(a.T(5)), math.Cos(a.T(5)), 0}, 
		{0, 0, 1}})
	a.r = []Matrix{R0_1, R1_2, R2_3, R3_4, R4_5, R5_6}

	fmt.Println("R0_1:")
	R0_1.Print()
	fmt.Println("R1_2:")
	R1_2.Print()
	fmt.Println("R2_3:")
	R2_3.Print()
	fmt.Println("R3_4:")
	R3_4.Print()
	fmt.Println("R4_5:")
	R4_5.Print()
	fmt.Println("R5_6:")
	R5_6.Print()

	fmt.Println("R0_6:")
	// R0_6 := MultMtx(R0_1, MultMtx(R1_2, MultMtx(R2_3, MultMtx(R3_4, MultMtx(R4_5, R5_6)))))
	R0_6 := MultMtxArr([]Matrix{R0_1, R1_2, R2_3, R3_4, R4_5, R5_6})
	R0_6.Map(func(f float64) float64 {
		return a.RadToDeg(f)
	})
	R0_6.Print()
}