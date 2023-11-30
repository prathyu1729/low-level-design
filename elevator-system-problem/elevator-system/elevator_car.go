package elevator_system

import "fmt"

type elevatorCar struct {
	button  []button
	state   state
	display display
	door    door
}

func (e *elevatorCar) move(destinationFloor int) {
	fmt.Println("Moving to floor: %d", destinationFloor)

	e.door.open()

	fmt.Println("People entering")

	e.door.close()

}
