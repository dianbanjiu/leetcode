package src
func MaxArea(height []int) int {
	var max int
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j != i; j-- {
			if height[i]<=height[j]&&height[i]*(j-i)>max {
				max = height[i]*(j-i)
			}else if height[i]>height[j]&&height[j]*(j-i)>max {
				max = height[j]*(j-i)
			}
		}
	}
	return max
}