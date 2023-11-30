package main

//
//import (
//	"fmt"
//	"sort"
//	"sync"
//)
//
//// ElevatorDirection represents the direction of the elevator (UP, DOWN, IDLE).
//type ElevatorDirection int
//
//const (
//	UP   ElevatorDirection = 1
//	DOWN ElevatorDirection = -1
//	IDLE ElevatorDirection = 0
//)
//
//// ElevatorButton represents a button on a floor to request an elevator.
//type ElevatorButton struct {
//	Floor  int
//	Up     chan struct{}
//	Down   chan struct{}
//	Cancel chan struct{}
//}
//
//// ElevatorCar represents an elevator car.
//type ElevatorCar struct {
//	ID        int
//	Current   int
//	Requests  map[int]bool
//	Direction ElevatorDirection
//	Mutex     sync.Mutex
//}
//
//// NewElevatorCar creates a new elevator car.
//func NewElevatorCar(id int) *ElevatorCar {
//	return &ElevatorCar{
//		ID:        id,
//		Current:   1,
//		Requests:  make(map[int]bool),
//		Direction: IDLE,
//	}
//}
//
//// ElevatorSystem represents an elevator system.
//type ElevatorSystem struct {
//	Cars           []*ElevatorCar
//	Floors         int
//	ElevatorButton [][]*ElevatorButton
//	Mutex          sync.Mutex
//}
//
//// NewElevatorSystem creates a new elevator system with the specified number of cars and floors.
//func NewElevatorSystem(numCars, numFloors int) *ElevatorSystem {
//	es := &ElevatorSystem{
//		Floors:         numFloors,
//		ElevatorButton: make([][]*ElevatorButton, numFloors+1),
//	}
//	for i := range es.ElevatorButton {
//		es.ElevatorButton[i] = make([]*ElevatorButton, numCars)
//		for j := range es.ElevatorButton[i] {
//			es.ElevatorButton[i][j] = &ElevatorButton{
//				Floor:  i,
//				Up:     make(chan struct{}),
//				Down:   make(chan struct{}),
//				Cancel: make(chan struct{}),
//			}
//		}
//	}
//	es.initCars(numCars)
//	return es
//}
//
//func (es *ElevatorSystem) initCars(numCars int) {
//	for i := 0; i < numCars; i++ {
//		car := NewElevatorCar(i + 1)
//		go es.runElevator(car)
//		es.Cars = append(es.Cars, car)
//	}
//}
//
//// RunElevator runs the elevator car logic.
//func (es *ElevatorSystem) runElevator(elevator *ElevatorCar) {
//	for {
//		select {
//		case <-elevator.RequestsChan():
//			es.processRequests(elevator)
//		}
//	}
//}
//
//// ProcessRequests processes elevator requests for a specific car.
//func (es *ElevatorSystem) processRequests(elevator *ElevatorCar) {
//	elevator.Mutex.Lock()
//	defer elevator.Mutex.Unlock()
//
//	floorRequests := make([]int, 0)
//	for floor, requested := range elevator.Requests {
//		if requested {
//			floorRequests = append(floorRequests, floor)
//		}
//	}
//
//	sort.Ints(floorRequests)
//	current := elevator.Current
//	direction := elevator.Direction
//
//	switch direction {
//	case UP:
//		for _, floor := range floorRequests {
//			if floor >= current {
//				es.moveToFloor(elevator, floor)
//			}
//		}
//	case DOWN:
//		for i := len(floorRequests) - 1; i >= 0; i-- {
//			floor := floorRequests[i]
//			if floor <= current {
//				es.moveToFloor(elevator, floor)
//			}
//		}
//	}
//
//	// Clear requests after processing
//	for floor := range elevator.Requests {
//		elevator.Requests[floor] = false
//	}
//}
//
//// MoveToFloor moves the elevator car to the specified floor.
//func (es *ElevatorSystem) moveToFloor(elevator *ElevatorCar, floor int) {
//	fmt.Printf("Elevator %d moving from floor %d to floor %d\n", elevator.ID, elevator.Current, floor)
//	elevator.Current = floor
//}
//
//// DispatchElevator dispatches an elevator to the specified floor.
//func (es *ElevatorSystem) DispatchElevator(floor, carID int, direction ElevatorDirection) {
//	if floor < 1 || floor > es.Floors {
//		fmt.Println("Invalid floor number")
//		return
//	}
//	if carID < 1 || carID > len(es.Cars) {
//		fmt.Println("Invalid car ID")
//		return
//	}
//
//	button := es.ElevatorButton[floor][carID-1]
//
//	switch direction {
//	case UP:
//		button.Up <- struct{}{}
//	case DOWN:
//		button.Down <- struct{}{}
//	}
//}
//
//// RequestElevator requests an elevator from a specific floor.
//func (es *ElevatorSystem) RequestElevator(floor, carID int, direction ElevatorDirection) {
//	if floor < 1 || floor > es.Floors {
//		fmt.Println("Invalid floor number")
//		return
//	}
//	if carID < 1 || carID > len(es.Cars) {
//		fmt.Println("Invalid car ID")
//		return
//	}
//
//	button := es.ElevatorButton[floor][carID-1]
//
//	switch direction {
//	case UP:
//		button.Up <- struct{}{}
//	case DOWN:
//		button.Down <- struct{}{}
//	}
//}
//
//// CancelRequest cancels a previous elevator request.
//func (es *ElevatorSystem) CancelRequest(floor, carID int) {
//	if floor < 1 || floor > es.Floors {
//		fmt.Println("Invalid floor number")
//		return
//	}
//	if carID < 1 || carID > len(es.Cars) {
//		fmt.Println("Invalid car ID")
//		return
//	}
//
//	button := es.ElevatorButton[floor][carID-1]
//	button.Cancel <- struct{}{}
//}
//
//// ElevatorButtonsChan returns a channel to listen for elevator button presses.
//func (es *ElevatorSystem) ElevatorButtonsChan(floor, carID int) <-chan struct{} {
//	return es.ElevatorButton[floor][carID-1].Up
//}
//
//// RequestFloor requests a floor inside an elevator.
//func (es *ElevatorSystem) RequestFloor(carID, floor int) {
//	if floor < 1 || floor > es.Floors {
//		fmt.Println("Invalid floor number")
//		return
//	}
//	if carID < 1 || carID > len(es.Cars) {
//		fmt.Println("Invalid car ID")
//		return
//	}
//
//	es.Cars[carID-1].Requests[floor] = true
//}
//
//// CancelFloorRequest cancels a previous floor request inside an elevator.
//func (es *ElevatorSystem) CancelFloorRequest(carID, floor int) {
//	if floor < 1 || floor > es.Floors {
//		fmt.Println("Invalid floor number")
//		return
//	}
//	if carID < 1 || carID > len(es.Cars) {
//		fmt.Println("Invalid car ID")
//		return
//	}
//
//	es.Cars[carID-1].Requests[floor] = false
//}
//
//// ElevatorRequestsChan returns a channel to listen for elevator requests.
//func (es *ElevatorSystem) ElevatorRequestsChan(carID int) <-chan struct{} {
//	return es.Cars[carID-1].RequestsChan()
//}
//
//// RequestsChan returns a channel to listen for elevator requests.
//func (elevator *ElevatorCar) RequestsChan() <-chan struct{} {
//	ch := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case <-elevator.RequestUpChan():
//				ch <- struct{}{}
//			case <-elevator.RequestDownChan():
//				ch <- struct{}{}
//			}
//		}
//	}()
//	return ch
//}
//
//// RequestUpChan returns a channel to listen for up requests.
//func (button *ElevatorButton) RequestUpChan() <-chan struct{} {
//	return button.Up
//}
//
//// RequestDownChan returns a channel to listen for down requests.
//func (button *ElevatorButton) RequestDownChan() <-chan struct{} {
//	return button.Down
//}
//
//// RequestCancelChan returns a channel to listen for cancel requests.
//func (button *ElevatorButton) RequestCancelChan() <-chan struct{} {
//	return button.Cancel
//}
//
//func main() {
//	numCars := 3
//	numFloors := 10
//
//	elevatorSystem := NewElevatorSystem(numCars, numFloors)
//
//	// Example: Dispatch an elevator to floor 5
//	elevatorSystem.DispatchElevator(5, 1, UP)
//
//	// Example: Request an elevator from floor 3
//	elevatorSystem.RequestElevator(3, 2, DOWN)
//
//	// Example: Request an elevator from floor 7
//	elevatorSystem.RequestElevator(7, 3, UP)
//
//	// Example: Inside the elevator, request to go to floor 8
//	elevatorSystem.RequestFloor(1, 8)
//
//	// Example: Cancel the request to go to floor 8
//	elevatorSystem.CancelFloorRequest(1, 8)
//
//	// Simulate elevator system running
//	for {
//		// Listen for elevator requests
//		select {
//		case <-elevatorSystem.ElevatorRequestsChan(1):
//			// Process elevator requests for car 1
//			fmt.Println("Elevator 1 received a request.")
//		case <-elevatorSystem.ElevatorRequestsChan(2):
//			// Process elevator requests for car 2
//			fmt.Println("Elevator 2 received a request.")
//		case <-elevatorSystem.ElevatorRequestsChan(3):
//			// Process elevator requests for car 3
//			fmt.Println("Elevator 3 received a request.")
//		}
//	}
//}
