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
	var cases = make([]int, 6)
	for i := 0; i < 6; i++ {
		cases[i] = i+1
	}

	for i := 1; i < n; i++ {
		clen:=len(cases)
		for j := 1; j <= 6; j++ {
			for k := 0; k < clen; k++ {
				cases = append(cases, cases[k]+j)
			}
		}
		cases = cases[int(math.Pow(6,float64(i))):]
	}

	var ans = make([]float64, 0)
	sort.Ints(cases)
	i, j := 0,0
	for j < len(cases) {
		if cases[i]==cases[j] {
			j++
		}else {
			p := float64(j-i)/float64(len(cases))
			ans = append(ans, p)
			i,j = j, j+1
		}
	}
	ans = append(ans,float64(j-i)/float64(len(cases)))
	return ans
}