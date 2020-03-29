package src

import (
	"fmt"
	"strconv"
)

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var flag = false
	if n<0{
		n = 0-n
		flag = true
	}

	var r float64 = 1
	for i := 0; i < n; i++ {
		r *= x
	}
	if flag{
		r = 1/r
	}
	rs := fmt.Sprintf("%.5f", r)
	r, _ = strconv.ParseFloat(rs, 64)
	return r
}
