package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	var cases = []struct{
		word1 string
		word2 string
		except string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, s := range cases {
		output := perday.MergeAlternately(s.word1, s.word2)
		if output != s.except {
			t.Errorf("word1 %v,word2 %v, except %v, but got %v", s.word1, s.word2,s.except, output)
		}
	}
}
