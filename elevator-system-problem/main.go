package main

import (
	elevator_system "elevator-system/elevator-system"
)

func main() {
	es := elevator_system.NewElevatorSystem(15, 3)
	es.Start()
	_ = es.PressExternalButton(1, 4, elevator_system.Up)
	_ = es.PressExternalButton(1, 5, elevator_system.Up)
	select {}
}
