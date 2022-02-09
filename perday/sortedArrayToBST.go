package perday

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// leetcode-cn number: 108. 将有序数组转换为二叉搜索树
// leetcode-cn url: https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	var root TreeNode
	buildBST(&root, nums, 0, len(nums)-1)
	return &root
}

func buildBST(root *TreeNode, nums []int, pre, last int) *TreeNode {
	if pre > last {
		return nil
	}
	middle := (pre + last) / 2
	root.Val = nums[middle]
	root.Left = buildBST(&TreeNode{}, nums, pre, middle-1)
	root.Right = buildBST(&TreeNode{}, nums, middle+1, last)
	return root
}
