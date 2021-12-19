package sortCode

// 插入排序
// 以下是实现的原地排序
func Insertion(nums []int) {

	for i := 1; i < len(nums); i++ {
		temp := i
		for j := i - 1; j >= 0; j-- {
			if nums[temp] < nums[j] {
				nums[temp], nums[j] = nums[j], nums[temp]
				temp--
			}
		}
	}
}
