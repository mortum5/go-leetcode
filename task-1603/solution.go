package task1603

import "github.com/emirpasic/gods/maps/hashmap"

type ParkingSystem struct {
	hashmap.Map
}

func Constructor(big int, medium int, small int) ParkingSystem {
	p := ParkingSystem{*hashmap.New()}
	p.Put(1, big)
	p.Put(2, medium)
	p.Put(3, small)
	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	count, ok := this.Get(carType)
	res := count.(int)
	if ok && count != 0 {
		this.Put(carType, res-1)
		return true
	}
	return false
}
