package sort

import (
	"fmt"
	"learngo/example/sort/algorithms"
)

// Run is a entry point
func Run() {
	fmt.Println("sort data is [5, 4, 3, 2, 1]")
	values := []int{5, 4, 3, 2, 1}
	algorithms.BubbleSort(values)
	fmt.Println("sort result:", values)

	values = []int{5, 4, 3, 2, 1}
	algorithms.Qsort(values)
	fmt.Println("sort result:", values)
}
