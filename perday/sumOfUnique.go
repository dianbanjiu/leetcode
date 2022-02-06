package perday

// leetcode-cn number: 1748. 唯一元素的和
// leetcode-cn url: https://leetcode-cn.com/problems/sum-of-unique-elements/
func sumOfUnique(nums []int) int {
	var elementMap = make(map[int]int)
	var sum int
	for _, num := range nums {
		count := elementMap[num]
		if count == 0 {
			sum += num
		} else if count == 1 {
			sum -= num
		}
		elementMap[num]++
	}
	return sum
}
