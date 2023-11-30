package parking_lot

type Vehicle interface {
	getType() VehicleType
}

type ParkingSlot interface {
	isFree() bool
	getType() VehicleType
	getId() int
}

type Payment interface {
	MakePayment(amount float64)
}
