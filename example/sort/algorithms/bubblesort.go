package algorithms

// BubbleSort is sort algorithm
func BubbleSort(values []int) {
	length := len(values)
	flag := true
	for i := 0; i < length-1; i++ {
		flag = true
		for j := 0; j < length-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
