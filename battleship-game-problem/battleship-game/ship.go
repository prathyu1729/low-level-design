package battleship_game

type ship struct {
	state state
	size  int
	point Point
}

func newShip(point Point, size int) ship {
	return ship{
		state: live,
		size:  size,
		point: point,
	}
}

func (s ship) setState(state state) {
	s.state = state
}
