package src

// 从一个排序数组创建一棵二叉搜索树
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var head = new(TreeNode)
	CreateBTS(nums, head)
	return head
}

func CreateBTS(nums []int, head *TreeNode) {
	if len(nums) == 0 {
		return
	}

	head.Val = nums[len(nums)/2]
	if len(nums) > 1 {
		head.Left = new(TreeNode)
		CreateBTS(nums[:len(nums)/2], head.Left)
	}
	if len(nums)/2+1 < len(nums) {
		head.Right = new(TreeNode)
		CreateBTS(nums[len(nums)/2+1:], head.Right)
	}
}
