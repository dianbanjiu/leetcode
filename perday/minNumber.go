package perday

import (
	"sort"
	"strconv"
	"strings"
)

// 剑指 Offer 45. 把数组排成最小的数
type str []string

func minNumber(nums []int) string {
	var strs = make(str, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Sort(strs)
	return strings.Join(strs, "")
}

func (s str) Len() int {
	return len(s)
}

func (s str) Less(i, j int) bool {
	if s[i]+s[j] > s[j]+s[i] {
		return false
	}
	return true
}
func (s str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
