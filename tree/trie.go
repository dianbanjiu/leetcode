package tree

import "errors"

// 前缀树

type TrieNode struct {
	Count    int           // 记录该节点代表的单词数
	Children [26]*TrieNode // 此处以仅包含 26 个小写英文字母的单词为例
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}

// insert 前缀树插入字符串
func (trie *TrieNode) insert(s string) {
	if trie == nil {
		trie = NewTrie()
	}

	var temp = trie
	for i := 0; i < len(s); i++ {
		if temp.Children[s[i]-'a'] == nil {
			temp.Children[s[i]-'a'] = NewTrie()
		}
		temp = temp.Children[s[i]-'a']
	}
	temp.Count++

}

// search 查找前缀树中是否存在指定的字符串
func (trie *TrieNode) search(s string) bool {
	var temp = trie
	for i := 0; i < len(s); i++ {
		if temp.Children[s[i]-'a'] == nil {
			return false
		}
		temp = temp.Children[s[i]-'a']
	}

	return temp.Count > 0
}

// delete 从树中删除指定的字符串，如果字符串不存在将会返回一个报错
func (trie *TrieNode) delete(s string) error {

	var temp = trie
	for i := 0; i < len(s); i++ {
		temp = temp.Children[s[i]-'a']
		if temp == nil {
			return errors.New("string is not exist. ")
		}
	}
	if temp.Count <= 0 {
		return errors.New("string is not exist. ")
	}
	temp.Count--
	return nil
}
