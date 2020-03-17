package src


func Permute(nums []int) [][]int {
	var r = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			t := append([]int{}, nums...)
			t[i], t[j] = t[j], t[i]
			r = append(r, t)
		}
	}
	return r
}