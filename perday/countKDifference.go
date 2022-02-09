package perday

// leetcode-cn number: 2006. 差的绝对值为 K 的数对数目
// leetcode-cn url: https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/
func countKDifference(nums []int, k int) int {
	var countMap = make(map[int]int)
	var result int
	for _, num := range nums {
		temp1 := num + k
		if count, ok := countMap[temp1]; ok {
			result += count
		}
		temp2 := num - k
		if count, ok := countMap[temp2]; ok {
			result += count
		}
		countMap[num]++
	}
	return result
}
