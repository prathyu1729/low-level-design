package elevator_system

type button struct {
	pressed bool
}

func (b button) press() {
	b.pressed = true
}

func (b button) unpress() {
	b.pressed = false
}
