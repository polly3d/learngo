package basic

import (
	"fmt"
)

func RunError() {
	panicAndDefer()
}

func panicAndDefer() {
	defer func() {
		fmt.Println("This is from defer")
	}()

	fmt.Println("This is from normal")

	defer func() {
		fmt.Println("other defer")
	}()

	panic("This is pannic")
}
