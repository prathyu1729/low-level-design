package elevator_system

import "fmt"

type door struct {
}

func (d door) open() {
	fmt.Println("Opening door")
}

func (d door) close() {
	fmt.Println("Closing door")
}
