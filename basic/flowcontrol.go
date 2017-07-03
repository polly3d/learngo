package basic

import (
	"fmt"
	"math"
)

// RunFlowControl is flow control
func RunFlowControl() {
	fmt.Println("==== flow control ====")
	normalFor()
	whileFor()
	normalIf(-10)
	normalIf(10)
	newton(2)
}

func normalFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
}

func whileFor() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func normalIf(x float64) {
	if x < 0 {
		fmt.Println("x < 0")
	} else {
		fmt.Println("x > 0")
	}
}

func newton(x float64) float64 {
	z := 1.0
	for {
		zl := z - (z*z-x)/(2*z)
		if math.Abs(z-zl) < 0.000001 {
			z = zl
			break
		}
		z = zl
	}
	fmt.Println(z)
	return z
}
