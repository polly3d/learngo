package array

//FindMin to find min number in array
func FindMin(arr []int) int {
	result := arr[0]
	for _, v := range arr {
		if v < result {
			result = v
		}
	}
	return result
}

//FindMin2 other find method
func FindMin2(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		m := arr[i]
		for j := 0; j < len(arr); j++ {
			n := arr[j]
			if m > n {
				result = n
			}
		}
	}
	return result
}
