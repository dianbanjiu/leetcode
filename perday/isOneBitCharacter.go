package perday

// leetcode-cn number: 717. 1比特与2比特字符
// leetcode-cn url: https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			if i == len(bits)-1 {
				return false
			}
		}
	}
	return true
}
