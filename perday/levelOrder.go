package perday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 32 - I. 从上到下打印二叉树
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result = make([]int, 0)
	var list = make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) != 0 {
		node := list[0]
		result = append(result, node.Val)
		list = list[1:]

		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
	}
	return result
}

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result = make([][]int, 0)
	var currentLevel = make([]*TreeNode, 0)
	currentLevel = append(currentLevel, root)
	for len(currentLevel) != 0 {
		var nextLevel = make([]*TreeNode, 0)
		var levelResult = make([]int, len(currentLevel))
		for i, node := range currentLevel {
			levelResult[i] = node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		currentLevel = nextLevel
		result = append(result, levelResult)
	}

	return result
}
