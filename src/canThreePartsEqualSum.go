package src
/*
将一个数组分成三个子数组，三个子数组的元素和相等，
也就是寻找数组和的三等分点
 */
func CanThreePartsEqualSum(A []int) bool {
	var flag = false
	i := 0
	for ; i < len(A)-2; i++ {
		if sum(A[:i+1]) == sum(A)/3 {
			break
		}
	}
	for j:=i+1; j < len(A)-1; j++ {
		if sum(A[i + 1:j + 1]) == sum(A)/3 {
			flag = true
			break
		}
	}

	return flag
}

func sum(n []int) int{
	var sum int
	for _, v := range n {
		sum+=v
	}
	return sum
}