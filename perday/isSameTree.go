package perday

// leetcode-cn number: 100. 相同的树
// leetcode-cn url: https://leetcode-cn.com/problems/same-tree/
// 深度优先解法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if q.Val == p.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
