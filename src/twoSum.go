package src

// 时间复杂度O(n^2), 空间复杂度O(1)
//func twoSum(nums []int, target int) []int {
//	var r = make([]int, 2)
//	for i := 0; i < len(nums)-1 && nums[i] < target; i++ {
//		for j := i + 1; j < len(nums) && nums[j] < target; j++ {
//			if nums[i]+nums[j] == target {
//				r[0] = nums[i]
//				r[1] = nums[j]
//				return r
//			}
//		}
//	}
//	return r
//}

//时间复杂度O(n)，空间复杂度 O(n)
func TwoSum(nums []int, target int) []int {
	var cases = make(map[int]int)
	for _, v := range nums {
		if _, ok := cases[v]; !ok {
			cases[v] = 1
		} else {
			cases[v] += 1
		}
	}

	var r = make([]int, 2)
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if _, ok := cases[target-nums[i]]; ok {
			r[0] = nums[i]
			r[1] = target - nums[i]
			break
		}
	}
	return r
}
