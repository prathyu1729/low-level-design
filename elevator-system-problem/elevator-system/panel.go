package elevator_system

type panel struct {
	display    display
	upButton   button
	DownButton button
}

func (p panel) pressButton(direction Direction) {
	if direction == Up {
		p.upButton.press()
	} else if direction == Down {
		p.DownButton.press()
	}
}
