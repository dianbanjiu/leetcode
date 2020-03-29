package src

import (
	"math"
)

func Massage(nums []int) int {
	if len(nums)==0{
		return 0
	}
	var dp = make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		var t1,t2 int
		if i-1<0{
			t1=0
		}else{
			t1 = dp[i-1]
		}
		if i-2<0{
			t2=0
		}else{
			t2 = dp[i-2]
		}

		dp[i]= int(math.Max(float64(nums[i]+t2), float64(t1)))
	}
	return dp[len(dp)-1]
}