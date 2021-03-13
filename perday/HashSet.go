package perday
// leetcode-cn number: 705. 设计哈希集合
// url: https://leetcode-cn.com/problems/design-hashset/
type MyHashSet struct {
 Element map[int]bool
}


/** Initialize your data structure here. */
func HashSetConstructor() MyHashSet {
	return MyHashSet{Element: map[int]bool{}}
}


func (this *MyHashSet) Add(key int)  {
	this.Element[key] = true
}


func (this *MyHashSet) Remove(key int)  {
	delete(this.Element, key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.Element[key]
}