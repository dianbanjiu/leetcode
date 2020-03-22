package src

type doubleSlice [][]int
func Merge(intervals [][]int) [][]int {
	var r = make(doubleSlice,0)
	r = append(r, intervals...)
	for i := 1; i < len(r); {
		if r[i][0] <= r[i-1][1] {
			r[i][0] = r[i-1][0]
			if r[i][1] <= r[i-1][1] {
				r[i][1] = r[i-1][1]
			}
			r = append(r[:i-1], r[i:]...)
			continue
		}

		i++
	}
	return r
}

func (d doubleSlice)Len()int{
	return len(d)
}

func (d doubleSlice)Less(i, j int)bool{
	return d[i][0]<=d[j][0]
}
func (d doubleSlice)Swap(i, j int){
	d[i], d[j] = d[j], d[i]
}