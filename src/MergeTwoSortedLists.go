package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	var t int
	if l1 != nil && l2 != nil && l1.Val > l2.Val {
		t = l2.Val - 1
	} else if l1 != nil && l2 != nil && l1.Val <= l2.Val {
		t = l1.Val - 1
	}

	newl1 := &ListNode{
		Val:  t,
		Next: l1,
	}
	p := newl1
	for p.Next != nil && l2 != nil {
		if p.Val <= l2.Val && p.Next.Val > l2.Val {
			temp := p.Next
			p.Next = l2
			l2 = l2.Next
			p.Next.Next = temp
			p = temp
			continue
		}
		p = p.Next
	}
	for p != nil {
		p = p.Next
	}
	p = l2
	return newl1.Next
}
