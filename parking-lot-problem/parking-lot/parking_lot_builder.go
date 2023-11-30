package parking_lot

type ParkingLotBuilder struct {
	parkingLot *ParkingLot
}

func NewParkingLotBuilder() *ParkingLotBuilder {
	return &ParkingLotBuilder{
		parkingLot: &ParkingLot{},
	}
}

func (b *ParkingLotBuilder) SetEntryPoints(entryPoints []string) *ParkingLotBuilder {
	// b.parkingLot.entries = entryPoints
	return b
}

func (b *ParkingLotBuilder) SetExitPoints(exitPoints []string) *ParkingLotBuilder {
	//b.parkingLot.entries = entryPoints
	return b
}

func (b *ParkingLotBuilder) SetSlots() *ParkingLotBuilder {

	return b
}
