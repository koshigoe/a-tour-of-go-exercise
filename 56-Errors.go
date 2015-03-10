package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := f
	zz := z + 1
	epsilon := 0.000001

	t := 0
	for math.Abs(zz-z) > epsilon {
		zz = z
		z = z - (z*z-f)/(2*z)
		t++
	}
	fmt.Printf("times = %d\n", t)

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
