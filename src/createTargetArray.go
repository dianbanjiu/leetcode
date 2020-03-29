package src

func CreateTargetArray(nums []int, index []int) []int {
	var target = make([]int, 0)
	i:=0
	j:=0
	for {
		if len(target) == len(nums){
			break
		}
		t := append([]int{}, target[index[i]:]...)
		target = append([]int{},target[:index[i]]... )
		target = append(target, nums[j])
		target = append(target, t...)

		i = (i+1)%len(index)
		j= j+1
	}
	return target
}