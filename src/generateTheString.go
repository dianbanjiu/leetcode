package src

func GenerateTheString(n int) string {
	var s string
	var flag int
	if n%2 == 1 {
		flag = n
	} else {
		flag = n - 1
	}
	for i := 0; i < flag; i++ {
		s += "a"
	}
	for i := 0; i < n-flag; i++ {
		s += "b"
	}
	return s
}
