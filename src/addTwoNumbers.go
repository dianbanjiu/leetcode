package src

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum = new(ListNode)
	head := sum
	var t int
	for l1 != nil && l2 != nil {
		sum.Next = new(ListNode)
		sum = sum.Next
		sum.Val = (l1.Val + l2.Val + t) % 10
		t = (l1.Val + l2.Val + t) / 10
		l1, l2 = l1.Next, l2.Next

	}

	for l1 != nil {
		sum.Next = new(ListNode)
		sum = sum.Next
		sum.Val = (l1.Val + t) % 10
		t = (l1.Val + t) / 10
		l1 = l1.Next
	}

	for l2 != nil {
		sum.Next = new(ListNode)
		sum = sum.Next

		sum.Val = (l2.Val + t) % 10
		t = (l2.Val + t) / 10
		l2 = l2.Next
	}

	if t != 0 {
		sum.Next = new(ListNode)
		sum = sum.Next

		sum.Val = t
	}
	return head.Next
}
