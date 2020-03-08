package src

type MaxQueue struct {
	Data []int
}

func QueueConstructor() MaxQueue {
	var q MaxQueue
	q.Data = make([]int, 0)
	return q
}

func (this *MaxQueue) Max_value() int {
	max := -1
	for i := 0; i < len(this.Data); i++ {
		if this.Data[i] > max {
			max = this.Data[i]
		}
	}
	return max
}

func (this *MaxQueue) Push_back(value int) {
	this.Data = append(this.Data, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Data) == 0 {
		return -1
	}
	front := this.Data[0]
	if len(this.Data) == 1 {
		this.Data = []int{}
	} else {
		this.Data = this.Data[1:]
	}
	return front
}
