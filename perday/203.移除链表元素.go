/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package perday
func removeElements(head *ListNode, val int) *ListNode {
	var tempHead = &ListNode{
		Next: head,
		Val: -1,
	}
	var runner = tempHead
	for runner.Next!=nil{
		if runner.Next.Val == val {
			runner.Next = runner.Next.Next
		}else{
			runner = runner.Next
		}
	}
	return tempHead.Next
}
// @lc code=end

