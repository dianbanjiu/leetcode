package perday

// leetcode-cn number: 1512. 好数对的数目
// url: https://leetcode-cn.com/problems/number-of-good-pairs/
func NumIdenticalPairs(nums []int) int {
	var count int
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				count++
			}
		}
	}

	return count
}

func numIdenticalPairs(nums []int) int {
	var numMap = make(map[int]int)
	for _, num := range nums {
		numMap[num] += 1
	}

	var result int
	for _, count := range numMap {
		result += count * (count - 1) / 2
	}
	return result
}
