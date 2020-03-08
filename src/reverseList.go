package src

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var cur = head
	var next = head.Next
	head.Next = nil
	for next != nil {
		next.Next, next, cur = cur, next.Next, next
	}
	return cur
}
