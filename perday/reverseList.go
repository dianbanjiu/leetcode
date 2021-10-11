package perday

// 剑指 Offer 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil
	for next != nil {
		temp := next.Next
		next.Next = head
		head = next
		next = temp
	}
	return head
}
