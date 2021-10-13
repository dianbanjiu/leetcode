package src

// merge two sort list
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var p1, p2 *ListNode
	if l1.Val > l2.Val {
		p1, p2 = l2, l1
	} else {
		p1, p2 = l1, l2
	}

	for p1.Next != nil && p2.Next != nil {
		if p1.Val <= p2.Val && p1.Next.Val > p2.Val {
			p1, p2, p1.Next, p2.Next = p2, p2.Next, p2, p1.Next
		} else {
			p1 = p1.Next
		}
	}

	if p2.Next == nil {
		for ; p1.Next != nil; p1 = p1.Next {
			if p1.Val <= p2.Val && p1.Next.Val > p2.Val {
				p1.Next, p2.Next = p2, p1.Next
				break
			}
		}

	}
	if p2 != nil {
		p1.Next = p2
	}

	if l1.Val > l2.Val {
		return l2
	}
	return l1
}
