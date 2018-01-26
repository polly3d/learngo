package algorithms

func qsort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		qsort(values, left, p-1)
	}
	if right-p > 1 {
		qsort(values, p+1, right)
	}
}

// Qsort is a sort alogrithm
func Qsort(values []int) {
	qsort(values, 0, len(values)-1)
}
