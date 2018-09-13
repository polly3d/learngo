package main

import (
	"fmt"
	"learngo/leetcode"
)

func main() {
	//result := leetcode.TwoSumNormal([]int{2, 7, 11, 15}, 9)
	//fmt.Println(result)
	//result = leetcode.TwoSumHash([]int{2, 7, 11, 15}, 9)
	//fmt.Println(result)

	//tool.Ls()

	l1 := &leetcode.ListNode{Value: 2, Next: &leetcode.ListNode{Value: 4, Next: &leetcode.ListNode{Value: 3}}}
	l2 := &leetcode.ListNode{Value: 5, Next: &leetcode.ListNode{Value: 6, Next: &leetcode.ListNode{Value: 4}}}
	result := leetcode.AddTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Value)
		result = result.Next
	}
}

// func main() {
// 	// data := tool.FetchStars("summerblue/laravel-shop")
// 	// fmt.Println(data)

// 	result := []tool.Repos{}
// 	reposes := tool.FetchURLs("./tool/test.md")

// 	fmt.Println("start fetch repos info...")
// 	for _, repos := range reposes {
// 		data := tool.FetchStars(repos)
// 		result = append(result, data)
// 	}

// 	fmt.Println(result)

// 	fmt.Println("result sorting...")
// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].Stars > result[j].Stars
// 	})

// 	fmt.Println("result is: ")
// 	fmt.Println(result)

// 	for _, repos := range result {
// 		fmt.Printf("%s %d\n", repos.Name, repos.Stars)
// 	}

// }
