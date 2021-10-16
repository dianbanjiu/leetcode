package perday

// 剑指 Offer 28. 对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return root == nil || recurSymmetric(root.Left, root.Right)
}

func recurSymmetric(left, right *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return recurSymmetric(left.Left, right.Right) && recurSymmetric(left.Right, right.Left)
}
