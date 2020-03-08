package src

// 给定一个数组，将数组元素向右移动 k 位
// 如 [1,2,3,4,5,6,7], k=3, [5,6,7,1,2,3,4]
// 移动方法，先反转数组，再反转前面 k%len(nums) 个元素，再反转后面 len(nums)-k%len(nums) 个元素

func Rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
