package src

import (
	"math"
	"sort"
)

// 时间复杂度O(n^2), 空间复杂度O(1)
//func twoSum(nums []int, target int) []int {
//	var r = make([]int, 2)
//	for i := 0; i < len(nums)-1 && nums[i] < target; i++ {
//		for j := i + 1; j < len(nums) && nums[j] < target; j++ {
//			if nums[i]+nums[j] == target {
//				r[0] = nums[i]
//				r[1] = nums[j]
//				return r
//			}
//		}
//	}
//	return r
//}

//时间复杂度O(n)，空间复杂度 O(n)
func TwoSum(nums []int, target int) []int {
	var cases = make(map[int]int)
	for _, v := range nums {
		if _, ok := cases[v]; !ok {
			cases[v] = 1
		} else {
			cases[v] += 1
		}
	}

	var r = make([]int, 2)
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if _, ok := cases[target-nums[i]]; ok {
			r[0] = nums[i]
			r[1] = target - nums[i]
			break
		}
	}
	return r
}

// 给出n个骰子，计算所有可能出现点数和的概率
func TwoSum2(n int)[]float64{
	var sum = make([]float64,6*n)	// 保存返回结果
	var cases = make(map[int]int)	// 保存每种情况出现的次数
	// 初始化基础点数
	for i:=0;i<6;i++{
		cases[i+1] = 1
	}

	// 遍历计算所有情况
	for i := 1; i < n; i++ {
		for j := 1; j <= 6; j++ {
			tc := cases
			for k,_:=range tc{
				cases[k+j] += 1
			}
		}
	}

	i := 0
	for _, v := range cases {
		sum[i] = float64(v)/math.Pow(6,float64(n))
		i++
	}
	sort.Float64s(sum)
	return sum
}