/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
package perday
func lengthOfLastWord(s string) int {
	var wordLen int
	for i := len(s) - 1; i >= 0; i-- {
		if wordLen != 0 && s[i] == ' ' {
			break
		}else if s[i] == ' '{
			continue
		}
		wordLen++
	}
	return wordLen
}

// @lc code=end
