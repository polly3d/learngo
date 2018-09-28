package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2}
	index := 1
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice[index+1 : len(slice)])
}

func stringCompara(s1, s2 string) string {
	s1Value := 0
	for _, s := range s1 {
		s1Value += int(s)
	}

	s2Value := 0
	for _, s := range s2 {
		s2Value += int(s)
	}

	if s1Value < s2Value {
		return s1
	}
	return s2
}
