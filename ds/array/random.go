package array

import (
	"math/rand"
	"time"
)

//Random a array
func Random(arr []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
