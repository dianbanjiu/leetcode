package perday

// leetcode-cn number: 1189. “气球” 的最大数量
// leetcode-cn url: https://leetcode-cn.com/problems/maximum-number-of-balloons/
func maxNumberOfBalloons(text string) int {
	var balloonMap = make(map[byte]int)
	for i := 0; i < len(text); i++ {
		balloonMap[text[i]]++
	}
	cnta, cntb, cntl, cnto, cntn := balloonMap['a'], balloonMap['b'], balloonMap['l']/2, balloonMap['o']/2, balloonMap['n']
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	return min(min(min(min(cnta, cntb), cntl), cnto), cntn)
}
