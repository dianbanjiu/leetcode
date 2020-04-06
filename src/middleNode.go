package src

func MiddleNode(head *ListNode) *ListNode {
	var cases = make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		cases = append(cases, p)
	}
	return cases[len(cases)/2]
}
