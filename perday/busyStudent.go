package perday

// leetcode-cn number: 1450. 在既定时间做作业的学生人数
// url: https://leetcode-cn.com/problems/number-of-students-doing-homework-at-a-given-time/submissions/
func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	var count int
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			count ++
		}
	}

	return count
}