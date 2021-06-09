package perday

// leetcode-cn number: 1641. 统计字典序元音字符串的数目
// leetcode-cn url: https://leetcode-cn.com/problems/count-sorted-vowel-strings/

func CountVowelStrings(n int) int {
	var result = []int{1,1,1,1,1}
	for i := 1; i < n; i++ {
		for j := 1;j < 5; j++{
			result[j] += result[j-1]
		}
	}
	
	return result[0]+result[1]+result[2]+result[3]+result[4]
}