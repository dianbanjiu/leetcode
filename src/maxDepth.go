package src

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var tree = make([]*TreeNode, 0)
	tree = append(tree, root)
	var count int
	for len(tree) != 0 {
		var temp = make([]*TreeNode, 0)
		for i := 0; i < len(tree); i++ {
			if tree[i].Left != nil {
				temp = append(temp, tree[i].Left)
			}
			if tree[i].Right != nil {
				temp = append(temp, tree[i].Right)
			}
		}
		count++
		tree = temp
	}
	return count
}
