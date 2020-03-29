package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteNode(head *ListNode, val int) *ListNode {
	if head == nil{
		return nil
	}else if head!=nil&&head.Val == val{
		return head.Next
	}

	for p:=head; p.Next!=nil;p=p.Next{
		if p.Next.Val == val{
			p.Next = p.Next.Next
			break
		}
	}
	return head
}
