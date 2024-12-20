package lab4

import (
	"math"
	"fmt"
)

func calculateY(x float64) float64 {
	y := math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
	return y
}

func TaskA(xn, xk, xdel float64) []float64 {
	var yValues []float64
	for x := xn; x <= xk; x += xdel {
		yValues = append(yValues, calculateY(x))
	}
	return yValues
}

func TaskB(xv []float64) []float64 {
	var yValues []float64
	for _, x := range xv {
		yValues = append(yValues, calculateY(x))
	}
	return yValues
}
func RunLab4() {
	var xn float64 = 0.26
	var xk float64 = 0.66
	var xdel float64 = 0.08
	var xv []float64 = []float64{0.1, 0.35, 0.4, 0.55, 0.6}

	var resultA []float64 = TaskA(xn, xk, xdel)
	fmt.Println(resultA)
	var resultB []float64 = TaskB(xv)
	fmt.Println(resultB)
}