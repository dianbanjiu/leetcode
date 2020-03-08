package src

func HammingWeight(num uint32) int {
	dnum := int(num)
	var count int
	for dnum != 0 {
		t := dnum % 2
		if t == 1 {
			count += 1
		}
		dnum /= 2
	}
	return count
}
