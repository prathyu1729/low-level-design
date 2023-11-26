package elevator_system

type elevatorCar struct {
	button  internalButton
	state   state
	display display
	door    door
}

func (e *elevatorCar) move(destinationFloor int) {

}
