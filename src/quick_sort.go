package src

func Sort(data []int) {
	QuickSort(data, 0, len(data)-1)
}

// Golang 版的快排
func QuickSort(data []int, start, end int){
	if start >= end {
		return
	}
	last := end-1
	for i := start; i < last; {
		if data[i] > data[end] {
			data[i], data[last] = data[last], data[i]
			last--
			continue
		}
		i++
	}

	if data[end] < data[last] {
		data[end],data[last] = data[last], data[end]
	}else{
		data[end],data[last+1] = data[last+1], data[end]
		last++
	}
	QuickSort(data, start, last-1)
	QuickSort(data, last+1, end)
}