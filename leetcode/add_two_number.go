package leetcode

type ListNode struct {
	Value int
	Next  *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	temp := result
	sum := 0
	for l1 != nil || l2 != nil {
		sum /= 10

		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}

		temp.Next = &ListNode{Value: sum % 10}
		temp = temp.Next
	}

	if sum/10 == 1 {
		result.Next = &ListNode{Value: 1}
	}

	return result.Next
}
