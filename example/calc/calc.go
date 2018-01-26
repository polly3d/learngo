package calc

import (
	"fmt"
	"learngo/example/calc/simpleMath"
	"os"
	"strconv"
)

var usage = func() {
	fmt.Println("Usage: calc command [arguments]")
}

// Run is a entry for calc
func Run() {
	args := os.Args
	if len(args) < 2 {
		usage()
		return
	}

	switch args[1] {
	case "add":
		v1, _ := strconv.Atoi(args[2])
		v2, _ := strconv.Atoi(args[3])
		result := simpleMath.Add(v1, v2)
		fmt.Println("result: ", result)
	case "sqrt":
	}
}
