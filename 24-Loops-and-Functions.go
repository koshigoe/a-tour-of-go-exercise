package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	for x := 1; x <= 15; x++ {
		fmt.Printf("Sqrt(%d) = %f, math.Sqrt(%d) = %f\n", x, Sqrt(float64(x)), x, math.Sqrt(float64(x)))
	}
}
