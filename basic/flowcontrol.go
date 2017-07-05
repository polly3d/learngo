package basic

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// RunFlowControl is flow control
func RunFlowControl() {
	fmt.Println("==== flow control ====")
	normalFor()
	whileFor()
	normalIf(-10)
	normalIf(10)
	newton(2)
	normalSwitch()
	switchNoCondition()
	normalDefer()
	forDefer()
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

func normalSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoog!")
	default:
		fmt.Println("Good evening!")

	}
}

func normalDefer() {
	defer fmt.Println("world !")

	fmt.Print("hello ")
}

func forDefer() {
	fmt.Println("start loop...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end loop")
}
