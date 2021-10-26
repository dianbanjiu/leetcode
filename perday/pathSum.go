package perday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 剑指 Offer 34. 二叉树中和为某一值的路径
func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	pathDfs(root, target)
	return res
}

var res = make([][]int, 0)
var path = make([]int, 0)

func pathDfs(root *TreeNode, target int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		var temp = make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}

	pathDfs(root.Left, target)
	pathDfs(root.Right, target)
	if len(path) != 0 {
		path = path[:len(path)-1]
	}
}
