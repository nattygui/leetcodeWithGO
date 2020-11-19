package parkingsystem

// ParkingSystem 停车系统
type ParkingSystem struct {
	last [3]int
}

// Constructor 新建一个
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		last: [3]int{small, medium, big},
	}
}

// AddCar 增加停车
func (ps *ParkingSystem) AddCar(carType int) bool {
	if ps.last[carType-1] > 0 {
		ps.last[carType]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
