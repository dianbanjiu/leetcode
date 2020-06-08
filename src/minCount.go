package src

func MinCount(coins []int) int {
	var count int
	for i := 0; i < len(coins); i++ {
		if coins[i]%2==1{
			count+= 1+(coins[i]-1)/2
		}else {
			count += coins[i]/2
		}
	}
	return count
}