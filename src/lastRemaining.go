package src
func LastRemaining(n int, m int) int {
	//var list = make([]int,n)
	//for i := 0; i < n; i++ {
	//	list[i] = i
	//}
	//p := 0
	//for len(list)!=1 {
	//	i := (p+m-1)%len(list)
	//	list = append(list[:i], list[i+1:]...)
	//	p = i%len(list)
	//}
	//return list[0]
	if n == 1 {
		return 0
	}

	return (LastRemaining(n-1,m)+m)%n
}