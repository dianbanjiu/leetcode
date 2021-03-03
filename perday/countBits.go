package perday

// leetcode-cn number: 338. 比特位计数
// url:https://leetcode-cn.com/problems/counting-bits/

// 基础题解
//func CountBits(num int) []int {
//	var result = make([]int, num+1)
//	var count = func(n int) int {
//		var c int
//		for n != 0 {
//			if n%2 != 0 {
//				c += 1
//			}
//			n /= 2
//		}
//		return c
//	}
//	for i := 0; i <= num; i++ {
//		result[i] = count(i)
//	}
//	return result
//}
//

// 进阶题解
func CountBits(num int) []int {
	var result = make([]int, num+1)
	temp := 1
	for temp < num+1 {
		result[temp] = 1
		for i := 1; i < temp && temp+i <= num; i++ {
			if result[temp+i] == 1 {
				continue
			}
			result[temp+i] = result[temp] + result[i]
		}
		temp *= 2

	}

	return result
}
