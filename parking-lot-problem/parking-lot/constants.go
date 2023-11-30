package parking_lot

type VehicleType int

const (
	Car VehicleType = iota
	Bike
)

type EntryPoint int

const (
	EntryA EntryPoint = iota
	EntryB
	EntryC
)

type ExitPoint int

const (
	ExitA ExitPoint = iota
	ExitB
	ExitC
)
