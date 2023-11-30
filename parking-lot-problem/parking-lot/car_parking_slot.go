package parking_lot

type CarParkingSlot struct {
	Id   int
	free bool
}

func (c CarParkingSlot) isFree() bool {
	return c.free
}
