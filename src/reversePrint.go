package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
	var ans = make([]int, 0)
	for p:=head;p!=nil;p=p.Next{
		ans = append([]int{p.Val}, ans...)
	}
	return ans
}