package AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	first := l1
	second := l2

	temp := 0

	current := result
	for first != nil || second != nil {
		sum := first.Val + second.Val + temp
		current.Val = sum % 10
		sum = sum / 10

		if first.Next == nil {
			first = &ListNode{Val: 0, Next: nil}
		} else {
			first = first.Next
		}
		if second.Next == nil {
			second = &ListNode{Val: 0, Next: nil}
		} else {
			second = second.Next
		}

		newNode := &ListNode{}
		current.Next = newNode
		current = newNode
	}

	return result
}
