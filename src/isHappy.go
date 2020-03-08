package src

func IsHappy(n int) bool {
	var cases = make(map[int]int)
	var flag bool
	for {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 {
			flag = true
			break
		} else if _, ok := cases[sum]; ok {
			flag = false
			break
		} else {
			cases[sum] = 1
			n = sum
		}
	}
	return flag
}
