package perday

func TopKFrequent(words []string, k int) []string {
	var countMap = make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}
	var result = make([]string, 0)
	for key, _ := range countMap {
		result = append(result, key)
	}
	if k-len(result) <= 0 {
		result = result[:k]
	}
	return result
}
