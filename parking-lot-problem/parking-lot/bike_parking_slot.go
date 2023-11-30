package parking_lot

type BikeParkingSlot struct {
	Id   int
	free bool
}

func (b BikeParkingSlot) isFree() bool {
	return b.free
}
