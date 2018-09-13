package leetcode

//TwoSumNormal https://leetcode.com/problems/two-sum/description/
func TwoSumNormal(input []int, target int) []int {
	inputLen := len(input)
	if inputLen == 0 {
		return nil
	}

	for i := 0; i < inputLen; i++ {
		for j := i + 1; j < inputLen-1; j++ {
			if (input[i] + input[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//TwoSumHash is sum by hash
func TwoSumHash(input []int, target int) []int {
	m := make(map[int]int)
	for index, value := range input {
		if j, ok := m[target-value]; ok {
			return []int{j, index}
		}
		m[value] = index
	}
	return nil
}
