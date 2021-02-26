package perday

// leetcode-cn number:1431. 拥有最多糖果的孩子
// url:https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxValue int
	for _, candy := range candies {
		if candy > maxValue {
			maxValue = candy
		}
	}

	var result = make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= maxValue {
			result[i] = true
		}
	}

	return result
}
