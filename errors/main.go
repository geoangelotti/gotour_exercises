package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for math.Abs(z*z-x) > 1e-6 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	value := 42.0
	fmt.Println(Sqrt(value))
	fmt.Println(Sqrt(-value))
}
