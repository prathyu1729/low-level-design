package parking_lot

type BikeVehicle struct {
}

func (c BikeVehicle) getType() VehicleType {
	return Bike
}
