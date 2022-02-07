package perday

import "sort"

// leetcode-cn number: 1405. 最长快乐字符串
// leetcode-cn url: https://leetcode-cn.com/problems/longest-happy-string/
func longestDiverseString(a int, b int, c int) string {
	var result = make([]byte, 0)
	type str struct {
		val byte
		cnt int
	}
	allStrings := []str{{'a', a}, {'b', b}, {'c', c}}
	for {
		sort.Slice(allStrings, func(i, j int) bool {
			return allStrings[i].cnt > allStrings[j].cnt
		})

		var hasNext bool
		for i, strElement := range allStrings {
			if strElement.cnt <= 0 {
				continue
			}
			if len(result) >= 2 && result[len(result)-2] == strElement.val && result[len(result)-1] == strElement.val {
				continue
			}
			result = append(result, strElement.val)
			allStrings[i].cnt--
			hasNext = true
			break
		}
		if !hasNext {
			return string(result)
		}
	}
}
