package basic

import (
	"fmt"
)

// RunFunctions is run functions test
func RunFunctions() {
	fmt.Println("==== func ====")
	sum := functionReturnOne(1, 2)
	fmt.Println(sum)
	a, b := functionReturnMutilple(1, 2)
	fmt.Println(a, b)
	c, d := functionReturnNamed(9)
	fmt.Println(c, d)
}

func functionReturnOne(x, y int) int {
	return x + y
}

func functionReturnMutilple(x, y int) (int, int) {
	return y, x
}

func functionReturnNamed(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum
	return
}
