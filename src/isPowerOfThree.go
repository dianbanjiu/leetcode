package src

func IsPowerOfThree(n int) bool {
	var r = true
	var left int
	for n != 0 {
		left += n % 3
		n /= 3
	}
	if left != 1 {
		r = false
	}
	return r
}
