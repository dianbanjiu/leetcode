package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetKthFromEnd(head *ListNode, k int) *ListNode {
	i := 0
	p:=head
	for ; p!=nil; p=p.Next{
		if i == k{
			break
		}
		i++
	}

	q:=head
	for ; p!=nil; p=p.Next{
		q = q.Next
	}
	return q
}
