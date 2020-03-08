package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle(head *ListNode) bool {
	var t = make(map[int]interface{})
	var flag = false
	for p := head; p != nil; p = p.Next {
		if _, ok := t[p.Val]; ok && t[p.Val] == &p.Val {
			flag = true
			break
		}
		t[p.Val] = &p.Val
	}
	return flag
}
