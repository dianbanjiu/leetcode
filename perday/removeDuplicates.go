package perday

// leetcode-cn number: 1047. 删除字符串中的所有相邻重复项
// url: https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
func RemoveDuplicates(S string) string {
	var result = string(S[0])
	for i := 1; i < len(S); i++ {
		if len(result) > 0 && S[i] == result[len(result)-1] {
			result = result[:len(result)-1]
			continue
		}
		result += string(S[i])
	}
	return result

}
