package perday

// leetcode-cn number:1672.最富有客户的资产总量
// url：https://leetcode-cn.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		var result int
		for _, ac := range account {
			result += ac
		}

		if result > max {
			max = result
		}
	}

	return max
}