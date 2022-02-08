package perday

// leetcode-cn number: 168. Excel表列名称
// leetcode-cn url: https://leetcode-cn.com/problems/excel-sheet-column-title/
func convertToTitle(columnNumber int) string {
	var result string
	for columnNumber != 0 {
		columnNumber--
		result = string(byte(columnNumber%26+'A')) + result
		columnNumber /= 26
	}
	return result
}
