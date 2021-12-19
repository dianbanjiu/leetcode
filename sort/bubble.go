package sortCode

// 冒泡排序
// 冒泡排序是非原地排序算法，Go 的切片是引用类型，可以无需额外的返回值
func Bubble(nums []int){
	
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]>=nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}