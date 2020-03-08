package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	p := l1
	for l1.Next != nil && l2.Next != nil {
		if l1.Val <= l2.Val && l1.Next.Val > l2.Val {
			l1, l2, l1.Next, l2.Next = l2, l2.Next, l2, l1.Next
		}
		l1 = l1.Next
	}
	if l2 != nil {
		l1 = l2
	}
	return p
}
