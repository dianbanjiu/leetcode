package perday

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 剑指 Offer 55 - II. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := recur(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := recur(root.Right)
	if rightDepth == -1 {
		return -1
	}
	if math.Abs(float64(rightDepth-leftDepth)) <= 1 {
		if rightDepth > leftDepth {
			return rightDepth + 1
		}
		return leftDepth + 1

	}
	return -1
}
