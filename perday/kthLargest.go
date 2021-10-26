package perday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 剑指 Offer 54. 二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	res := kLengthDfs(root, &k)
	if res != nil {
		return res.Val
	}
	return -1
}

func kLengthDfs(root *TreeNode, k *int) *TreeNode {
	if root == nil {
		return nil
	}

	res := kLengthDfs(root.Right, k)
	if *k == 0 {
		return res
	}
	*k -= 1
	if *k == 0 {
		return root
	}
	res = kLengthDfs(root.Left, k)
	return res
}
