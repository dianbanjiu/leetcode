package src

// 时间复杂度过高
func SumFourDivisors(nums []int) int {
	var cases = make(map[int]int)
	for i:=0;i<len(nums);i++{
		cases[nums[i]] += 1
	}
	var sum int
	for k,v := range cases{
		t := div(k)
		if len(t)==4{
			sum += (t[0]+t[1]+t[2]+t[3])*v
			delete(cases,t[0])
			delete(cases,t[1])
			delete(cases,t[2])
			delete(cases,t[3])
		}
	}
	
	return sum
}

func div(n int) []int {
	var r = make([]int, 0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			r = append(r, i)
		}
		if len(r)>4{
			return []int{}
		}
	}
	r = append(r,n)
	return r
}