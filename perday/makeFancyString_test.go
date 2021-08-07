package perday

import "testing"

func TestMakeFancyString(t *testing.T) {
	// 此处仅使用了示例中的两个测试用例
	var cases = []struct {
		Give string
		Want string
	}{
		{"leeetcode", "leetcode"},
		{"aaabaaaa", "aabaa"},
	}

	for _, ss := range cases {
		actual := makeFancyString(ss.Give)
		if actual != ss.Want {
			t.Errorf("Give %s, Want %s, but got %s", ss.Give, ss.Want, actual)
		}
	}
}
