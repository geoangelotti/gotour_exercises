package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func check(x float64) string {
	return fmt.Sprintf("Error %v", Sqrt(x)-math.Sqrt(x))
}

func main() {
	value := 42.0
	fmt.Println(Sqrt(value), check(value))
}
