package perday

import "strconv"

// leetcode-cn number: 5686. 移动所有球到每个盒子所需的最小操作数
func MinOperations(boxes string) []int {

	var boxList = make([]int, 0)
	for boxIndex, box := range boxes {
		boxValue, _ := strconv.Atoi(string(box))
		if boxValue == 1 {
			boxList = append(boxList, boxIndex)
		}
	}
	var answers = make([]int, len(boxes))
	for answerIndex, _ := range answers {
		for _, box := range boxList {
			if box == answerIndex {
				continue
			}

			if box > answerIndex {
				answers[answerIndex] += box - answerIndex
			} else {
				answers[answerIndex] += answerIndex - box
			}
		}
	}

	return answers
}
