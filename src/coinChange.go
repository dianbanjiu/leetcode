package src

import "math"

func CoinChange(coins []int, amount int) int {
	var res = make([]int, amount+1)
	res[0] = 0
	for i := 1; i < amount+1; i++ {
		res[i] = math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && res[i-coins[j]] != math.MaxInt64 {
				res[i] = int(math.Min(float64(res[i]), float64(res[i-coins[j]]+1)))
			}
		}
	}
	if res[amount] == math.MaxInt64 || res[amount] == 0 {
		return -1
	}
	return res[amount]
}
