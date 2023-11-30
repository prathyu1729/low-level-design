package elevator_system

import (
	"errors"
	"fmt"
	"sync"
)

var once sync.Once
var instance *ElevatorSystem

type ElevatorSystem struct {
	elevators        []elevatorCar
	floors           []floor
	dispatchStrategy dispatchStrategy
	floorCount       int
	elevatorCount    int
	requestChannel   chan request
}

func NewElevatorSystem(floorNumbers, elevatorNumbers int) *ElevatorSystem {
	once.Do(func() {
		instance = &ElevatorSystem{
			elevators:        make([]elevatorCar, elevatorNumbers),
			floors:           make([]floor, floorNumbers),
			dispatchStrategy: &fcfsStrategy{},
			floorCount:       floorNumbers,
			elevatorCount:    elevatorNumbers,
			requestChannel:   make(chan request),
		}
	})

	return instance
}

func (e *ElevatorSystem) PressExternalButton(elevatorCarId, floorNumber int, direction Direction) error {
	if floorNumber >= e.floorCount || floorNumber < 0 {
		return errors.New("invalid floorNumber")
	}

	if elevatorCarId >= e.elevatorCount || elevatorCarId < 0 {
		return errors.New("invalid elevatorId")
	}

	elevatorCar := e.elevators[elevatorCarId]
	e.requestChannel <- request{floorNumber: floorNumber, direction: direction, elevatorCar: elevatorCar}
	return nil
}

func (e *ElevatorSystem) PressInternalButton(elevatorCarId int, destination int) {

}

func (e *ElevatorSystem) Start() {
	fmt.Println("Elevator system started running")
	go func() {
		for request := range e.requestChannel {
			fmt.Println("Received a request")
			e.dispatchStrategy.processRequest(request)
		}
	}()

	go func() {
		e.dispatchStrategy.start()
	}()

}
