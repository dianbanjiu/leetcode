package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindromeList(head *ListNode) bool {
	var r = make([]int, 0)
	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}

	var flag = true
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			flag = false
			break
		}
	}
	return flag
}
