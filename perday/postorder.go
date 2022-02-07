package perday

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// leetcode-cn number: 590. N 叉树的后序遍历
// leetcode-cn url: https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
func postorder(root *NAnyTreeNode) []int {
	if root == nil {
		return nil
	}

	var result = make([]int, 0)
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}
