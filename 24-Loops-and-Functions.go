package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	zz := z + 1
	epsilon := 0.000001

	t := 0
	for math.Abs(zz-z) > epsilon {
		zz = z
		z = z - (z*z-x)/(2*z)
		t++
	}
	fmt.Printf("times = %d\n", t)

	return z
}

func main() {
	for x := 1; x <= 15; x++ {
		fmt.Printf("Sqrt(%d) = %f, math.Sqrt(%d) = %f\n", x, Sqrt(float64(x)), x, math.Sqrt(float64(x)))
	}
}
