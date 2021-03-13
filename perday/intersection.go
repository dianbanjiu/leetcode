package perday

// leetcode-cn number: 349. 两个数组的交集
// url: https://leetcode-cn.com/problems/intersection-of-two-arrays/

func Intersection(nums1 []int, nums2 []int) []int {
	var nums1Map = make(map[int]bool)
	for _, num := range nums1 {
		nums1Map[num] = true
	}

	var innerMap = make(map[int]bool)
	for _, num := range nums2 {
		if nums1Map[num] {
			innerMap[num] = true
		}
	}

	var result = make([]int, 0)
	for key := range innerMap {
		result = append(result, key)
	}

	return result
}
