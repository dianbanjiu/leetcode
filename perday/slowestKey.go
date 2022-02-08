package perday

// leetcode-cn number: 1629. 按键持续时间最长的键
// leetcode-cn url: https://leetcode-cn.com/problems/slowest-key/
func slowestKey(releaseTimes []int, keysPressed string) byte {
	result := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		continueTime := releaseTimes[i] - releaseTimes[i-1]
		if continueTime > maxTime {
			maxTime = continueTime
			result = keysPressed[i]
		} else if continueTime == maxTime && keysPressed[i] > result {
			result = keysPressed[i]
		}
	}
	return result
}
