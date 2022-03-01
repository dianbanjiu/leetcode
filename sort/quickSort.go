package sortCode

func Quick(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, pre, last int){
	if pre>=last {
		return
	}
	tempPre, tempLast := pre, last
	middle := (pre+last)/2
	for pre<last {
		if nums[pre]>=nums[middle]&&nums[last] < nums[middle]{
			nums[pre], nums[last] = nums[last], nums[pre]
			pre++
			last--
			continue
		}

		if nums[pre]< nums[middle] {
			pre++
		}
		if nums[last]>=nums[middle] {
			last--
		}
	}
	quickSort(nums, tempPre, middle-1)
	quickSort(nums, middle+1, tempLast)
}