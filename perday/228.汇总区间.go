package perday

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var result = make([]string, 0)
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			if i == len(nums)-1 {
				result = append(result, fmt.Sprintf("%d->%d", temp, nums[i]))
			}
			continue
		}
		if temp == nums[i-1] {
			result = append(result, strconv.Itoa(temp))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", temp, nums[i-1]))
		}
		temp = nums[i]
		if i == len(nums)-1 {
			result = append(result, strconv.Itoa(temp))
		}
	}
	return result
}
