package perday

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// leetcode-cn number: 589. N 叉树的前序遍历
// leetcode-cn url: https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
func preorder(root *NAnyTreeNode) []int {
	if root == nil {
		return nil
	}

	var result = make([]int, 0)
	result = append(result, root.Val)
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}
	return result
}
