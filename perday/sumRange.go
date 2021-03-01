package perday

// leetcode-cn number: 303. 区域和检索 - 数组不可变
// url: https://leetcode-cn.com/problems/range-sum-query-immutable/
type NumArray struct {
	Array []int
}

func Constructor(nums []int) NumArray {
	return NumArray{Array: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	var result = 0
	for i <= j {
		result += this.Array[i]
		i++
	}

	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
