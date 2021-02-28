package perday

// leetcode-cn number: 5689. 统计匹配检索规则的物品数量
// url:https://leetcode-cn.com/problems/count-items-matching-a-rule/
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	var reslut int
	var matches = func(ruleIndex int) int {
		var count int
		for _, item := range items {
			if item[ruleIndex] == ruleValue {
				count++
			}
		}
		return count
	}
	switch ruleKey {
	case "type":
		reslut = matches(0)
	case "color":
		reslut = matches(1)
	case "name":
		reslut = matches((2))
	}

	return reslut
}
