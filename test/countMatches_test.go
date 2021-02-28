package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestCountMatches(t *testing.T) {
	var testCases = []struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		expect    int
	}{
		{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver", 1},
		{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "type", "phone", 2},
	}

	for _, testCase := range testCases {
		output := perday.CountMatches(testCase.items, testCase.ruleKey, testCase.ruleValue)
		if output != testCase.expect {
			t.Errorf("item is %v, rulekey is %v, rulevalue is %v, expect %v, but got %v", testCase.items, testCase.ruleKey, testCase.ruleValue, testCase.expect, output)
		}
	}
}
