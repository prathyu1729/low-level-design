package elevator_system

import "errors"

type ElevatorSystem struct {
	elevators        []elevatorCar
	floors           []floor
	dispatchStrategy dispatchStrategy
	floorCount       int
	elevatorCount    int
}

func NewElevatorSystem(floorNumbers, elevatorNumbers int) *ElevatorSystem {

}

func (e *ElevatorSystem) PressExternalButton(elevatorCarId, floorNumber int, direction Direction) error {
	if floorNumber >= e.floorCount || floorNumber < 0 {
		return errors.New("invalid floorNumber")
	}

	if elevatorCarId >= e.elevatorCount || elevatorCarId < 0 {
		return errors.New("invalid elevatorId")
	}

	floor := e.floors[floorNumber]
	panel := floor.panels[elevatorCarId]
	panel.pressButton(direction)

	return nil
}

func (e *ElevatorSystem) PressInternalButton(elevatorCarId int, destination int) {

}
