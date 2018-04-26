package array

func rotate(arr []int, d int, n int) []int {
	var result []int
	temp := arr[:d]
	result = arr[d:n]
	result = append(result, temp...)
	return result
}

func rotateOneByOne(arr []int, d int, n int) []int {
	len := len(arr)
	for i := 0; i < d; i++ {
		temp := arr[0]
		for j := 1; j < len; j++ {
			arr[j-1] = arr[j]
		}
		arr[len-1] = temp
	}
	return arr
}
