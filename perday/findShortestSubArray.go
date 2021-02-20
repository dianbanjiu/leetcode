package perday

// leetcode-cn numeber: 697:数组的度
// url: https://leetcode-cn.com/problems/degree-of-an-array/

func FindShortestSubArray(nums []int) int {
	type degreeInfo struct {
		degree int
		arrayStart int
		arrayEnd int
	}

	var numsMap = make(map[int]degreeInfo)
	var maxNum int
	for numsIndex, num := range nums {
		degree := numsMap[num].degree + 1

		var arrayStart, arrayEnd int
		if degree == 1 {
			arrayStart = numsIndex
		}else {
			arrayStart = numsMap[num].arrayStart
		}
		arrayEnd = numsIndex

		isShort := (numsMap[maxNum].arrayEnd-numsMap[maxNum].arrayStart - (arrayEnd - arrayStart))>0
		if degree > numsMap[maxNum].degree||(degree == numsMap[maxNum].degree&&isShort){
			maxNum = num
		}

		numsMap[num] = degreeInfo{
			degree: degree,
			arrayStart: arrayStart,
			arrayEnd: arrayEnd,
		}
	}


	return numsMap[maxNum].arrayEnd-numsMap[maxNum].arrayStart+1
}
