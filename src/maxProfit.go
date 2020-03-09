package src

func MaxProfit(prices []int) int {
	var profit = make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		var max = 0
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > max {
				max = prices[i] - prices[j]
			}
		}
		profit[i] = max
	}
	var max int
	for _, v := range profit {
		if v > max {
			max = v
		}
	}
	return max
}
