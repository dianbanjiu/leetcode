package perday

// leetcode-cn number: 面试题 16.01. 交换数字
// url: https://leetcode-cn.com/problems/swap-numbers-lcci/
func SwapNumbers(numbers []int) []int {
	// 使用语言提供的语法糖达成效果
	//numbers[0], numbers[1] = numbers[1], numbers[0]
	// return numbers

	// 使用简单的数学逻辑达成效果
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]
	return numbers
}
