package perday

// 剑指 Offer 63. 股票的最大利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrices := prices[0]
	prices[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
		if prices[i-1] < prices[i]-minPrices {
			prices[i] = prices[i] - minPrices
		} else {
			prices[i] = prices[i-1]
		}
	}

	return prices[len(prices)-1]
}
