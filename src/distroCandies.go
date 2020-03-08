package src

func DistributeCandies(candies int, num_people int) []int {
	if num_people == 0 {
		return nil
	}
	var r = make([]int, num_people)
	var c int
	for i := 0; ; {
		c += 1
		if candies-c <= 0 {
			r[i] += candies
			break
		}
		candies -= c
		r[i] += c
		i = (i + 1) % num_people
	}
	return r
}
