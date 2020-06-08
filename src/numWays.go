package src
func NumWays(n int, relation [][]int, k int) int {
	var EndCases = make(map[int][]int)  	// 保存所有第二个元素为终点的节点
	var normalCases  = make(map[int][]int)	// 保存普通的节点
	for i := 0; i < len(relation); i++ {
		if relation[i][1]==n-1 {
			EndCases[i] = relation[i]
		}else {
			normalCases[i] = relation[i]
		}
	}

	var ans int
	for _, ec := range EndCases {
		firstElementFlag := ec[0]	// 标记每个节点的首元素
		kcounter := 0	// k 轮计数器
		i:=0
		for{
			for _, nc := range normalCases {
				if nc[1] == firstElementFlag {
					firstElementFlag = nc[0]
					kcounter++
				}
				i++
				if kcounter>=k||i>=len(relation) {
					break
				}
			}

			if firstElementFlag==0&&(kcounter>=k||i>len(relation)) {
				ans+=1
				break
			}
		}
	}
	return ans
}