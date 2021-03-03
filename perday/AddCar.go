package perday

// leetcode-cn number: 1603. 设计停车系统
// url:https://leetcode-cn.com/problems/design-parking-system/
type ParkingSystem struct {
	Big, Medium, Small int
}

func CarConstructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	var result bool
	switch carType {
	case 1:
		this.Big--
		result = this.Big >= 0
	case 2:
		this.Medium--
		result = this.Medium >= 0
	case 3:
		this.Small--
		result = this.Small >= 0
	}

	return result
}
