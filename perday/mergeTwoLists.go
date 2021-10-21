package perday

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 剑指 Offer 25. 合并两个排序的链表
// 这道题总是做了忘，忘了做，其实挺简单的一道题，只需要多加一个头节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	var temp = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}

	return head.Next

}
