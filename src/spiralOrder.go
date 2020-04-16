package src

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var ans = make([]int, 0)
	cs, ce := 0, len(matrix[0])-1
	rs, re := 0, len(matrix)-1
	ansLen := (ce + 1) * (re + 1)
	for i := 0; i <= len(matrix)/2; i++ {
		for e := cs; e <= ce && len(ans) < ansLen; e++ {
			ans = append(ans, matrix[rs][e])
		}
		rs++
		for e := rs; e <= re && len(ans) < ansLen; e++ {
			ans = append(ans, matrix[e][ce])
		}
		ce--
		for e := ce; e >= cs && len(ans) < ansLen; e-- {
			ans = append(ans, matrix[re][e])
		}
		re--
		for e := re; e >= rs && len(ans) < ansLen; e-- {
			ans = append(ans, matrix[e][cs])
		}
		cs++
	}
	return ans
}
