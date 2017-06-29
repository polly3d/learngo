package basic

import (
	"fmt"
	"math"
)

// RunVariables is running by main
func RunVariables() {
	fmt.Println("==== basic ====")
	variables()
	iniVarables()
	shortVarables()
	typeConversion()
	constData()
}

func variables() {
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java)
}

func iniVarables() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func shortVarables() {
	i, j := 1, 2
	k := "hello world"
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

func typeConversion() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}

func constData() {
	const WORD = "世界"
	fmt.Println("Hello", WORD)
}
